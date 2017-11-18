package withGLUT

import "github.com/vitalibaumtrok/glut/glut-master"

func StartGLUT() {
	glut.Init()
	glut.InitDisplayMode(glut.SINGLE | glut.RGBA)
	glut.InitWindowSize(640, 480)
	glut.CreateWindow("Testing GLUT binding for Go")
	glut.ReshapeFunc(reshape)
	glut.DisplayFunc(display)
	glut.KeyboardFunc(keyboard)
	glut.MainLoop()
}

func reshape(width, height int) {
	println("reshape")
}

func display() {
	println("display")
}

func keyboard(key uint8, x, y int) {
	if key==27 { // escape
		glut.DestroyWindow(glut.GetWindow())
	} else {
		if (glut.GetModifiers() & glut.ACTIVE_CTRL) > 0 {
			println("key pressed: ctrl +", key)
		} else {
			println("key pressed:", key)
		}
	}
}