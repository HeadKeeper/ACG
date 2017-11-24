package cube

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func doMovement() {
	var cameraSpeed float32 = 0.1
	if keys[glfw.KeyUp] {
		cameraPos = cameraPos.Add(cameraFront.Mul(cameraSpeed))
	}
	if keys[glfw.KeyDown] {
		cameraPos = cameraPos.Sub(cameraFront.Mul(cameraSpeed))
	}
	if keys[glfw.KeyLeft] {
		cameraPos = cameraPos.Sub(cameraFront.Cross(cameraUp).Normalize().Mul(cameraSpeed))
	}
	if keys[glfw.KeyRight] {
		cameraPos = cameraPos.Add(cameraFront.Cross(cameraUp).Normalize().Mul(cameraSpeed))
	}
}

func onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
	if key >= 0 && key < 1024 {
		if action == glfw.Press {
			keys[key] = true
		} else if action == glfw.Release {
			keys[key] = false
		}
	}
}

func onMouse(w *glfw.Window, xPos float64, yPos float64)  {
	if mousePressed {
		if firstMouse {
			lastX = xPos
			lastY = yPos
			firstMouse = false
		}

		xOffset := float32(xPos - lastX)
		yOffset := float32(lastY - yPos) // Reversed since y-coordinates go from bottom to left
		lastX = xPos
		lastY = yPos

		sensitivity := float32(0.03) // Change this value to your liking
		xOffset *= sensitivity
		yOffset *= sensitivity

		yaw += xOffset
		pitch += yOffset

		// Make sure that when pitch is out of bounds, screen doesn't get flipped
		if pitch > 89.0 {
			pitch = 89.0
		}
		if pitch < -89.0 {
			pitch = -89.0
		}

		front := mgl32.Vec3{
			float32(math.Cos(float64(mgl32.DegToRad(yaw)) * math.Cos(float64(mgl32.DegToRad(pitch))))),
			float32(math.Sin(float64(mgl32.DegToRad(pitch)))),
			float32(math.Sin(float64(mgl32.DegToRad(yaw)) * math.Cos(float64(mgl32.DegToRad(pitch))))),
		}
		cameraFront = front.Normalize()
	}
}

func onMouseButton(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey)  {
	if button == glfw.MouseButtonLeft && action == glfw.Press {
		mousePressed = true
	} else if button == glfw.MouseButtonLeft && action == glfw.Release {
		mousePressed = false
		firstMouse = true
	}
}