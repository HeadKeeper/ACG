package lab3

import (
	"github.com/veandco/go-sdl2/sdl"
	"main/types"
	"main/utils"
)

var (
	WINDOW = types.Shape {

		Lines: []types.Line {
			{
				BeginPos: types.Point {
					X: 50,
					Y: 50,
				},
				EndPos: types.Point {
					X: 150,
					Y: 50,
				},
			},
			{
				BeginPos: types.Point {
					X: 150,
					Y: 50,
				},
				EndPos: types.Point {
					X: 150,
					Y: 100,
				},
			},
			{
				BeginPos: types.Point {
					X: 150,
					Y: 100,
				},
				EndPos: types.Point {
					X: 50,
					Y: 100,
				},
			},
			{
				BeginPos: types.Point {
					X: 50,
					Y: 100,
				},
				EndPos: types.Point {
					X: 50,
					Y: 50,
				},
			},
		},
		Color: sdl.Color { 228, 148, 8 , 0 },
	}

	EXAMPLE_SHAPE = types.Shape {
		Lines: []types.Line {
			{
				BeginPos: types.Point {
					X: 30,
					Y: 30,
				},
				EndPos: types.Point {
					X: 130,
					Y: 30,
				},
			},
			{
				BeginPos: types.Point {
					X: 130,
					Y: 30,
				},
				EndPos: types.Point {
					X: 130,
					Y: 80,
				},
			},
			{
				BeginPos: types.Point {
					X: 130,
					Y: 80,
				},
				EndPos: types.Point {
					X: 30,
					Y: 80,
				},
			},
			{
				BeginPos: types.Point {
					X: 30,
					Y: 80,
				},
				EndPos: types.Point {
					X: 30,
					Y: 30,
				},
			},
		},
		Color: sdl.Color { 144, 228, 8 , 0 },
	}

	LINE = types.Line {
		BeginPos: types.Point {
			X: 25,
			Y: 75,
		},
		EndPos: types.Point {
			X: 175,
			//Y: 125,
			Y: 75,
		},
		Invisible: false,
	}
)

func CreateWindows() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer.SetDrawColor(255, 255, 255, 0)
	renderer.Clear()

	shapeRect := utils.AnalyzeShapeInWindow(EXAMPLE_SHAPE, WINDOW)

	utils.DrawByBresenhamAlgorithm([]types.Shape {WINDOW, shapeRect}, 3)

	renderer.Present()

	sdl.Delay(5000)
}