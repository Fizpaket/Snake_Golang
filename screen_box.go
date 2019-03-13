package main

import (
	"github.com/nsf/termbox-go"
)

// Экраны
//----------------------------------------//
func newArena(width, height int) *box {
	arena := &box{
		upperLeftX:  width/2 - 10,
		upperLeftY:  height/2 - 5,
		lowerRightX: width/2 + 10,
		lowerRightY: height/2 + 5,
	}
	return arena
}

func newOuterBox(width, height int) *box {
	outerBox := &box{
		upperLeftX:  width/2 - 12,
		upperLeftY:  height/2 - 7,
		lowerRightX: width/2 + 16,
		lowerRightY: height/2 + 6,
	}
	return outerBox
}

type box struct {
	upperLeftX  int
	upperLeftY  int
	lowerRightX int
	lowerRightY int
}

func (b *box) render() []*screenCell {
	fg := termbox.ColorGreen
	bg := termbox.ColorBlack
	cells := make([]*screenCell, 0)
	for i := b.upperLeftX + 1; i < b.lowerRightX; i++ {
		cell := &screenCell{
			x:    i,
			y:    b.upperLeftY,
			char: '―',
			fg:   fg,
			bg:   bg,
		}
		cells = append(cells, cell)
	}
	for i := b.upperLeftX + 1; i < b.lowerRightX; i++ {
		cell := &screenCell{
			x:    i,
			y:    b.lowerRightY,
			char: '―',
			fg:   fg,
			bg:   bg,
		}
		cells = append(cells, cell)
	}
	for i := b.upperLeftY + 1; i < b.lowerRightY; i++ {
		cell := &screenCell{
			x:    b.upperLeftX,
			y:    i,
			char: '|',
			fg:   fg,
			bg:   bg,
		}
		cells = append(cells, cell)
	}
	for i := b.upperLeftY + 1; i < b.lowerRightY; i++ {
		cell := &screenCell{
			x:    b.lowerRightX,
			y:    i,
			char: '|',
			fg:   fg,
			bg:   bg,
		}
		cells = append(cells, cell)
	}
	cornCellLU := &screenCell{
		x:    b.upperLeftX,
		y:    b.upperLeftY,
		char: '+',
		fg:   fg,
		bg:   bg,
	}
	cells = append(cells, cornCellLU)
	cornCellRU := &screenCell{
		x:    b.lowerRightX,
		y:    b.upperLeftY,
		char: '+',
		fg:   fg,
		bg:   bg,
	}
	cells = append(cells, cornCellRU)
	cornCellLD := &screenCell{
		x:    b.upperLeftX,
		y:    b.lowerRightY,
		char: '+',
		fg:   fg,
		bg:   bg,
	}
	cells = append(cells, cornCellLD)

	cornCellRD := &screenCell{
		x:    b.lowerRightX,
		y:    b.lowerRightY,
		char: '+',
		fg:   fg,
		bg:   bg,
	}
	cells = append(cells, cornCellRD)

	return cells
}
