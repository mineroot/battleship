package sprites

import (
	"embed"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image"
	_ "image/png"
)

type CanvasType int
type tileCoordinates [2]int
type canvasLayout [][]tileCoordinates

const (
	ButtonYellowDefault CanvasType = iota
	ButtonYellowMouseDown
	ButtonBlueDefault
	ButtonBlueMouseDown
)

const (
	tileSize        = 16
	tileScale       = 4
	tileScaled      = tileSize * tileScale
	halfTileScaled  = tileScaled / 2
	tileMargin      = 2
	spriteSheetPath = "assets/sprite-sheet.png"
)

var (
	isInit          = false
	spriteSheet     pixel.Picture
	canvasesTiles   = make(map[CanvasType][][]*pixel.Sprite)
	canvasesLayouts = map[CanvasType]canvasLayout{
		ButtonYellowDefault: {
			[]tileCoordinates{{9, 30}, {10, 30}, {10, 30}, {11, 30}},
		},
		ButtonYellowMouseDown: {
			[]tileCoordinates{{9, 32}, {10, 32}, {10, 32}, {11, 32}},
		},
		ButtonBlueDefault: {
			[]tileCoordinates{{27, 30}, {28, 30}, {28, 30}, {29, 30}},
		},
		ButtonBlueMouseDown: {
			[]tileCoordinates{{27, 32}, {28, 32}, {28, 32}, {29, 32}},
		},
	}
)

func Init(assetsDir embed.FS) {
	loadSpriteSheet(assetsDir)
	initTiles()
	isInit = true
}

func CreateCanvas(canvasType CanvasType) *pixelgl.Canvas {
	checkInit()
	canvasTiles := canvasesTiles[canvasType]
	tilesHeight := len(canvasTiles)
	tilesWidth := len(canvasTiles[0])
	canvas := pixelgl.NewCanvas(pixel.R(0, 0, float64(tileScaled*tilesWidth), float64(tileScaled*tilesHeight)))
	for i, canvasColumnTiles := range canvasTiles {
		for j, tile := range canvasColumnTiles {
			reflectI := len(canvasTiles) - 1 - i // (0,0) top-left (as in layout) to bottom-left (as in game engine)
			x := float64(j*tileScaled + halfTileScaled)
			y := float64(reflectI*tileScaled + halfTileScaled)
			tile.Draw(canvas, pixel.IM.Scaled(pixel.ZV, tileScale).Moved(pixel.V(x, y)))
		}
	}
	return canvas
}

func checkInit() {
	if !isInit {
		panic(fmt.Errorf("sprites: not initialized"))
	}
}

func loadSpriteSheet(assetsDir embed.FS) {
	file, err := assetsDir.Open(spriteSheetPath)
	if err != nil {
		panic(fmt.Errorf("sprites: unable to open file: %w", err))
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(fmt.Errorf("sprites: unable to decode png: %w", err))
	}
	spriteSheet = pixel.PictureDataFromImage(img)
}

func initTiles() {
	for canvasType, layout := range canvasesLayouts {
		tiles := make([][]*pixel.Sprite, len(layout))
		if len(layout) == 0 {
			panic(fmt.Errorf("sprites: empty layout"))
		}
		layoutWidth := len(layout[0])
		for i, layoutColumn := range layout {
			if layoutWidth != len(layoutColumn) {
				panic(fmt.Errorf("sprites: wrong layout (must be rectangular)"))
			}
			tiles[i] = make([]*pixel.Sprite, len(layoutColumn))
			for j, coordinates := range layoutColumn {
				tiles[i][j] = getTile(coordinates[0], coordinates[1])
			}
		}
		canvasesTiles[canvasType] = tiles
	}
}

func getTile(x, y int) *pixel.Sprite {
	minX, minY := float64(x*tileSize+x*tileMargin), float64(y*tileSize+y*tileMargin)
	maxX, maxY := minX+tileSize, minY+tileSize

	return pixel.NewSprite(spriteSheet, pixel.R(minX, minY, maxX, maxY))
}
