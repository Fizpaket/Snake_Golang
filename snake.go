package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type snake struct {
	head      *screenCell
	body      []*screenCell
	direction string
	food      *screenCell
	width     int
	height    int
	arena     *box
}

func newSnake(width, height int, arena *box) *snake {
	head := &screenCell{
		x:    width / 2,
		y:    height / 2,
		char: '@',
		fg:   termbox.ColorMagenta,
		bg:   termbox.ColorBlack,
	}
	body := &screenCell{
		x:    width/2 - 1,
		y:    height / 2,
		char: '+',
		fg:   termbox.ColorYellow,
		bg:   termbox.ColorBlack,
	}
	s := &snake{
		height:    height,
		width:     width,
		head:      head,
		body:      make([]*screenCell, 0),
		direction: "right",
		arena:     arena,
	}

	s.body = append(s.body, body)
	s.generateFood()

	return s
}

func (s *snake) generateFood() {
	rand.Seed(time.Now().UnixNano())
	var coords []*struct {
		x int
		y int
	}
	for x := s.arena.upperLeftX + 1; x < s.arena.lowerRightX; x++ {
		for y := s.arena.upperLeftY + 1; y < s.arena.lowerRightY; y++ {
			suitable := true
			for _, body := range s.body {
				if x == body.x && y == body.y {
					suitable = false
				}
			}
			if x == s.head.x && y == s.head.y {
				suitable = false
			}
			if suitable {
				coord := &struct {
					x int
					y int
				}{
					x: x,
					y: y,
				}
				coords = append(coords, coord)
			}
		}

	}
	target := coords[rand.Intn(len(coords))]
	s.food = &screenCell{
		x:    target.x,
		y:    target.y,
		char: '$',
		fg:   termbox.ColorWhite,
		bg:   termbox.ColorBlack,
	}
}

func (s *snake) render() []*screenCell {
	cells := make([]*screenCell, 0, 10)
	cells = append(cells, s.head)
	cells = append(cells, s.food)
	cells = append(cells, s.body...)

	return cells
}

func (s *snake) checkCollision() bool {
	for _, cell := range s.body {
		if s.head.x == cell.x && s.head.y == cell.y {
			return true
		}
		if s.head.x <= s.arena.upperLeftX ||
			s.head.x >= s.arena.lowerRightX ||
			s.head.y <= s.arena.upperLeftY ||
			s.head.y >= s.arena.lowerRightY {
			return true
		}
	}
	return false
}

func (s *snake) changeDirection(dir string) {
	if dir == "left" && s.direction != "right" {
		s.direction = "left"
	}
	if dir == "right" && s.direction != "left" {
		s.direction = "right"
	}
	if dir == "up" && s.direction != "down" {
		s.direction = "up"
	}
	if dir == "down" && s.direction != "up" {
		s.direction = "down"
	}
}

func (s *snake) eat() bool {
	if s.head.x == s.food.x && s.head.y == s.food.y {
		body := &screenCell{
			char: '+',
			fg:   termbox.ColorYellow,
			bg:   termbox.ColorBlack,
			x:    -1,
			y:    -1,
		}
		s.body = append(s.body, body)
		s.generateFood()
		return true
	}
	return false
}

func (s *snake) move() {
	if len(s.body) > 1 {
		for i := len(s.body); i > 1; i-- {
			s.body[i-1].x = s.body[i-2].x
			s.body[i-1].y = s.body[i-2].y
		}
	}
	s.body[0].x = s.head.x
	s.body[0].y = s.head.y

	if s.direction == "right" {
		s.head.x++
	}
	if s.direction == "left" {
		s.head.x--
	}
	if s.direction == "up" {
		s.head.y--
	}
	if s.direction == "down" {
		s.head.y++
	}
}

func (s *snake) reset() {
	s.head.x = s.width / 2
	s.head.y = s.height / 2

	body := &screenCell{
		x:    s.width/2 - 1,
		y:    s.height / 2,
		char: '+',
		fg:   termbox.ColorYellow,
		bg:   termbox.ColorBlack,
	}
	s.body = make([]*screenCell, 0)
	s.body = append(s.body, body)
	s.direction = "right"
	s.generateFood()
}
