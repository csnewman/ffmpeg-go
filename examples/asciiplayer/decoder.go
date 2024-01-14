// Based on doc/examples/decode_video.c
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/csnewman/ffmpeg-go"
	"log"
	"slices"
	"unsafe"
)

type Decoder struct {
	fmtCtx        *ffmpeg.AVFormatContext
	decCtx        *ffmpeg.AVCodecContext
	vsi           int
	bufferSrcCtx  *ffmpeg.AVFilterContext
	bufferSinkCtx *ffmpeg.AVFilterContext
	timeBase      *ffmpeg.AVRational
}

func Open(url string, width int, height int) (*Decoder, error) {
	var fmtCtx *ffmpeg.AVFormatContext

	urlPtr := ffmpeg.ToCStr(url)
	defer urlPtr.Free()

	if _, err := ffmpeg.AVFormatOpenInput(&fmtCtx, urlPtr, nil, nil); err != nil {
		return nil, err
	}

	ffmpeg.AVDumpFormat(fmtCtx, 0, urlPtr, 0)

	// Find best stream
	if _, err := ffmpeg.AVFormatFindStreamInfo(fmtCtx, nil); err != nil {
		return nil, err
	}

	var dec *ffmpeg.AVCodec

	vsi, err := ffmpeg.AVFindBestStream(fmtCtx, ffmpeg.AVMediaTypeVideo, -1, -1, &dec, 0)
	if err != nil {
		return nil, err
	}

	log.Println("Stream index", vsi)

	// Create decoding context
	decCtx := ffmpeg.AVCodecAllocContext3(dec)

	vs := fmtCtx.Streams().Get(uintptr(vsi))

	if _, err := ffmpeg.AVCodecParametersToContext(decCtx, vs.Codecpar()); err != nil {
		return nil, err
	}

	// Init decoder
	if _, err := ffmpeg.AVCodecOpen2(decCtx, dec, nil); err != nil {
		return nil, err
	}

	name := ffmpeg.ToCStr("buffer")
	defer name.Free()

	buffersrc := ffmpeg.AVFilterGetByName(name)

	name = ffmpeg.ToCStr("buffersink")
	defer name.Free()

	buffersink := ffmpeg.AVFilterGetByName(name)
	timeBase := vs.TimeBase()
	filterGraph := ffmpeg.AVFilterGraphAlloc()

	args := fmt.Sprintf(
		"video_size=%vx%v:pix_fmt=%v:time_base=%v/%v:pixel_aspect=%v/%v\n",
		decCtx.Width(), decCtx.Height(),
		decCtx.PixFmt(),
		timeBase.Num(), timeBase.Den(),
		decCtx.SampleAspectRatio().Num(),
		decCtx.SampleAspectRatio().Den(),
	)

	var bufferSrcCtx *ffmpeg.AVFilterContext

	name = ffmpeg.ToCStr("in")
	defer name.Free()

	argsC := ffmpeg.ToCStr(args)
	defer argsC.Free()

	_, err = ffmpeg.AVFilterGraphCreateFilter(
		&bufferSrcCtx,
		buffersrc,
		name,
		argsC,
		nil,
		filterGraph,
	)
	if err != nil {
		return nil, err
	}

	var bufferSinkCtx *ffmpeg.AVFilterContext

	name = ffmpeg.ToCStr("out")
	defer name.Free()

	_, err = ffmpeg.AVFilterGraphCreateFilter(
		&bufferSinkCtx,
		buffersink,
		name,
		nil,
		nil,
		filterGraph,
	)
	if err != nil {
		return nil, err
	}

	name = ffmpeg.ToCStr("pix_fmts")
	defer name.Free()

	pixFmts := []ffmpeg.AVPixelFormat{
		ffmpeg.AVPixFmtGray8,
	}

	if _, err := ffmpeg.AVOptSetSlice(bufferSinkCtx.RawPtr(), name, pixFmts, ffmpeg.AVOptSearchChildren); err != nil {
		return nil, err
	}

	outputs := ffmpeg.AVFilterInoutAlloc()
	inputs := ffmpeg.AVFilterInoutAlloc()

	outputs.SetName(ffmpeg.ToCStr("in"))
	outputs.SetFilterCtx(bufferSrcCtx)
	outputs.SetPadIdx(0)
	outputs.SetNext(nil)

	inputs.SetName(ffmpeg.ToCStr("out"))
	inputs.SetFilterCtx(bufferSinkCtx)
	inputs.SetPadIdx(0)
	inputs.SetNext(nil)

	desc := ffmpeg.ToCStr(fmt.Sprintf(
		"scale=w=%v:h=%v,fps=fps=10,settb=expr=%v/%v",
		width, height,
		timeBase.Num(), timeBase.Den(),
	))
	defer desc.Free()

	if _, err := ffmpeg.AVFilterGraphParsePtr(filterGraph, desc, &inputs, &outputs, nil); err != nil {
		return nil, err
	}

	if _, err := ffmpeg.AVFilterGraphConfig(filterGraph, nil); err != nil {
		return nil, err
	}

	ffmpeg.AVFilterInoutFree(&outputs)
	ffmpeg.AVFilterInoutFree(&inputs)

	return &Decoder{
		fmtCtx:        fmtCtx,
		decCtx:        decCtx,
		vsi:           vsi,
		bufferSrcCtx:  bufferSrcCtx,
		bufferSinkCtx: bufferSinkCtx,
		timeBase:      timeBase,
	}, nil
}

