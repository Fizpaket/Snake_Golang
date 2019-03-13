package main

import (
	"github.com/nsf/termbox-go"
	"strconv"
)

// Score
//--------------------------------------//
type score struct {
	x     int
	y     int
	count int
}

func newScore(x, y int) *score {
	s := &score{
		x: x,
		y: y,
	}
	return s
}

func (s *score) incr() {
	s.count += 10
}

func (s *score) reset() {
	s.count = 0
}

func (s *score) render() []*screenCell {
	cells := make([]*screenCell, 0)
	text := strconv.Itoa(s.count)
	for index, char := range text {
		cell := &screenCell{
			x:    s.x + index,
			y:    s.y,
			char: char,
			fg:   termbox.ColorGreen,
			bg:   termbox.ColorBlack,
		}
		cells = append(cells, cell)
	}
	return cells
}

// Screen text
//--------------------------------------//

type screenText struct {
	text string
	x    int
	y    int
}

func (s *screenText) render() []*screenCell {
	cells := make([]*screenCell, 0)
	for index, char := range s.text {
		cell := &screenCell{
			x:    s.x + index,
			y:    s.y,
			char: char,
			fg:   termbox.ColorGreen,
			bg:   termbox.ColorBlack,
		}
		cells = append(cells, cell)
	}
	return cells
}
