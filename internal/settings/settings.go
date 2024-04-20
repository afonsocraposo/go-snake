package settings

import (
    "github.com/gdamore/tcell"
)

type GameOver func()

type State int

const (
    PLAYING = iota
    GAME_OVER
)

type Settings struct {
    Screen tcell.Screen
    Padding int
    State
}

func (s *Settings) SqPadding() (int, int) {
    width, height := s.Screen.Size()
    playW, playH := width-s.Padding, height-s.Padding
    if playW < playH {
        return s.Padding, (height-playW)/2
    } else {
        return (width-playH)/2, s.Padding
    }
}

func (s *Settings) PlaySize() (int, int) {
    width, height := s.Screen.Size()
    paddingW, paddingH := s.SqPadding()
    return width-2*paddingW-2, height-2*paddingH-2
}
