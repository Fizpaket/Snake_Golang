package main

import (
	"github.com/nsf/termbox-go"
)

type screenCell struct {
	x    int
	y    int
	char rune
	fg   termbox.Attribute
	bg   termbox.Attribute
}

func draw(screen []*screenCell) {
	for _, cell := range screen {
		termbox.SetCell(cell.x, cell.y, cell.char, cell.fg, cell.bg)
	}
}

type renderer interface {
	render() []*screenCell
}

func drawObjects(m map[string]renderer) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	for _, object := range m {
		draw(object.render())
	}
	termbox.Flush()
}
