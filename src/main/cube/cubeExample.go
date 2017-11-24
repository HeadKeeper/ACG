package cube

import (
	_ "image/png"
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"fmt"
)

const windowWidth = 800
const windowHeight = 600

var (
	// Camera
	cameraPos = mgl32.Vec3{0.0, 0.0, 8.0}
	cameraFront = mgl32.Vec3{0.0, 0.0, -1.0}
	cameraUp = mgl32.Vec3{0.0, 1.0, 1.0}

	// Keys control
	keys [1024]bool
	firstMouse = true
	mousePressed = false
	yaw = float32(-90.0)
	pitch = float32(0.0)
	lastX = float64(windowWidth / 2.0)
	lastY = float64(windowHeight / 2.0)

	// Light
	lightPos = mgl32.Vec3{1.2, 1.0, 2.0}
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func StartCube() {
	window := createWindow()

	program, model, modelUniform, camera, cameraUniform, projection, _ := configureWindow()
	lightProgram := configureLight()

	// Set the container's VAO (and containerVBO)
	var containerVBO, containerVAO uint32
	gl.GenVertexArrays(1, &containerVAO)
	gl.GenBuffers(1, &containerVBO)

	gl.BindBuffer(gl.ARRAY_BUFFER, containerVBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)

	gl.BindVertexArray(containerVAO)

	// Position attribute
	vertexAttribute := uint32(0)/*uint32(gl.GetAttribLocation(program, gl.Str("vertexIn\x00")))*/
	gl.VertexAttribPointer(vertexAttribute, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(vertexAttribute)
	gl.BindVertexArray(vertexAttribute)

	/*
	// Load the texture
	texture := newTexture("square.png")
	textureCoordinateAttribute := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(textureCoordinateAttribute)
	gl.VertexAttribPointer(textureCoordinateAttribute, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
	*/

	// Set the light's VAO (VBO stays the same. After all, the vertices are the same for the light object (also a 3D cube)
	var lightVAO uint32
	gl.GenVertexArrays(1, &lightVAO)
	gl.BindVertexArray(lightVAO)
	// We only need to bind to the VBO (to link it with gl.VertexAttribPointer), no need to fill it; the VBO's data already contains all we need.
	gl.BindBuffer(gl.ARRAY_BUFFER, containerVBO)
	// Set the vertex attributes (only position data for the lamp))
	vertexLightAttribute := vertexAttribute/*uint32(gl.GetAttribLocation(program, gl.Str("vertexIn\x00")))*/
	gl.VertexAttribPointer(vertexLightAttribute, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(vertexLightAttribute)
	gl.BindVertexArray(vertexLightAttribute)

	angle := 0.0
	previousFrame := glfw.GetTime()

	for !window.ShouldClose() {

		// Calculate delta time of current frame
		currentFrame := glfw.GetTime()
		elapsed := currentFrame - previousFrame
		previousFrame = currentFrame

		// Check if any events have been activiated (key pressed, mouse moved etc.) and call corresponding response functions
		glfw.PollEvents()
		doMovement()

		// Clear data for new frame
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Use corresponding shader when setting uniforms/drawing objects
		gl.UseProgram(program)
		objectColorLoc := gl.GetUniformLocation(program, gl.Str("objectColor\x00"))
		lightColorLoc  := gl.GetUniformLocation(program, gl.Str("lightColor\x00"))
		gl.Uniform3f(objectColorLoc, 1.0, 0.5, 0.31)
		gl.Uniform3f(lightColorLoc, 1.0, 0.5, 1.0)

		// Change camera
		camera = mgl32.LookAtV(cameraPos, cameraPos.Add(cameraFront), cameraUp)
		gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

		// Change model
		angle += elapsed
		model = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0.1, 1, 0})
		gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

		/*
		// Make texture active
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, texture)
		*/

		// Draw the light object (using light's vertex attributes)
		gl.BindVertexArray(containerVAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 36)
		gl.BindVertexArray(0)

		// Also draw the lamp object, again binding the appropriate shader
		gl.UseProgram(lightProgram)
		modelLightUniform := gl.GetUniformLocation(lightProgram, gl.Str("model\x00"))
		viewLightUniform := gl.GetUniformLocation(lightProgram, gl.Str("view\x00"))
		projectionLightUniform := gl.GetUniformLocation(lightProgram, gl.Str("projection\x00"))
		// Set matrices
		gl.UniformMatrix4fv(viewLightUniform, 1, false, &camera[0])
		gl.UniformMatrix4fv(projectionLightUniform, 1, false, &projection[0])
		var modelLight mgl32.Mat4
		modelLight = modelLight.Mul4(mgl32.Translate3D(lightPos.X(), lightPos.Y(), lightPos.Z()))
		modelLight = modelLight.Mul4(mgl32.Scale3D(0.2, 0.2, 0.2)) // Make it a smaller cube
		gl.UniformMatrix4fv(modelLightUniform, 1, false, &model[0])

		// Draw the light object (using light's vertex attributes)
		gl.BindVertexArray(lightVAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 36)
		gl.BindVertexArray(0)

		// Maintenance
		window.SwapBuffers()
	}

}


func createWindow() *glfw.Window {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetKeyCallback(onKey)
	window.SetCursorPosCallback(onMouse)
	window.SetMouseButtonCallback(onMouseButton)

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	return window
}

func configureWindow() (uint32, mgl32.Mat4, int32, mgl32.Mat4, int32, mgl32.Mat4, int32) {
	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	program := NewShaderProgram("object.vertex.shader", "object.fragment.shader")
	gl.UseProgram(program)

	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(windowWidth)/windowHeight, 0.1, 100.0)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	cameraUniform := gl.GetUniformLocation(program, gl.Str("view\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	/*
	textureUniform := gl.GetUniformLocation(program, gl.Str("textureUni\x00"))
	gl.Uniform1i(textureUniform, 0)
	*/
	//gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))

	return program, model, modelUniform, camera, cameraUniform, projection, projectionUniform
}

func configureLight() (uint32) {
	program := NewShaderProgram("lamp.vertex.shader", "lamp.fragment.shader")

	gl.UseProgram(program)

	objectColor := mgl32.Vec3{1, 0.5, 0.31}
	objectColorUniform := gl.GetUniformLocation(program, gl.Str("objectColor\x00"))
	gl.Uniform3fv(objectColorUniform, 1, &objectColor[0])

	lightColor := mgl32.Vec3{1, 1, 1}
	lightColorUniform := gl.GetUniformLocation(program, gl.Str("lightColor\x00"))
	gl.Uniform3fv(lightColorUniform, 1, &lightColor[0])

	return program
}