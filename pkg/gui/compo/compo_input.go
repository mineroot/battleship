package compo

// TODO: move text to the left when text width greater than Input.Rect().W()
import (
	"battleship/pkg/gui/sprites"
	"battleship/pkg/gui/typing"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.design/x/clipboard"
	"golang.org/x/image/colornames"
	"image/color"
	"regexp"
	"time"
)

type Input struct {
	base
	RegexpPattern     string
	placeholder       string
	placeholderCanvas *pixelgl.Canvas
	caretCanvas       *pixelgl.Canvas
	caretIndex        int
	showCaret         bool
	focusOutChan      chan struct{}
	text              string
	padding           pixel.Vec
}

func NewInput(pos pixel.Vec, placeholder string) *Input {
	input := &Input{
		base:        newBase(sprites.CreateCanvas(sprites.Input, nil)),
		placeholder: placeholder,
	}
	input.Pos = pos
	input.Size = input.Canvas.Bounds().Size()
	input.placeholderCanvas = pixelgl.NewCanvas(input.Canvas.Bounds())
	input.caretCanvas = pixelgl.NewCanvas(input.Canvas.Bounds())
	typing.TypeOnCanvas(input.placeholderCanvas, input.placeholder, typing.Left, pixel.V(30, 20), typing.Size26, colornames.Lightslategrey)
	input.padding = pixel.V(20, 20)
	input.focusOutChan = make(chan struct{})

	caret := imdraw.New(nil)
	caret.Color = color.RGBA{R: 20, G: 167, B: 225, A: 255} // shape of blue
	caret.Push(pixel.V(20, 16), pixel.V(23, 40))
	caret.Rectangle(0)
	caret.Draw(input.caretCanvas)

	input.On(FocusIn, func(data ...any) {
		go func() {
			caretTicker := time.NewTicker(time.Millisecond * 500)
			defer caretTicker.Stop()
			input.showCaret = true
			for {
				select {
				case <-input.focusOutChan:
					input.showCaret = false
					return
				case <-caretTicker.C:
					input.showCaret = !input.showCaret
				}
			}
		}()
	})
	input.On(FocusOut, func(data ...any) {
		input.focusOutChan <- struct{}{}
	})
	input.On(Click, func(data ...any) {
		relativeMousePos := data[1].(pixel.Vec)
		relativeWithoutPadding := relativeMousePos.X - input.padding.X
		i := 0
		for ; i < len(input.text); i++ {
			textBounds := input.textBoundsOfLen(i)
			diff := textBounds.W() - relativeWithoutPadding
			if diff > 0 {
				break
			}
		}
		input.caretIndex = i
	})

	return input
}

func (i *Input) textBoundsOfLen(length int) pixel.Rect {
	return typing.BoundsOf(i.text[:length], i.padding, typing.Size26)
}

func (i *Input) Update(win *pixelgl.Window, dt float64) {
	i.base.Update(win, dt)
	if !i.isFocus {
		return
	}

	controlPressed := win.Pressed(pixelgl.KeyLeftControl) || win.Pressed(pixelgl.KeyRightControl)
	vPressed := win.JustPressed(pixelgl.KeyV) || win.Repeated(pixelgl.KeyV)
	var typed string
	if controlPressed && vPressed {
		copied := clipboard.Read(clipboard.FmtText)
		typed = string(copied)
	} else {
		typed = win.Typed()
	}

	if typed != "" {
		tmpText := i.text[:i.caretIndex] + typed + i.text[i.caretIndex:]
		ok, err := regexp.MatchString(i.RegexpPattern, tmpText)
		if err != nil {
			panic(fmt.Errorf("compo: wrong regexp pattern: %w", err))
		}
		if ok {
			i.text = tmpText
			i.caretIndex += len(typed)
		}
	}

	backSpacePressed := win.JustPressed(pixelgl.KeyBackspace) || win.Repeated(pixelgl.KeyBackspace)
	if backSpacePressed && len(i.text) != 0 {
		if i.caretIndex > 0 && i.caretIndex != len(i.text)+1 {
			i.text = i.text[:i.caretIndex-1] + i.text[i.caretIndex:]
			i.caretIndex -= 1
		}
	}
	deletePressed := win.JustPressed(pixelgl.KeyDelete) || win.Repeated(pixelgl.KeyDelete)
	if deletePressed && len(i.text) != 0 {
		if i.caretIndex != len(i.text) {
			i.text = i.text[:i.caretIndex] + i.text[i.caretIndex+1:]
		}
	}
	if win.JustPressed(pixelgl.KeyLeft) || win.Repeated(pixelgl.KeyLeft) {
		if i.caretIndex != 0 {
			i.caretIndex -= 1
		}
	}
	if win.JustPressed(pixelgl.KeyRight) || win.Repeated(pixelgl.KeyRight) {
		if i.caretIndex != len(i.text) {
			i.caretIndex += 1
		}
	}
	if win.JustReleased(pixelgl.KeyHome) {
		i.caretIndex = 0
	}
	if win.JustReleased(pixelgl.KeyEnd) {
		i.caretIndex = len(i.text)
	}
}

func (i *Input) Draw(target pixel.ComposeTarget, dt float64) {
	i.base.Draw(target, dt)
	if i.showCaret {
		bounds := i.textBoundsOfLen(i.caretIndex)
		caretPos := i.Pos.Add(pixel.V(bounds.W(), 0))
		i.caretCanvas.Draw(target, pixel.IM.Moved(caretPos))
	}
	if i.text == "" {
		i.placeholderCanvas.Draw(target, pixel.IM.Moved(i.Pos))
	} else {
		textCanvas := pixelgl.NewCanvas(i.Canvas.Bounds())
		typing.TypeOnCanvas(textCanvas, i.text, typing.Left, i.padding, typing.Size26, colornames.Gray)
		textCanvas.Draw(target, pixel.IM.Moved(i.Pos))
	}
}

func (i *Input) Dispose() error {
	close(i.focusOutChan)
	return nil
}
