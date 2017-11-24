package cube

import (
	"strings"
	"github.com/go-gl/gl/v3.3-core/gl"
	"os"
	"log"
	"bufio"
	"path/filepath"
)

func NewShaderProgram(vertexShaderPath, fragmentShaderPath string) uint32 {
	vertexShader, fragmentShader, err := loadShaders(vertexShaderPath, fragmentShaderPath)
	if err != nil {
		panic(err)
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)

	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		loggedData := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(loggedData))

		panic("failed to link program: " + loggedData)
	}

	gl.DetachShader(program, vertexShader)
	gl.DetachShader(program, fragmentShader)

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program
}

func loadShaders(vertexFilePath, fragmentFilePath string) (uint32, uint32, error) {

	// Compile vertex shader
	vertexShaderID := compileShader(readShaderCode(vertexFilePath), gl.VERTEX_SHADER)

	// Compile fragment shader
	fragmentShaderID := compileShader(readShaderCode(fragmentFilePath), gl.FRAGMENT_SHADER)

	return vertexShaderID, fragmentShaderID, nil
}

// Compile shader. Source is null terminated c string. shader type is self
// explanatory
func compileShader(source string, shaderType uint32) uint32 {

	// Create new shader
	shader := gl.CreateShader(shaderType)
	// Convert shader string to null terminated c string
	shaderCode, free := gl.Strs(source)
	defer free()
	gl.ShaderSource(shader, 1, shaderCode, nil)

	// Compile shader
	gl.CompileShader(shader)

	// Check shader status
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		loggedData := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(loggedData))

		panic("failed to compile \n" + source + "\n:" + loggedData)
	}
	return shader
}

// Read shader code from file
func readShaderCode(fileName string) string {
	code := ""
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(dir + "/src/main/cube/shaders/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		code += "\n" + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	code += "\x00"
	return code
}
