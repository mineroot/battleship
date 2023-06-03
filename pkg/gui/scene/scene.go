package scene

import "github.com/faiface/pixel/pixelgl"

type Scene interface {
	Update(win *pixelgl.Window, dt float64) Scene
	Draw(win *pixelgl.Window, dt float64)
}
