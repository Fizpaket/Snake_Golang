package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

const (
	minWidth  = 35
	minHeight = 14
)

func main() {
	// Initialize termbox screen
	//-------------------------------//
	err := termbox.Init()
	if err != nil {
		fmt.Println("Error:", err)
	}
	width, height := termbox.Size()
	// Screen is not big enough for the game
	if width < minWidth || height < minHeight {
		termbox.Close()
		fmt.Printf("This game requires at least %v x %v terminal window to run.\n", minWidth, minHeight)
		fmt.Printf("Current window is %v x %v.\n", width, height)
		return
	}
	defer termbox.Close()
	//-------------------------------//

	// Initialize game objects
	//-------------------------------//
	screenObjects := make(map[string]renderer)

	arena := newArena(width, height)
	screenObjects["arena"] = arena
	screenObjects["outerBox"] = newOuterBox(width, height)

	title := &screenText{
		text: "SNAKE",
		x:    width/2 - 1,
		y:    height/2 - 6,
	}
	screenObjects["title"] = title

	scoreTxt := &screenText{
		text: "Score",
		x:    width/2 + 11,
		y:    height/2 - 3,
	}
	screenObjects["scoreTxt"] = scoreTxt

	bestScoreTxt := &screenText{
		text: "Best",
		x:    width/2 + 11,
		y:    height / 2,
	}
	screenObjects["bestScoreTxt"] = bestScoreTxt

	score := newScore(width/2+11, height/2-2)
	screenObjects["score"] = score

	bestScore := newScore(width/2+11, height/2+1)
	screenObjects["bestScore"] = bestScore
	//-------------------------------//

	// Poll for keyboard events
	//-------------------------------//
	eventChan := make(chan termbox.Event, 1)
	go func() {
		for {
			eventChan <- termbox.PollEvent()
		}
	}()
	//-------------------------------//

	// Startup window
	//-------------------------------//
	welcomeTxt := &screenText{text: "Welcome!", x: width/2 - 4, y: height/2 - 3}
	screenObjects["welcomeTxt"] = welcomeTxt

	enterTxt := &screenText{text: "Enter to start", x: width/2 - 7, y: height/2 - 1}
	screenObjects["enterTxt"] = enterTxt

	arrowsTxt := &screenText{text: "arrows to move", x: width/2 - 7, y: height / 2}
	screenObjects["arrowsTxt"] = arrowsTxt

	rTxt := &screenText{text: "R to restart", x: width/2 - 6, y: height/2 + 1}
	screenObjects["rTxt"] = rTxt

	qTxt := &screenText{text: "Q to quit", x: width/2 - 4, y: height/2 + 2}
	screenObjects["qTxt"] = qTxt

	drawObjects(screenObjects)

StartLoop:
	for {
		event := <-eventChan
		switch event.Key {
		// Start game
		case termbox.KeyEnter:
			break StartLoop
		default:
		}
		switch event.Ch {
		// Quit
		case 'q', 'Q':
			return
		default:
		}

	}
	delete(screenObjects, "welcomeTxt")
	delete(screenObjects, "enterTxt")
	delete(screenObjects, "arrowsTxt")
	delete(screenObjects, "rTxt")
	delete(screenObjects, "qTxt")
	snake := newSnake(width, height, arena)
	screenObjects["snake"] = snake
	//-------------------------------//

	// Game loop
	//-------------------------------//
	ticker := time.NewTicker(time.Millisecond * 200)
	restart := false
GameLoop:
	for {
		// Input Handling
		//-------------------------------//
		select {
		case event := <-eventChan:
			switch event.Key {
			case termbox.KeyArrowLeft:
				snake.changeDirection("left")
			case termbox.KeyArrowRight:
				snake.changeDirection("right")
			case termbox.KeyArrowUp:
				snake.changeDirection("up")
			case termbox.KeyArrowDown:
				snake.changeDirection("down")
			default:
			}
			switch event.Ch {
			// Quit
			case 'q', 'Q':
				break GameLoop
			// Restart
			case 'r', 'R':
				restart = true
			}
		default:
		}
		//-------------------------------//
		// Game actions
		//-------------------------------//
		snake.move()
		if snake.eat() {
			score.incr()
		}
		if snake.checkCollision() || restart {
			if score.count > bestScore.count {
				bestScore.count = score.count
			}
			snake.reset()
			score.reset()
			restart = false
		}
		drawObjects(screenObjects)
		<-ticker.C
		//-------------------------------//
	}
}
