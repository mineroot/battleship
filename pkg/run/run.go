package run

import (
	"battleship/pkg/gui/scene"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"log"
	"time"
)

type runner struct {
	currentScene scene.Scene
}

func (r *runner) run(win *pixelgl.Window, dt float64) {
	newScene := r.currentScene.Update(win, dt)
	if newScene != nil {
		_ = r.currentScene.Dispose()
		r.currentScene = newScene
		return
	}
	r.currentScene.Draw(win, dt)
}

func Run() {
	bounds := pixel.R(0, 0, 1250, 700)
	cfg := pixelgl.WindowConfig{
		Title:     "Battleship 2023",
		Bounds:    bounds,
		Resizable: false,
	}
	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		log.Fatalln(err)
	}
	r := runner{currentScene: scene.NewStartScene(bounds)}
	const targetFPS = 60
	frameDuration := time.Second / targetFPS
	frameTicker := time.NewTicker(frameDuration)
	defer frameTicker.Stop()

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		r.run(win, dt)
		win.Update()

		<-frameTicker.C
	}
}
