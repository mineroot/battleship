package main

import (
	"battleship/pkg/gui/sprites"
	"battleship/pkg/run"
	"embed"
	"github.com/faiface/pixel/pixelgl"
)

//go:embed assets
var assetsDir embed.FS

func main() {
	sprites.Init(assetsDir)
	pixelgl.Run(run.Run)
}
