package main

import (
	"math/rand"
	"time"

	termbox "github.com/nsf/termbox-go"
)

var width, height int

var mx, my int

var status string

var staticBoard [][]bool

const maxFrameDuration = 1000
const minFrameDuration = 10

var frameDuration float64

func randCell() bool {
	//taken from https://github.com/pennello/go_util/blob/master/fix/math/rand/rand.go:
	return rand.Int63()&1 == 1
}

func newBoard() [][]bool {
	b := make([][]bool, height)
	for i := 0; i < height; i++ {
		b[i] = make([]bool, width)
	}
	return b
}

func randomBoard() [][]bool {
	b := newBoard()
	for y := range b {
		for x := range b[y] {
			b[y][x] = randCell()
		}
	}
	return b
}

func newNeighborBoard() [][]int {
	n := make([][]int, height)
	for i := 0; i < height; i++ {
		n[i] = make([]int, width)
	}
	return n
}

func getNeighborBoard(b [][]bool) [][]int {
	n := newNeighborBoard()
	for y := range b {
		isTopRow := y == 0
		isBottomRow := y == len(b)-1

		for x := range b[y] {
			isLeftCol := x == 0
			isRightCol := x == len(b[y])-1

			if !isLeftCol {
				//upper left
				if !isTopRow {
					if b[y-1][x-1] {
						n[y][x]++
					}
				}
				//left
				if b[y][x-1] {
					n[y][x]++
				}
				//bottom left
				if !isBottomRow {
					if b[y+1][x-1] {
						n[y][x]++
					}
				}
			}

			if !isTopRow {
				//top neighbor
				if b[y-1][x] {
					n[y][x]++
				}
			}

			if !isBottomRow {
				//bottom neighbor
				if b[y+1][x] {
					n[y][x]++
				}
			}

			if !isRightCol {
				//upper right
				if !isTopRow {
					if b[y-1][x+1] {
						n[y][x]++
					}
				}
				//right
				if b[y][x+1] {
					n[y][x]++
				}
				//bottom right
				if !isBottomRow {
					if b[y+1][x+1] {
						n[y][x]++
					}
				}
			}
		}
	}
	return n
}

func nextGeneration(b [][]bool) [][]bool {
	n := getNeighborBoard(b)
	newB := newBoard()
	for y := range b {
		for x := range b[y] {
			newB[y][x] = (b[y][x] && (n[y][x] == 2 || n[y][x] == 3)) ||
				(!b[y][x] && n[y][x] == 3)
		}
	}
	return newB
}

func drawBoard(b [][]bool) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := range b {
		for x := range b[y] {
			if b[y][x] {
				termbox.SetCell(x, y, ' ', termbox.ColorBlue, termbox.ColorBlue)
			}
		}
	}
	termbox.Flush()
}

func itsAlive(gameBoard [][]bool, events chan termbox.Event) {
	for {
		select {
		case ev := <-events:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeySpace:
					staticBoard = gameBoard
					return

				case termbox.KeyArrowUp:
					if frameDuration > minFrameDuration {
						frameDuration = frameDuration * .9
					} else {
						frameDuration = minFrameDuration
					}
				case termbox.KeyArrowDown:
					if frameDuration < maxFrameDuration {
						frameDuration = frameDuration * 1.1
					} else {
						frameDuration = maxFrameDuration
					}
				}

			}
		default:
			gameBoard = nextGeneration(gameBoard)
			drawBoard(gameBoard)
			time.Sleep(time.Duration(frameDuration) * time.Millisecond)
		}
	}
}

func toggleCell(x, y int) {
	staticBoard[y][x] = !staticBoard[y][x]
	drawBoard(staticBoard)
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	events := make(chan termbox.Event)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()

	rand.Seed(int64(time.Now().Nanosecond()))

	width, height = termbox.Size()

	frameDuration = 250

	staticBoard = randomBoard()
	drawBoard(staticBoard)

playGame:
	for {
		mx, my = -1, -1
		select {
		case ev := <-events:
			switch ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeySpace:
					itsAlive(staticBoard, events)
				case termbox.KeyEsc:
					break playGame
				case termbox.KeyDelete:
					staticBoard = newBoard()
					drawBoard(staticBoard)
				case termbox.KeyEnter:
					staticBoard = nextGeneration(staticBoard)
					drawBoard(staticBoard)
				}
			case termbox.EventMouse:
				if ev.Key == termbox.MouseLeft {
					mx, my = ev.MouseX, ev.MouseY
				}
			}
			if mx != -1 && my != -1 {
				toggleCell(mx, my)
			}
		}
	}
}
