package object

import (
	"github.com/afonsocraposo/go-snake/internal/draw"
	"github.com/afonsocraposo/go-snake/internal/events"
	"github.com/afonsocraposo/go-snake/internal/settings"
)

type Object interface {
    Render(draw draw.Draw, settings *settings.Settings)
    Update(settings *settings.Settings, gameEvents events.GameEventsChannel)
    HandleGameEvent(settings *settings.Settings, event events.GameEvent)
}
