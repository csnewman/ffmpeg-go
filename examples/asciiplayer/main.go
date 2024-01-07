package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	d := NewDisplay()
	defer d.Close()

	d.Show()

	w, h := d.Size()

	log.Println(w, h)

	e, err := Open(os.Args[1], w, h)
	if err != nil {
		log.Panicln(err)
	}

	ch := make(chan *Frame)

	ctx := context.Background()

	go func() {
		err := e.Start(ctx, ch)
		if err != nil {
			ch <- nil

			log.Println(err)
		}
	}()

	fc := 0

	lastPts := int64(0)

	var lastPresent time.Time

	for {
		if d.ShouldExit() {
			return
		}

		f := <-ch
		if f == nil {
			return
		}

		ptsDiff := f.pts - lastPts
		now := time.Now()

		tgt := lastPresent.Add(time.Microsecond * time.Duration(ptsDiff))

		sleep := tgt.Sub(now)
		if sleep.Milliseconds() > 0 {
			time.Sleep(sleep)
		}

		lastPts = f.pts
		lastPresent = now

		pos := 0

		for y := 0; y < f.height; y++ {
			for x := 0; x < f.width; x++ {
				b := f.data[pos+x]

				d.SetPixel(x, y, b)
			}
			pos += f.ls
		}

		d.DrawText(0, 0, d.boxStyle, fmt.Sprintf("Press ESC to exit - frame=%v sleep=%v w=%v h=%v", fc, sleep, w, h))
		d.Show()

		fc++
	}

}
