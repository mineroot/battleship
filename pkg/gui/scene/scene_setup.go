package scene

import (
	"battleship/pkg/game"
	"battleship/pkg/gui/compo"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type SetupScene struct {
	components []compo.Compo
	bounds     pixel.Rect
}

func NewSetupScene(bounds pixel.Rect, player game.Player) *SetupScene {
	newGame := game.NewGame()
	_ = newGame
	components := []compo.Compo{
		compo.NewBg(bounds),
		compo.NewField(),
	}
	return &SetupScene{
		components: components,
	}
}

func (s *SetupScene) Update(win *pixelgl.Window, dt float64) Scene {
	for _, component := range s.components {
		component.Update(win, dt)
	}
	return nil
}

func (s *SetupScene) Draw(win *pixelgl.Window, dt float64) {
	for _, component := range s.components {
		component.Draw(win, dt)
	}
}
