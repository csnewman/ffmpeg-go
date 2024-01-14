// Port of doc/examples/transcode.c
package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/csnewman/ffmpeg-go"
)

func main() {
	slog.Info("Transcode")

	t := &Transcoder{}

	t.openIn()
	t.openOut()
	t.initFilters()

	packet := ffmpeg.AVPacketAlloc()

	inStreams := t.ifmtCtx.Streams()
	outStreams := t.ofmtCtx.Streams()

	for {
		if _, err := ffmpeg.AVReadFrame(t.ifmtCtx, packet); err != nil {
			if errors.Is(err, ffmpeg.AVErrorEOF) {
				log.Println("End of file")
				break
			}

			log.Panicln(err)
		}

		idx := packet.StreamIndex()

		stream := t.streams[idx]

		if stream.filterGraph != nil {

			if _, err := ffmpeg.AVCodecSendPacket(stream.decCtx, packet); err != nil {
				log.Panicln(err)
			}

			for {
				if _, err := ffmpeg.AVCodecReceiveFrame(stream.decCtx, stream.decFrame); err != nil {
					if errors.Is(err, ffmpeg.AVErrorEOF) || errors.Is(err, ffmpeg.EAgain) {
						break
					}

					log.Panicln(err)
				}

				stream.decFrame.SetPts(stream.decFrame.BestEffortTimestamp())

				stream.Write(
					stream.decFrame,
					outStreams.Get(uintptr(idx)).TimeBase(),
					t.ofmtCtx,
				)
			}

		} else {

			ffmpeg.AVPacketRescaleTs(
				packet,
				inStreams.Get(uintptr(idx)).TimeBase(),
				outStreams.Get(uintptr(idx)).TimeBase(),
			)

			if _, err := ffmpeg.AVInterleavedWriteFrame(t.ofmtCtx, packet); err != nil {
				log.Panicln(err)
			}
		}

		ffmpeg.AVPacketUnref(packet)
	}

	for idx, stream := range t.streams {
		stream.Flush(outStreams.Get(uintptr(idx)).TimeBase(), t.ofmtCtx)
	}

	if _, err := ffmpeg.AVWriteTrailer(t.ofmtCtx); err != nil {
		log.Panicln(err)
	}

	ffmpeg.AVPacketFree(&packet)
	for _, stream := range t.streams {
		stream.Free()
	}
	ffmpeg.AVFormatCloseInput(&t.ifmtCtx)

	if t.ofmtCtx.Oformat().Flags()&ffmpeg.AVFmtNofile == 0 {
		if _, err := ffmpeg.AVIOClose(t.ofmtCtx.Pb()); err != nil {
			log.Panicln(err)
		}

		t.ofmtCtx.SetPb(nil)
	}

	ffmpeg.AVFormatFreeContext(t.ofmtCtx)
}

type Transcoder struct {
	ifmtCtx *ffmpeg.AVFormatContext
	ofmtCtx *ffmpeg.AVFormatContext

	streams []*StreamContext
}

type StreamContext struct {
	id int

	decCtx *ffmpeg.AVCodecContext
	encCtx *ffmpeg.AVCodecContext

	decFrame      *ffmpeg.AVFrame
	encPkt        *ffmpeg.AVPacket
	filteredFrame *ffmpeg.AVFrame

	bufferSrcCtx  *ffmpeg.AVFilterContext
	bufferSinkCtx *ffmpeg.AVFilterContext
	filterGraph   *ffmpeg.AVFilterGraph
}

func (c *StreamContext) Write(frame *ffmpeg.AVFrame, outTb *ffmpeg.AVRational, ofmtCtx *ffmpeg.AVFormatContext) {

	if _, err := ffmpeg.AVBuffersrcAddFrameFlags(c.bufferSrcCtx, frame, 0); err != nil {
		log.Panicln(err)
	}

	for {
		if _, err := ffmpeg.AVBuffersinkGetFrame(c.bufferSinkCtx, c.filteredFrame); err != nil {
			if errors.Is(err, ffmpeg.EAgain) || errors.Is(err, ffmpeg.AVErrorEOF) {
				break
			}

			log.Panicln(err)
		}

		c.filteredFrame.SetTimeBase(ffmpeg.AVBuffersinkGetTimeBase(c.bufferSinkCtx))
		c.filteredFrame.SetPictType(ffmpeg.AVPictureTypeNone)

		c.EncodeWrite(false, outTb, ofmtCtx)

		ffmpeg.AVFrameUnref(c.filteredFrame)
	}

}

