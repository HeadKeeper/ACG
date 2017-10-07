package utils

import (
	"main/types"
	"github.com/veandco/go-sdl2/sdl"
	"math"
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

func drawLineByBresenhamAlgorithm(renderer *sdl.Renderer, x1 int, y1 int, x2 int, y2 int, color sdl.Color) {
	renderer.SetDrawColor(color.R, color.G, color.B, color.A)

	dx := int(math.Abs(float64(x2 - x1)))
	dy := int(math.Abs(float64(y2 - y1)))

	var sx int
	if x2 >= x1 {
		sx = 1
	} else {
		sx = -1
	}

	var sy int
	if y2 >= y1 {
		sy = 1
	} else {
		sy = -1
	}

	if dy <= dx {
		d := (dy << 1) - dx
		d1 := dy << 1
		d2 := (dy - dx) << 1

		renderer.DrawPoint(x1, y1)

		x := x1 + sx
		y := y1
		for i := 1; i <= dx; i ++ {
			if d > 0 {
				d += d2
				y += sy
			} else {
				d += d1
			}

			renderer.DrawPoint(x, y)

			x += sx
		}
	} else {
		d := (dx << 1) - dy
		d1 := dx << 1
		d2 := (dx - dy) << 1

		renderer.DrawPoint(x1, y1)

		x := x1
		y := y1 + sy
		for i := 1; i <= dy; i ++ {
			if d > 0 {
				d += d2
				x += sx
			} else {
				d += d1
			}

			renderer.DrawPoint(x, y)

			y += sy
		}
	}
}

func DrawByBresenhamAlgorithm(shapes []types.Shape, scale float64) {
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
			drawLineByBresenhamAlgorithm(
				renderer,
				int(scale * side.BeginPos.X),
				int(scale * side.BeginPos.Y),
				int(scale * side.EndPos.X),
				int(scale * side.EndPos.Y),
				shape.Color,
			)
		}
	}

	renderer.Present()

	sdl.Delay(5000)
}