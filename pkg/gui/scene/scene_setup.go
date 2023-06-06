package scene

import (
	"battleship/pkg/game"
	f "battleship/pkg/game/field"
	"battleship/pkg/gui/compo"
	"battleship/pkg/gui/typing"
	"battleship/pkg/p2p"
	"errors"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type SetupScene struct {
	base
	bounds pixel.Rect
}

func NewSetupScene(bounds pixel.Rect, player game.Player) *SetupScene {
	scene := &SetupScene{}
	var inputOrLabel compo.Compo
	if player == game.PlayerOne {
		go p2p.CreateServer()
		label := compo.NewLabel(pixel.V(300, 100), "Copy your address 127.0.0.1:49152", typing.Left, typing.Size26, colornames.Blue)
		inputOrLabel = label
	} else {
		go p2p.CreateClient("127.0.0.1:49152") //TODO hardcode!
		input := compo.NewInput(pixel.V(300, 100), "Enter IP address")
		input.RegexpPattern = `^[\d\.:]{1,21}$`
		inputOrLabel = input
	}

	newGame := game.NewGame()
	_ = newGame

	field := compo.NewField()
	statusLbl := compo.NewLabel(pixel.V(300, 50), "", typing.Left, typing.Size39, colornames.Darkorange)
	btnReady := compo.NewYellowButton(pixel.V(1050, 100), "READY")
	btnReady.On(compo.Click, func(data ...any) {
		err := field.Validate()
		switch {
		case err == nil:
			statusLbl.SetCaption("Waiting for another player")
		case errors.Is(err, f.ErrVariantCount):
			statusLbl.SetCaption("Not all ships on field")
		case errors.Is(err, f.ErrShipOutOfField):
			statusLbl.SetCaption("Ship out of field")
		case errors.Is(err, f.ErrShipsOverlap):
			statusLbl.SetCaption("Ships overlap")
		default:
			statusLbl.SetCaption("Something wrong")
		}
	})
	scene.components = []compo.Compo{
		compo.NewBg(bounds),
		field,
		btnReady,
		statusLbl,
		inputOrLabel,
	}
	return scene
}

func (s *SetupScene) Update(win *pixelgl.Window, dt float64) Scene {
	for _, component := range s.components {
		component.Update(win, dt)
	}
	return nil
}
