package game

import (
	"fmt"

	"github.com/afonsocraposo/go-snake/internal/draw"
	"github.com/afonsocraposo/go-snake/internal/events"
	"github.com/afonsocraposo/go-snake/internal/food"
	"github.com/afonsocraposo/go-snake/internal/object"
	"github.com/afonsocraposo/go-snake/internal/settings"
	"github.com/afonsocraposo/go-snake/internal/snake"
	"github.com/afonsocraposo/go-snake/internal/utils"
	"github.com/gdamore/tcell"
)

const PADDING = 0

type GameObject int

const (
	SNAKE GameObject = iota
	FOOD
)

type GameObjects map[GameObject]object.
	Object

type Game struct {
	draw.Draw
	settings.Settings
	objects GameObjects
	score   int
	events  events.GameEventsChannel
}

func NewGame(s tcell.Screen, gameEvents events.GameEventsChannel) *Game {
	settings := settings.Settings{
		Screen:  s,
		Padding: PADDING,
		State:   settings.PLAYING,
	}

	game := Game{}
	game.Draw = draw.Draw{S: s}
	game.Settings = settings
	game.score = 0
	game.events = gameEvents
	game.objects = game.newGameObjects()

	go game.handleGameEvents(gameEvents)
	return &game
}

func (g *Game) drawScore() {
    paddingW, paddingH := g.Settings.SqPadding()
	g.Draw.Text(paddingW+2, paddingH, fmt.Sprintf(" Score: %d ", g.score))
}

func (g *Game) hud() {
	width, height := g.S.Size()
    paddingW, paddingH := g.Settings.SqPadding()
	tl := utils.Point{X: paddingW, Y: paddingH}
	tr := utils.Point{X: width - paddingW - 1, Y: paddingH}
	lr := utils.Point{X: width - paddingW - 1, Y: height - paddingH - 1}
	ll := utils.Point{X: paddingW, Y: height - paddingH - 1}
	g.Draw.Borders(tl, tr, lr, ll)
}

func (g *Game) Render() {
	g.S.Clear()
	g.hud()
	g.drawScore()

	for _, obj := range g.objects {
		obj.Render(g.Draw, &g.Settings)
	}

	if g.State == settings.GAME_OVER {
		g.drawGameOver()
	}
	g.S.Show()
}

func (g *Game) drawGameOver() {
	width, height := g.S.Size()
	g.Draw.Text(width/2-4, height/2, "Game Over")
}
func (g *Game) handleGameEvents(ch chan events.GameEvent) {
	for event := range ch {
		switch event {
		case events.GAME_OVER:
			g.Settings.State = settings.GAME_OVER
		case events.ENTER:
			if g.State == settings.GAME_OVER {
				g.restart()
			}
		case events.EAT:
			g.score++
		}
		for _, obj := range g.objects {
			obj.HandleGameEvent(&g.Settings, event)
		}
	}
}

func (g *Game) Update() {
	if g.Settings.State == settings.PLAYING {
        // we need to ensure the snake positions are updated before we check the
        // collision between food and snake
		for _, gameObj := range []GameObject{SNAKE, FOOD} {
            obj := g.objects[gameObj]
			obj.Update(&g.Settings, g.events)
		}
	}
}

func (g *Game) newGameObjects() GameObjects {
	snake := snake.New(&g.Settings)
	food := food.New(snake)
	food.NewPosition(&g.Settings)
	return GameObjects{
		SNAKE: snake,
		FOOD:  food,
	}
}

func (g *Game) restart() {
	g.objects = g.newGameObjects()
	g.State = settings.PLAYING
    g.score = 0
}
