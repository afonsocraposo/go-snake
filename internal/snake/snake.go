package snake

import (
	"github.com/afonsocraposo/go-snake/internal/draw"
	"github.com/afonsocraposo/go-snake/internal/events"
	"github.com/afonsocraposo/go-snake/internal/settings"
	. "github.com/afonsocraposo/go-snake/internal/utils"
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Snake struct {
	Body          []Point
	velocity      Point
	readyToChange bool
    grow bool
}

func New(settings *settings.Settings) *Snake {
	playW, playH := settings.PlaySize()
	center := Point{X: playW/2 - 1, Y: playH/2 - 1}
	body := []Point{
		center,
		{X: center.X - 1, Y: center.Y},
		{X: center.X - 2, Y: center.Y},
	}
	snake := Snake{body, Point{X: 1, Y: 0}, true, false}
	return &snake
}

func (snake *Snake) Render(draw draw.Draw, settings *settings.Settings) {
    paddingW, paddingH := settings.SqPadding()
    head := snake.Body[0]
    draw.Text(head.X+paddingW+1, head.Y+paddingH+1, "@")
    for _, point := range snake.Body[1:] {
		draw.Text(point.X+paddingW+1, point.Y+paddingH+1, "#")
	}
}

func (snake *Snake) Update(settings *settings.Settings, gameEvents events.GameEventsChannel) {
    pw, _ := settings.PlaySize()
	s := map[int]bool{}
	for _, p := range snake.Body {
		n := p.Y*pw + p.X
		s[n] = true
	}

    if snake.grow {
        snake.Body = append(snake.Body[0:1], snake.Body...)
        snake.grow = false
    } else {
        snake.Body = append(snake.Body[0:1], snake.Body[:len(snake.Body)-1]...)
    }
	snake.Body[0].X += snake.velocity.X
	snake.Body[0].Y += snake.velocity.Y

	snake.readyToChange = true

	head := snake.Body[0]
    n := head.Y*pw + head.X
    _, dead := s[n]
	playW, playH := settings.PlaySize()
	if dead || head.X >= playW || head.X < 0 || head.Y >= playH || head.Y < 0 {
		gameEvents <- events.GAME_OVER
	}
}

func (snake *Snake) changeDirection(d Direction) {
	if !snake.readyToChange {
		return
	}
	switch d {
	case UP:
		if snake.velocity.Y != 0 {
			return
		}
		snake.velocity = Point{X: 0, Y: -1}
	case DOWN:
		if snake.velocity.Y != 0 {
			return
		}
		snake.velocity = Point{X: 0, Y: 1}
	case LEFT:
		if snake.velocity.X != 0 {
			return
		}
		snake.velocity = Point{X: -1, Y: 0}
	case RIGHT:
		if snake.velocity.X != 0 {
			return
		}
		snake.velocity = Point{X: 1, Y: 0}
	}
	snake.readyToChange = false
}

func (snake *Snake) HandleGameEvent(settings *settings.Settings, event events.GameEvent) {
	switch event {
    case events.EAT:
        snake.grow = true
	case events.MOVE_UP:
        snake.changeDirection(UP)
    case events.MOVE_RIGHT:
        snake.changeDirection(RIGHT)
    case events.MOVE_DOWN:
        snake.changeDirection(DOWN)
    case events.MOVE_LEFT:
        snake.changeDirection(LEFT)
	}
}
