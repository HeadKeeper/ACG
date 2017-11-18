package triangle


import (
	"runtime"
	"log"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v3.3-core/gl"
	"fmt"
	"strings"
)

const (
	width  = 500
	height = 500

	openGLVersionMajor = 3
	openGLVersionMinor = 3
	vertexShaderSource = `
    #version 330
    layout (location = 0) in vec3 vp;
    void main() {
        gl_Position = vec4(vp, 1.0);
    }` + "\x00"

	fragmentShaderSource = `
    #version 330
    out vec4 frag_colour;
    void main() {
        frag_colour = vec4(1, 0.2, 0.1, 0);
    }` + "\x00"
)

var (
	triangleVertices = []float32{
		 0.5,  0.5, 0, // top
		 0.5, -0.5, 0, // right
		-0.5, -0.5, 0, // left
		-0.5,  0.5, 0, // bottom
	}

	triangles = []int32{
		0, 1, 2, // first
		1, 2, 3, // second
	}
)

func StartTriangle() {
	runtime.LockOSThread()

	window := initGLFW()
	defer glfw.Terminate()
	program := initOpenGL()

	VBO := makeVBO(triangleVertices)
	//EBO := makeEBO(triangles)

	VAO := makeVAOFromVBO(VBO)

	for !window.ShouldClose() {
		draw(VAO, window, program)
	}
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER) // Returning new vector which contains 4 params
	if err != nil {
		panic(err)
	}
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER) // Calculating colors of pixels
	if err != nil {
		panic(err)
	}

	program := gl.CreateProgram() // Creating program for shader
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program) // Make this program active for drawing
	return program
}

func initGLFW() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, openGLVersionMajor)
	glfw.WindowHint(glfw.ContextVersionMinor, openGLVersionMinor)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Conway's Game of Life", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}


func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	cSources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, cSources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logging := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logging))

		return 0, fmt.Errorf("failed to compile %v: %v", source, logging)
	}

	return shader, nil
}

func draw(vao uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangleVertices) / 3))
	gl.BindVertexArray(0)

	glfw.PollEvents()
	window.SwapBuffers()
}

func drawElements(vao uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawElements(gl.TRIANGLES, 6, uint32(len(triangleVertices) / 3), nil)
	gl.BindVertexArray(0)

	glfw.PollEvents()
	window.SwapBuffers()
}

func makeVBO(points []float32) uint32 {
	var vbo uint32 // Vertex buffer object
	gl.GenBuffers(1, &vbo) // Create OpenGL ID for vertex buffer object
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo) // Bind vertex buffer object to ARRAY_BUFFER
	gl.BufferData(gl.ARRAY_BUFFER, len(points) / 3, gl.Ptr(points), gl.STATIC_DRAW) // Copy points data into ARRAY_BUFFER

	/*
	* gl.STATIC_DRAW - for static data
	* gl.DYNAMIC_DRAW - for dynamic data
	* gl.STREAM_DRAW - for objects which will update on every rendering
	*/

	return vbo
}

func makeVAOFromVBO(vbo uint32) uint32 {
	var vao uint32 // Vertex array object
	gl.GenVertexArrays(1, &vao) // Create OpenGL ID for vertex array object
	gl.BindVertexArray(vao) // Bind vertex array object to OpenGL instance
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil) // Set pointers on vertex data
	// First arg - location from vertex shader
	// Second arg - size of arg in vertex shader
	// Third arg - data type
	// Fourth arg - normalization (true -> all data will be placed between 0 and 1 (or -1 for sign values)
	// Fifth arg - distance between data sets (0 -> auto, but only for packed-data)
	// Sixth arg - pointer to array begin (nil -> array beginning from 0)

	return vao
}

func makeVAOFromEBO(vbo, ebo uint32) uint32 {
	var vao uint32 // Vertex array object
	gl.GenVertexArrays(1, &vao) // Create OpenGL ID for vertex array object
	gl.BindVertexArray(vao) // Bind vertex array object to OpenGL instance
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil) // Set pointers on vertex data

	return vao
}

func makeEBO(indexes []int32) uint32 {
	var ebo uint32                                                                        // Element buffer object
	gl.GenBuffers(1, &ebo)                                                             // Create OpenGL ID for element buffer object
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)                                           // Bind vertex buffer object to ARRAY_BUFFER
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indexes), gl.Ptr(indexes), gl.STATIC_DRAW) // Copy indexes data into ARRAY_BUFFER

	return ebo
}