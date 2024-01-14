package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/csnewman/ffmpeg-go"
)

func main() {
	slog.Info("Metadata")

	var ctx *ffmpeg.AVFormatContext

	url := ffmpeg.ToCStr(os.Args[1])
	defer url.Free()

	_, err := ffmpeg.AVFormatOpenInput(&ctx, url, nil, nil)
	if err != nil {
		log.Panicln(err)
	}

	defer ffmpeg.AVFormatFreeContext(ctx)

	if _, err := ffmpeg.AVFormatFindStreamInfo(ctx, nil); err != nil {
		log.Panicln(err)
	}

	streams := ctx.Streams()

	for i := uintptr(0); i < uintptr(ctx.NbStreams()); i++ {
		s := streams.Get(i)

		slog.Info("  Stream", "i", i, "id", s.Id(), "dur", s.Duration())

		meta := s.Metadata()

		var entry *ffmpeg.AVDictionaryEntry

		for {
			entry = ffmpeg.AVDictGet(meta, ffmpeg.GlobalCStr(""), entry, ffmpeg.AVDictIgnoreSuffix)
			if entry == nil {
				break
			}

			slog.Info("    Meta", "key", entry.Key(), "value", entry.Value())
		}
	}
}
