package scene

import (
	"battleship/pkg/gui"
	"battleship/pkg/gui/compo"
	"github.com/faiface/pixel/pixelgl"
)

type Scene interface {
	gui.Disposer
	Update(win *pixelgl.Window, dt float64) Scene
	Draw(win *pixelgl.Window, dt float64)
}

type base struct {
	components []compo.Compo
}

func (b *base) Draw(win *pixelgl.Window, dt float64) {
	for _, component := range b.components {
		component.Draw(win, dt)
	}
}

func (b *base) Dispose() error {
	for _, component := range b.components {
		if c, ok := component.(gui.Disposer); ok {
			_ = c.Dispose()
		}
	}
	return nil
}