type Frame struct {
	data   []byte
	pts    int64
	width  int
	height int
	ls     int
}

func (d *Decoder) Start(ctx context.Context, c chan<- *Frame) error {
	frame := ffmpeg.AVFrameAlloc()
	filtFrame := ffmpeg.AVFrameAlloc()
	packet := ffmpeg.AVPacketAlloc()

	for {
		if err := ctx.Err(); err != nil {
			return err
		}

		if _, err := ffmpeg.AVReadFrame(d.fmtCtx, packet); err != nil {
			return err
		}

		if packet.StreamIndex() == d.vsi {

			if _, err := ffmpeg.AVCodecSendPacket(d.decCtx, packet); err != nil {
				return err
			}

			for {
				_, err := ffmpeg.AVCodecReceiveFrame(d.decCtx, frame)
				if errors.Is(err, ffmpeg.EAgain) || errors.Is(err, ffmpeg.AVErrorEOF) {
					break
				}

				if err != nil {
					return err
				}

				frame.SetPts(frame.BestEffortTimestamp())

				if _, err := ffmpeg.AVBuffersrcAddFrameFlags(d.bufferSrcCtx, frame, ffmpeg.AVBuffersrcFlagKeepRef); err != nil {
					return err
				}

				for {
					_, err := ffmpeg.AVBuffersinkGetFrame(d.bufferSinkCtx, filtFrame)
					if errors.Is(err, ffmpeg.EAgain) || errors.Is(err, ffmpeg.AVErrorEOF) {
						break
					}

					if err != nil {
						return err
					}

					pts := int64(0)

					if filtFrame.Pts() != ffmpeg.AVNoptsValue {
						pts = ffmpeg.AVRescaleQ(filtFrame.Pts(), d.timeBase, ffmpeg.AVTimeBaseQ)
					}

					ptr := filtFrame.Data().Get(0)
					ls := filtFrame.Linesize().Get(0)

					data := unsafe.Slice((*byte)(ptr), filtFrame.Height()*ls)
					dataCopy := slices.Clone(data)

					c <- &Frame{
						data:   dataCopy,
						pts:    pts,
						width:  filtFrame.Width(),
						height: filtFrame.Height(),
						ls:     ls,
					}

					ffmpeg.AVFrameUnref(filtFrame)
				}

				ffmpeg.AVFrameUnref(frame)
			}

		}

		ffmpeg.AVPacketUnref(packet)
	}
}