func (c *StreamContext) EncodeWrite(flush bool, outTb *ffmpeg.AVRational, ofmtCtx *ffmpeg.AVFormatContext) {
	filtFrame := c.filteredFrame
	if flush {
		filtFrame = nil
	}

	ffmpeg.AVPacketUnref(c.encPkt)

	if filtFrame != nil && filtFrame.Pts() != ffmpeg.AVNoptsValue {
		filtFrame.SetPts(
			ffmpeg.AVRescaleQ(filtFrame.Pts(), filtFrame.TimeBase(), c.encCtx.TimeBase()),
		)
	}

	if _, err := ffmpeg.AVCodecSendFrame(c.encCtx, filtFrame); err != nil {
		log.Panicln(err)
	}

	for {
		if _, err := ffmpeg.AVCodecReceivePacket(c.encCtx, c.encPkt); err != nil {
			if errors.Is(err, ffmpeg.EAgain) || errors.Is(err, ffmpeg.AVErrorEOF) {
				break
			}

			log.Panicln(err)
		}

		// prepare packet for muxing
		c.encPkt.SetStreamIndex(c.id)
		ffmpeg.AVPacketRescaleTs(c.encPkt, c.encCtx.TimeBase(), outTb)

		// mux encoded frame
		if _, err := ffmpeg.AVInterleavedWriteFrame(ofmtCtx, c.encPkt); err != nil {
			log.Panicln(err)
		}
	}

}

func (c *StreamContext) Flush(outTb *ffmpeg.AVRational, ofmtCtx *ffmpeg.AVFormatContext) {
	if c.filterGraph == nil {
		return
	}

	if _, err := ffmpeg.AVCodecSendPacket(c.decCtx, nil); err != nil {
		log.Panicln(err)
	}

	for {
		if _, err := ffmpeg.AVCodecReceiveFrame(c.decCtx, c.decFrame); err != nil {
			if errors.Is(err, ffmpeg.AVErrorEOF) {
				break
			}

			log.Panicln(err)
		}

		c.decFrame.SetPts(c.decFrame.BestEffortTimestamp())

		c.Write(c.decFrame, outTb, ofmtCtx)
	}

	c.Write(nil, outTb, ofmtCtx)

	if c.encCtx.Codec().Capabilities()&ffmpeg.AVCodecCapDelay != 0 {
		c.EncodeWrite(true, outTb, ofmtCtx)
	}
}

func (c *StreamContext) Free() {

	ffmpeg.AVCodecFreeContext(&c.decCtx)
	ffmpeg.AVCodecFreeContext(&c.encCtx)

	if c.filterGraph != nil {
		ffmpeg.AVFilterGraphFree(&c.filterGraph)
		ffmpeg.AVPacketFree(&c.encPkt)
		ffmpeg.AVFrameFree(&c.filteredFrame)
	}

	ffmpeg.AVFrameFree(&c.decFrame)

}

