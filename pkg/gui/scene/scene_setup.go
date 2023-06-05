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

	field := compo.NewField()
	statusLbl := compo.NewLabel(pixel.V(300, 100), "", typing.Left, typing.Size39, colornames.Darkorange)
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
	components := []compo.Compo{
		compo.NewBg(bounds),
		field,
		btnReady,
		statusLbl,
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
