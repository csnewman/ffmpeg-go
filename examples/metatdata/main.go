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

	empty := ffmpeg.ToCStr("")
	defer empty.Free()

	for i := uint(0); i < ctx.NbStreams(); i++ {
		s := ctx.StreamsEntry(i)

		slog.Info("  Stream", "i", i, "id", s.Id(), "dur", s.Duration())

		meta := s.Metadata()

		var entry *ffmpeg.AVDictionaryEntry

		for {
			entry = ffmpeg.AVDictGet(meta, empty, entry, ffmpeg.AVDictIgnoreSuffix)
			if entry == nil {
				break
			}

			slog.Info("    Meta", "key", entry.Key(), "value", entry.Value())
		}
	}
}