func (t *Transcoder) openIn() {
	urlPtr := ffmpeg.ToCStr(os.Args[1])
	defer urlPtr.Free()

	if _, err := ffmpeg.AVFormatOpenInput(&t.ifmtCtx, urlPtr, nil, nil); err != nil {
		log.Panicln(err)
	}

	if _, err := ffmpeg.AVFormatFindStreamInfo(t.ifmtCtx, nil); err != nil {
		log.Panicln(err)
	}

	slog.Info("streams", "nb", t.ifmtCtx.NbStreams())

	streams := t.ifmtCtx.Streams()

	for i := uintptr(0); i < uintptr(t.ifmtCtx.NbStreams()); i++ {
		slog.Info("Stream", "i", i)
		stream := streams.Get(i)

		cid := stream.Codecpar().CodecId()

		slog.Info(" > ", "codec", ffmpeg.AVCodecGetName(cid).String())

		dec := ffmpeg.AVCodecFindDecoder(cid)
		if dec == nil {
			log.Panicln("no decoder")
		}

		codecCtx := ffmpeg.AVCodecAllocContext3(dec)
		if codecCtx == nil {
			log.Panicln("failed to alloc context")
		}

		if _, err := ffmpeg.AVCodecParametersToContext(codecCtx, stream.Codecpar()); err != nil {
			log.Panicln(err)
		}

		// Inform the decoder about the timebase for the packet timestamps. This is highly recommended, but not
		// mandatory.
		codecCtx.SetPktTimebase(stream.TimeBase())

		cType := codecCtx.CodecType()

		slog.Info(" > ", "type", ffmpeg.AVGetMediaTypeString(cType).String())

		if cType == ffmpeg.AVMediaTypeVideo || cType == ffmpeg.AVMediaTypeAudio {

			if cType == ffmpeg.AVMediaTypeVideo {
				fr := ffmpeg.AVGuessFrameRate(t.ifmtCtx, stream, nil)

				slog.Info(" > ", "fr", fr)

				codecCtx.SetFramerate(fr)
			}

			if _, err := ffmpeg.AVCodecOpen2(codecCtx, dec, nil); err != nil {
				log.Panicln(err)
			}
		}

		frame := ffmpeg.AVFrameAlloc()

		t.streams = append(t.streams, &StreamContext{
			id:       int(i),
			decCtx:   codecCtx,
			encCtx:   nil,
			decFrame: frame,
		})
	}

	ffmpeg.AVDumpFormat(t.ifmtCtx, 0, urlPtr, 0)
}

func (t *Transcoder) openOut() {
	namePtr := ffmpeg.ToCStr(os.Args[2])
	defer namePtr.Free()

	if _, err := ffmpeg.AVFormatAllocOutputContext2(&t.ofmtCtx, nil, nil, namePtr); err != nil {
		log.Panicln(err)
	}

	slog.Info("streams", "nb", t.ifmtCtx.NbStreams())

	streams := t.ifmtCtx.Streams()

	for i := uintptr(0); i < uintptr(t.ifmtCtx.NbStreams()); i++ {
		slog.Info("Stream", "i", i)

		outStream := ffmpeg.AVFormatNewStream(t.ofmtCtx, nil)
		inStream := streams.Get(i)
		stream := t.streams[i]
		decCtx := stream.decCtx

		if decCtx.CodecType() == ffmpeg.AVMediaTypeVideo || decCtx.CodecType() == ffmpeg.AVMediaTypeAudio {
			// in this example, we choose transcoding to same codec
			encoder := ffmpeg.AVCodecFindEncoder(decCtx.CodecId())
			if encoder == nil {
				log.Panicln("failed to find encoder")
			}

			stream.encCtx = ffmpeg.AVCodecAllocContext3(encoder)
			encCtx := stream.encCtx

			// In this example, we transcode to same properties (picture size, sample rate etc.). These properties can
			// be changed for output streams easily using filters
			if decCtx.CodecType() == ffmpeg.AVMediaTypeVideo {
				encCtx.SetHeight(decCtx.Height())
				encCtx.SetWidth(decCtx.Width())
				encCtx.SetSampleAspectRatio(decCtx.SampleAspectRatio())

				fmts := encoder.PixFmts()
				if fmts != nil {
					encCtx.SetPixFmt(fmts.Get(0))
				} else {
					encCtx.SetPixFmt(decCtx.PixFmt())
				}

				encCtx.SetTimeBase(ffmpeg.AVInvQ(decCtx.Framerate()))
			} else {
				encCtx.SetSampleRate(decCtx.SampleRate())

				if _, err := ffmpeg.AVChannelLayoutCopy(encCtx.ChLayout(), decCtx.ChLayout()); err != nil {
					log.Panicln(err)
				}

				encCtx.SetSampleFmt(encoder.SampleFmts().Get(0))
				encCtx.SetTimeBase(ffmpeg.AVMakeQ(1, encCtx.SampleRate()))
			}

			if t.ofmtCtx.Oformat().Flags()&ffmpeg.AVFmtGlobalheader != 0 {
				encCtx.SetFlags(encCtx.Flags() | ffmpeg.AVCodecFlagGlobalHeader)
			}

			// Third parameter can be used to pass settings to encoder
			if _, err := ffmpeg.AVCodecOpen2(encCtx, encoder, nil); err != nil {
				log.Panicln(err)
			}

			if _, err := ffmpeg.AVCodecParametersFromContext(outStream.Codecpar(), encCtx); err != nil {
				log.Panicln(err)
			}

			outStream.SetTimeBase(encCtx.TimeBase())
		} else if decCtx.CodecType() == ffmpeg.AVMediaTypeUnknown {
			log.Panicln("unknown media type")
		} else {
			// if this stream must be remuxed
			if _, err := ffmpeg.AVCodecParametersCopy(outStream.Codecpar(), inStream.Codecpar()); err != nil {
				log.Panicln(err)
			}

			outStream.SetTimeBase(inStream.TimeBase())
		}
	}

	ffmpeg.AVDumpFormat(t.ofmtCtx, 0, namePtr, 1)

	if t.ofmtCtx.Oformat().Flags()&ffmpeg.AVFmtNofile == 0 {
		var pb *ffmpeg.AVIOContext

		if _, err := ffmpeg.AVIOOpen(&pb, namePtr, ffmpeg.AVIOFlagWrite); err != nil {
			log.Panicln(err)
		}

		t.ofmtCtx.SetPb(pb)
	}

	if _, err := ffmpeg.AVFormatWriteHeader(t.ofmtCtx, nil); err != nil {
		log.Panicln(err)
	}
}

