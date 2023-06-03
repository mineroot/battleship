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
	//newGame := game.NewGame()
	//_ = newGame
	components := []compo.Compo{
		compo.NewBg(bounds),
		compo.NewField(),
		compo.NewShip(1, pixel.V(775, 125)),
		compo.NewShip(1, pixel.V(875, 125)),
		compo.NewShip(1, pixel.V(975, 125)),
		compo.NewShip(1, pixel.V(1075, 125)),
		compo.NewShip(2, pixel.V(800, 225)),
		compo.NewShip(2, pixel.V(950, 225)),
		compo.NewShip(2, pixel.V(1100, 225)),
		compo.NewShip(3, pixel.V(825, 325)),
		compo.NewShip(3, pixel.V(1025, 325)),
		compo.NewShip(4, pixel.V(850, 425)),
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
