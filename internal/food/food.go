package food

import (
	"math/rand"

	"github.com/afonsocraposo/go-snake/internal/draw"
	"github.com/afonsocraposo/go-snake/internal/events"
	"github.com/afonsocraposo/go-snake/internal/settings"
	"github.com/afonsocraposo/go-snake/internal/snake"
	"github.com/afonsocraposo/go-snake/internal/utils"
)

type Food struct {
	position utils.Point
	*snake.Snake
}

func New(snake *snake.Snake) *Food {
	f := &Food{}
	f.Snake = snake
	return f
}

func (f *Food) Render(draw draw.Draw, settings *settings.Settings) {
    paddingW, paddingH := settings.SqPadding()
	draw.Text(f.position.X+paddingW+1, f.position.Y+paddingH+1, "O")
}
func (f *Food) Update(settings *settings.Settings, gameEvents events.GameEventsChannel) {
	head := f.Snake.Body[0]
	if head.X == f.position.X && head.Y == f.position.Y {
		gameEvents <- events.EAT
	}
}

func (f *Food) HandleGameEvent(settings *settings.Settings, event events.GameEvent) {
	switch event {
	case events.EAT:
		f.NewPosition(settings)
	}
}

func (f *Food) NewPosition(settings *settings.Settings) {
	pw, ph := settings.PlaySize()
	s := map[int]bool{}
	for _, p := range f.Snake.Body {
		n := p.Y*pw + p.X
		s[n] = true
	}
	free := []utils.Point{}
	for x := 0; x < pw; x++ {
		for y := 0; y < ph; y++ {
			n := y*pw + x
			if _, ok := s[n]; !ok {
				free = append(free, utils.Point{X: x, Y: y})
			}
		}
	}
	pick := rand.Intn(len(free))
	f.position = free[pick]
}
