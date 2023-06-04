package scene

import (
	"battleship/pkg/game"
	"battleship/pkg/gui/compo"
	"battleship/pkg/p2p"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type SetupScene struct {
	components []compo.Compo
	bounds     pixel.Rect
}

func NewSetupScene(bounds pixel.Rect, player game.Player) *SetupScene {
	if player == game.PlayerOne {
		go p2p.CreateServer()
	} else {
		go p2p.CreateClient("127.0.0.1:49152") //TODO hardcode!
	}

	newGame := game.NewGame()
	_ = newGame
	btnReady := compo.NewYellowButton(pixel.V(300, 100), "READY")
	btnReady.On(compo.Click, func(data ...any) {
		fmt.Println("ready")
	})
	components := []compo.Compo{
		compo.NewBg(bounds),
		compo.NewField(),
		btnReady,
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