func (t *Transcoder) initFilters() {

	for i, stream := range t.streams {
		slog.Info("Stream", "i", i)

		decCtx := stream.decCtx

		if decCtx.CodecType() != ffmpeg.AVMediaTypeVideo && decCtx.CodecType() != ffmpeg.AVMediaTypeAudio {
			continue
		}

		filterGraph := ffmpeg.AVFilterGraphAlloc()
		var bufferSrcCtx *ffmpeg.AVFilterContext
		var bufferSinkCtx *ffmpeg.AVFilterContext

		var filterSpec string

		if decCtx.CodecType() == ffmpeg.AVMediaTypeVideo {
			filterSpec = "null"

			bufferSrc := ffmpeg.AVFilterGetByName(ffmpeg.GlobalCStr("buffer"))
			bufferSink := ffmpeg.AVFilterGetByName(ffmpeg.GlobalCStr("buffersink"))

			if bufferSrc == nil || bufferSink == nil {
				log.Panicln("filtering src/sink not found")
			}

			pktTimebase := decCtx.PktTimebase()
			args := fmt.Sprintf(
				"video_size=%vx%v:pix_fmt=%v:time_base=%v/%v:pixel_aspect=%v/%v",
				decCtx.Width(), decCtx.Height(),
				decCtx.PixFmt(),
				pktTimebase.Num(), pktTimebase.Den(),
				decCtx.SampleAspectRatio().Num(),
				decCtx.SampleAspectRatio().Den(),
			)

			argsC := ffmpeg.ToCStr(args)
			defer argsC.Free()

			_, err := ffmpeg.AVFilterGraphCreateFilter(
				&bufferSrcCtx,
				bufferSrc,
				ffmpeg.GlobalCStr("in"),
				argsC,
				nil,
				filterGraph,
			)
			if err != nil {
				log.Panicln(err)
			}

			_, err = ffmpeg.AVFilterGraphCreateFilter(
				&bufferSinkCtx,
				bufferSink,
				ffmpeg.GlobalCStr("out"),
				nil,
				nil,
				filterGraph,
			)
			if err != nil {
				log.Panicln(err)
			}

			pixFmts := []ffmpeg.AVPixelFormat{
				stream.encCtx.PixFmt(),
			}

			_, err = ffmpeg.AVOptSetSlice(
				bufferSinkCtx.RawPtr(),
				ffmpeg.GlobalCStr("pix_fmts"),
				pixFmts,
				ffmpeg.AVOptSearchChildren,
			)
			if err != nil {
				log.Panicln(err)
			}

		} else {
			filterSpec = "anull"

			bufferSrc := ffmpeg.AVFilterGetByName(ffmpeg.GlobalCStr("abuffer"))
			bufferSink := ffmpeg.AVFilterGetByName(ffmpeg.GlobalCStr("abuffersink"))

			if bufferSrc == nil || bufferSink == nil {
				log.Panicln("filtering src/sink not found")
			}

			if decCtx.ChLayout().Order() == ffmpeg.AVChannelOrderUnspec {
				log.Println("unspec")
				ffmpeg.AVChannelLayoutDefault(decCtx.ChLayout(), decCtx.ChLayout().NbChannels())
			}

			layoutPtr := ffmpeg.AllocCStr(64)
			defer layoutPtr.Free()

			if _, err := ffmpeg.AVChannelLayoutDescribe(decCtx.ChLayout(), layoutPtr, 64); err != nil {
				log.Panicln(err)
			}

			layout := layoutPtr.String()

			log.Println("src layout", layout)

			pktTimebase := decCtx.PktTimebase()
			args := fmt.Sprintf(
				"time_base=%v/%v:sample_rate=%v:sample_fmt=%v:channel_layout=%v",
				pktTimebase.Num(), pktTimebase.Den(),
				decCtx.SampleRate(),
				ffmpeg.AVGetSampleFmtName(decCtx.SampleFmt()),
				layout,
			)

			log.Println(args)

			argsC := ffmpeg.ToCStr(args)
			defer argsC.Free()

			_, err := ffmpeg.AVFilterGraphCreateFilter(
				&bufferSrcCtx,
				bufferSrc,
				ffmpeg.GlobalCStr("in"),
				argsC,
				nil,
				filterGraph,
			)
			if err != nil {
				log.Panicln(err)
			}

			_, err = ffmpeg.AVFilterGraphCreateFilter(
				&bufferSinkCtx,
				bufferSink,
				ffmpeg.GlobalCStr("out"),
				nil,
				nil,
				filterGraph,
			)
			if err != nil {
				log.Panicln(err)
			}

			sampleFmts := []ffmpeg.AVSampleFormat{
				stream.encCtx.SampleFmt(),
			}

			_, err = ffmpeg.AVOptSetSlice(
				bufferSinkCtx.RawPtr(),
				ffmpeg.GlobalCStr("sample_fmts"),
				sampleFmts,
				ffmpeg.AVOptSearchChildren,
			)
			if err != nil {
				log.Panicln(err)
			}

			layoutPtr = ffmpeg.AllocCStr(64)
			defer layoutPtr.Free()

			if _, err := ffmpeg.AVChannelLayoutDescribe(stream.encCtx.ChLayout(), layoutPtr, 64); err != nil {
				log.Panicln(err)
			}

			layout = layoutPtr.String()

			log.Println("sink layout", layout)

			_, err = ffmpeg.AVOptSet(
				bufferSinkCtx.RawPtr(),
				ffmpeg.GlobalCStr("ch_layouts"),
				layoutPtr,
				ffmpeg.AVOptSearchChildren,
			)
			if err != nil {
				log.Panicln(err)
			}

			sampleRates := []int{
				stream.encCtx.SampleRate(),
			}

			_, err = ffmpeg.AVOptSetSlice(
				bufferSinkCtx.RawPtr(),
				ffmpeg.GlobalCStr("sample_rates"),
				sampleRates,
				ffmpeg.AVOptSearchChildren,
			)
			if err != nil {
				log.Panicln(err)
			}
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

		filterSpecC := ffmpeg.ToCStr(filterSpec)
		defer filterSpecC.Free()

		if _, err := ffmpeg.AVFilterGraphParsePtr(filterGraph, filterSpecC, &inputs, &outputs, nil); err != nil {
			log.Panicln(err)
		}

		if _, err := ffmpeg.AVFilterGraphConfig(filterGraph, nil); err != nil {
			log.Panicln(err)
		}

		stream.bufferSrcCtx = bufferSrcCtx
		stream.bufferSinkCtx = bufferSinkCtx
		stream.filterGraph = filterGraph

		ffmpeg.AVFilterInoutFree(&inputs)
		ffmpeg.AVFilterInoutFree(&outputs)

		stream.encPkt = ffmpeg.AVPacketAlloc()
		stream.filteredFrame = ffmpeg.AVFrameAlloc()
	}
}
