package scene

import (
	"battleship/pkg/game"
	"battleship/pkg/gui/compo"
	"battleship/pkg/gui/typing"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type StartScene struct {
	base
	bounds    pixel.Rect
	player    game.Player
	gotoSetup bool
	exit      bool
}

func NewStartScene(bounds pixel.Rect) *StartScene {
	scene := &StartScene{
		bounds: bounds,
	}
	midX := bounds.W() / 2
	btnCreate := compo.NewBlueButton(pixel.V(midX, 400), "CREATE")
	btnConnect := compo.NewBlueButton(pixel.V(midX, 300), "CONNECT")
	btnExit := compo.NewBlueButton(pixel.V(midX, 200), "EXIT")

	btnCreate.On(compo.Click, func(...any) {
		scene.gotoSetup = true
		scene.player = game.PlayerOne
	})
	btnConnect.On(compo.Click, func(...any) {
		scene.gotoSetup = true
		scene.player = game.PlayerTwo
	})
	btnExit.On(compo.Click, func(...any) {
		scene.exit = true
	})

	scene.components = []compo.Compo{
		compo.NewBg(bounds),
		compo.NewLabel(pixel.V(midX, 550), "BattleShip 2023", typing.Center, typing.Size78, colornames.Lightblue),
		compo.NewLabel(pixel.V(midX-3, 550+3), "BattleShip 2023", typing.Center, typing.Size78, colornames.Deepskyblue), // shadow effect
		btnCreate,
		btnConnect,
		btnExit,
	}
	return scene
}

func (s *StartScene) Update(win *pixelgl.Window, dt float64) Scene {
	for _, component := range s.components {
		component.Update(win, dt)
	}
	if s.exit {
		win.SetClosed(true)
		return nil
	}
	if s.gotoSetup {
		return NewSetupScene(s.bounds, s.player)
	}
	return nil
}
