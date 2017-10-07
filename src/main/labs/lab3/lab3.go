package lab3

import "github.com/veandco/go-sdl2/sdl"

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

	subWindow, subRenderer, err := sdl.CreateWindowAndRenderer(400, 300, sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}
	defer subWindow.Destroy()
	subRenderer.SetDrawColor(255, 0, 255, 0)
	subRenderer.Clear()

	renderer.Present()
	subRenderer.Present()

	sdl.Delay(5000)
}