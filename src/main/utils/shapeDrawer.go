package utils

import (
	"main/types"
	"github.com/veandco/go-sdl2/sdl"
)

func Draw(shapes []types.Shape, scale float64) {
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

	for _, shape := range shapes {
		renderer.SetDrawColor(shape.Color.R, shape.Color.G, shape.Color.B, shape.Color.A)
		for _, side := range shape.Sides {
			renderer.DrawLine(
				int(scale * side.BeginPos.X),
				int(scale * side.BeginPos.Y),
				int(scale * side.EndPos.X),
				int(scale * side.EndPos.Y),
			)
		}
	}

	renderer.Present()

	sdl.Delay(5000)
}