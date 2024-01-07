package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
)

type Display struct {
	s           tcell.Screen
	defStyle    tcell.Style
	boxStyle    tcell.Style
	pixelStyles []tcell.Style
}

func NewDisplay() *Display {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Panicln(err)
	}

	if err := s.Init(); err != nil {
		log.Panicln(err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	pixelStyles := make([]tcell.Style, 256)

	// Clear screen
	s.SetStyle(defStyle)
	s.Clear()

	nColors := s.Colors()
	if nColors > 256 {
		nColors = 256
	}

	palette := make([]tcell.Color, nColors)
	for i := 0; i < nColors; i++ {
		palette[i] = tcell.Color(i) | tcell.ColorValid
	}

	for i := 0; i < len(pixelStyles); i++ {
		col := tcell.FindColor(tcell.NewRGBColor(int32(i), int32(i), int32(i)), palette)
		pixelStyles[i] = tcell.StyleDefault.Background(col).Foreground(col)
	}

	return &Display{
		s:           s,
		defStyle:    defStyle,
		boxStyle:    boxStyle,
		pixelStyles: pixelStyles,
	}
}

func (d *Display) Close() {
	d.s.Fini()
}

func (d *Display) Size() (int, int) {
	return d.s.Size()
}

func (d *Display) ShouldExit() bool {
	if d.s.HasPendingEvent() {
		ev := d.s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			d.s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return true
			}
		}
	}

	return false
}

func (d *Display) DrawText(x1 int, y1 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		d.s.SetContent(col, row, r, nil, style)
		col++

	}
}

func (d *Display) SetPixel(x int, y int, col byte) {
	d.s.SetContent(x, y, ' ', nil, d.pixelStyles[col])
}

func (d *Display) Show() {
	d.s.Show()
}
