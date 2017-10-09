package lab2

import (
	"github.com/veandco/go-sdl2/sdl"
	"main/utils"
	"main/types"
)

func CreateInitialShape(lines []types.Line) types.Shape {
	return CreateShape(
		lines,
		utils.CreateRandomColor(),
	)
}

func CreateShape(lines []types.Line, color sdl.Color) types.Shape {
	return types.Shape{
		Lines: lines,
		Color: color,
	}
}

func CreateNewShape(shape types.Shape) types.Shape {
	return CreateNewShapeWithNu(shape, NU)
}

func CreateNewShapeWithNuAndSameColor(shape types.Shape, nu float64) types.Shape {
	newShape := CreateNewShapeWithNu(shape, nu)
	newShape.Color = shape.Color
	return newShape
}

func CreateNewShapeWithNu(shape types.Shape, nu float64) types.Shape {
	return CreateShape(
		createNewSides(nu, shape),
		utils.CreateRandomColor(),
	)
}

func createNewSides(nu float64, shape types.Shape) []types.Line {
	var newSides []types.Line
	for index := 0; index < len(shape.Lines); index++ {
		newSideStartPos := getPointFromSide(shape.Lines[index], nu)
		newSideEndPos := getPointFromSide(utils.GetNextSide(index, shape.Lines), nu)
		newSides = append(newSides, types.Line{
			BeginPos: newSideStartPos,
			EndPos: newSideEndPos,
		})
	}
	return newSides
}

func getPointFromSide(line types.Line, nu float64) types.Point {
	pointX := line.BeginPos.X + nu * (line.EndPos.X - line.BeginPos.X)
	if pointX < 0 {
		pointX = pointX * -1
	}
	pointY := line.BeginPos.Y + nu * (line.EndPos.Y - line.BeginPos.Y)
	if pointY < 0 {
		pointY = pointY * -1
	}
	return types.Point{
		X: pointX,
		Y: pointY,
	}
}

