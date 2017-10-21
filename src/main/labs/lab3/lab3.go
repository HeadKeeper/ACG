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

	SCALE = 3.0
	ROTATE_ANGLE = 2.0
)

func Perform() {
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

	var animatedShape types.Animation
	var shapesToDrawing []types.Shape
	var currentShape types.Shape

	animatedShape.Setup(EXAMPLE_SHAPE)
	EXAMPLE_SHAPE = utils.AnalyzeShapeInWindow(EXAMPLE_SHAPE, WINDOW)
	shapesToDrawing = append(shapesToDrawing, WINDOW, EXAMPLE_SHAPE)
	utils.UpdateShapes(renderer, shapesToDrawing, SCALE)

	mousePressed := false
	prevMouseX := 0
	prevMouseY := 0

	stop := false
	lastTime := uint32(0)

	for !stop {
		currentTime := sdl.GetTicks()

		if animatedShape.Playing {
			if currentTime - lastTime > 50 {
				animatedShape.Play()
				utils.UpdateShapes(renderer, shapesToDrawing, SCALE)
				lastTime = currentTime
			}
		}

		if mousePressed {
			mouseX, mouseY, _ := sdl.GetMouseState()
			dx := mouseX - prevMouseX
			dy := mouseY - prevMouseY

			prevMouseX = mouseX
			prevMouseY = mouseY

			if currentTime - lastTime > 50 {
				currentShape.Rotate(ROTATE_ANGLE)
				lastTime = currentTime
			}

			currentShape.Move(float64(dx), float64(dy))

			utils.UpdateShapes(renderer, shapesToDrawing, SCALE)
		}


		ev := sdl.WaitEvent()
		if ev == nil {
			sdl.Delay(1000)
			continue
		}
		switch ev.(type) {
		case *sdl.QuitEvent:
			stop = true
			break
		case *sdl.KeyDownEvent:
			stop = true
			break
		case *sdl.MouseButtonEvent:
			switch ev.(*sdl.MouseButtonEvent).Button {
			case sdl.BUTTON_LEFT:
				if mousePressed {
					mousePressed = false
				} else {
					if !animatedShape.Playing {
						mouseX, mouseY, _ := sdl.GetMouseState()
						shape := utils.GetTargetShape(shapesToDrawing, float64(mouseX), float64(mouseY))
						mousePressed = true
						prevMouseX = mouseX
						prevMouseY = mouseY
						currentShape = shape
					}
				}
			default:
				continue
			}
		/*case *sdl.MouseMotionEvent:
			mouseX, mouseY, _ := sdl.GetMouseState()
			prevMouseX = mouseX
			prevMouseY = mouseY*/
		default:
			continue
		}
	}
}