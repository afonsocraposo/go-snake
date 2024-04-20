package draw

import (
	"github.com/afonsocraposo/go-snake/internal/utils"
	"github.com/gdamore/tcell"
)

type Draw struct {
    S tcell.Screen
}

func (d *Draw) setContentWithDefaultStyle(x, y int, r rune) {
	d.S.SetContent(x, y, r, nil, tcell.StyleDefault)
}

func (d *Draw) Text(x, y int, text string) {
	row := y
	col := x
	for _, r := range []rune(text) {
		d.setContentWithDefaultStyle(col, row, r)
		col++
	}
}

func (d *Draw) Borders(ul utils.Point, ur utils.Point, lr utils.Point, ll utils.Point) {
	for x := ul.X+1; x < ur.X; x++ {
		d.setContentWithDefaultStyle(x, ul.Y, tcell.RuneHLine)
		d.setContentWithDefaultStyle(x, ll.Y, tcell.RuneHLine)
	}
	for y := ul.Y; y < ll.Y; y++ {
		d.setContentWithDefaultStyle(ul.X, y, tcell.RuneVLine)
		d.setContentWithDefaultStyle(ur.X, y, tcell.RuneVLine)
	}
	d.setContentWithDefaultStyle(ul.X, ul.Y, tcell.RuneULCorner)
	d.setContentWithDefaultStyle(ur.X, ur.Y, tcell.RuneURCorner)
	d.setContentWithDefaultStyle(lr.X, lr.Y, tcell.RuneLRCorner)
	d.setContentWithDefaultStyle(ll.X, ll.Y, tcell.RuneLLCorner)
}

