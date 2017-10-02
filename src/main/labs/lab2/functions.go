package lab2

import (
	"github.com/veandco/go-sdl2/sdl"
	"main/utils"
	"main/types"
)

func CreateInitialShape(sides []types.Side) types.Shape {
	return CreateShape(
		sides,
		utils.CreateRandomColor(),
	)
}

func CreateShape(sides []types.Side, color sdl.Color) types.Shape {
	return types.Shape{
		Sides: sides,
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

func createNewSides(nu float64, shape types.Shape) []types.Side {
	var newSides []types.Side
	for index := 0; index < len(shape.Sides); index++ {
		newSideStartPos := getPointFromSide(shape.Sides[index], nu)
		newSideEndPos := getPointFromSide(getNextSide(index, shape.Sides), nu)
		newSides = append(newSides, types.Side{
			BeginPos: newSideStartPos,
			EndPos: newSideEndPos,
		})
	}
	return newSides
}

func getPointFromSide(side types.Side, nu float64) types.Point {
	pointX := side.BeginPos.X + nu * (side.EndPos.X - side.BeginPos.X)
	if pointX < 0 {
		pointX = pointX * -1
	}
	pointY := side.BeginPos.Y + nu * (side.EndPos.Y - side.BeginPos.Y)
	if pointY < 0 {
		pointY = pointY * -1
	}
	return types.Point{
		X: pointX,
		Y: pointY,
	}
}

func getNextSideNumber(currentSideNumber int, amountOfSides int) int {
	var newSideNumber int
	if currentSideNumber >= amountOfSides - 1 {
		newSideNumber = 0
	} else {
		newSideNumber = currentSideNumber + 1
	}

	return newSideNumber
}

func getNextSide(currentSideNumber int, sides []types.Side) types.Side {
	return sides[getNextSideNumber(currentSideNumber, len(sides))]
}