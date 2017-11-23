package main

import (
	"main/lab6WithoutGLUT"
	"main/example/triangle"
	"main/example/cube"
	"main/example/withGLUT"
	"main/lab6"
)

func main() {
	//makeLab6()
	//makeLab6WithoutGlut()
	//makeTriangleExample()
	makeCubeExample()
	//makeGLUTExample()
}

func makeLab6()  {
	lab6.StartLab6()
}

func makeLab6WithoutGlut()  {
	lab6WithoutGLUT.Perform()
}

func makeTriangleExample()  {
	triangle.StartTriangle()
}

func makeCubeExample()  {
	cube.StartCube()
}

func makeGLUTExample() {
	withGLUT.StartGLUT()
}