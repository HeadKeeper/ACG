package lab3

import (
	"github.com/veandco/go-sdl2/sdl"
	"main/types"
	"main/utils"
)

var (
	RECT = []types.Line {
		{
			BeginPos: types.Point{
				X: 50,
				Y: 50,
			},
			EndPos: types.Point{
				X: 150,
				Y: 50,
			},
		},
		{
			BeginPos: types.Point{
				X: 150,
				Y: 50,
			},
			EndPos: types.Point{
				X: 150,
				Y: 100,
			},
		},
		{
			BeginPos: types.Point{
				X: 150,
				Y: 100,
			},
			EndPos: types.Point{
				X: 50,
				Y: 100,
			},
		},
		{
			BeginPos: types.Point{
				X: 50,
				Y: 100,
			},
			EndPos: types.Point{
				X: 50,
				Y: 50,
			},
		},
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

	shape := types.Shape {
		Lines: RECT,
		Color: sdl.Color { 228, 148, 8 , 0 },
	}

	utils.GetInsideLinePart(
		LINE,
		shape,
	)

	utils.Draw([]types.Shape{ shape, {[]types.Line{LINE}, sdl.Color { 148, 8, 228 , 0 }}}, 2)

	renderer.Present()

	sdl.Delay(5000)
}