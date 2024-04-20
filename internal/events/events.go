package events

type GameEvent int

const (
    GAME_OVER = iota
    MOVE_UP
    MOVE_DOWN
    MOVE_LEFT
    MOVE_RIGHT
    ENTER
    EAT
)

type GameEventsChannel chan GameEvent
