package engine

import (
	"log"
	"time"

	"github.com/afonsocraposo/go-snake/internal/events"
	. "github.com/afonsocraposo/go-snake/internal/game"
	"github.com/gdamore/tcell"
)

const FPS = 15

type Engine struct {
    *Game
    s tcell.Screen
    quit chan struct{}
    GameEvents events.GameEventsChannel
}

func (engine *Engine) eventLoop() {
	defer close(engine.quit)
    defer close(engine.GameEvents)

	for {
		// Poll event
		ev := engine.s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			engine.s.Sync()
		case *tcell.EventKey:
			k := ev.Key()
			r := ev.Rune()
			if k == tcell.KeyCtrlC || r == 'q' {
				return
			}
            if k == tcell.KeyLeft || r == 'a' {
                engine.GameEvents <- events.MOVE_LEFT
            }
            if k == tcell.KeyRight || r == 's' {
                engine.GameEvents <- events.MOVE_RIGHT
            }
            if k == tcell.KeyDown || r == 'r' {
                engine.GameEvents <- events.MOVE_DOWN
            }
            if k == tcell.KeyUp || r == 'w' {
                engine.GameEvents <- events.MOVE_UP
            }
            if k == tcell.KeyEnter {
                engine.GameEvents <- events.ENTER
            }
		}
	}
}

func (engine *Engine) renderLoop() {
	ticker := time.NewTicker(time.Second / FPS)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			engine.Game.Update()
			engine.Game.Render()
		case <-engine.quit:
			return
		}
	}
}

func Start() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// Set default text style
	s.SetStyle(tcell.StyleDefault)

	// Clear screen
	s.Clear()

	quit := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	q := make(chan struct{})
    GameEvents := make(events.GameEventsChannel)
	game := NewGame(s, GameEvents)

    engine := Engine{game, s, q, GameEvents}

	go engine.renderLoop()

	engine.eventLoop()
}
