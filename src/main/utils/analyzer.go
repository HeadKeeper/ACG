package utils

import (
	"main/types"
)

// Cyrus & Beck algorithm
func AnalyzeLineInWindow(line types.Line, window types.Shape) []types.Line {
	beginT := 0.0
	endT := 1.0
	var lines []types.Line

	vectorBegin, vectorEnd := CreateVectorsForLine(line)
	vectorD := vectorEnd.Subtract(vectorBegin)

	var vectorsN []types.LineVector = GetShapeSidesVectors(window)

	for sideNumber, edge := range window.Lines {
		valueP := vectorsN[sideNumber].MultiplyScalarOnVector(vectorD)
		vectorF, _ := CreateVectorsForLine(edge)
		vectorW := vectorBegin.Subtract(vectorF)
		valueQ := vectorsN[sideNumber].MultiplyScalarOnVector(vectorW)

   		if valueP == 0 {
			if valueQ < 0 {
			//	Outside window
				line.Invisible = true
				lines = append(lines, line)
				return lines
			}
		}

		t := -1 * valueQ / valueP
		if t > 1 || t < 0 {
			//	T is outside window
			continue
		}

		if valueP < 0 && beginT <= t && endT > t {
			endT = t
		}

		if valueP > 0 && endT >= t && beginT < t {
			beginT = t
		}

	}

	lineParameter := GetLineParameterVersion(line)

	if beginT != 0 {
		lines = append(lines, types.Line{
			BeginPos:  lineParameter(0),
			EndPos:    lineParameter(beginT),
			Invisible: true,
		})
	}
	lines = append(lines, types.Line{
		BeginPos:  lineParameter(beginT),
		EndPos:    lineParameter(endT),
	})
	if endT != 1 {
		lines = append(lines, types.Line{
			BeginPos:  lineParameter(endT),
			EndPos:    lineParameter(1),
			Invisible: true,
		})
	}

	return lines
}

func AnalyzeShapeInWindow(shape types.Shape, window types.Shape) types.Shape {
	var newShape types.Shape
	newShape.Color = shape.Color

	for lineIndex := range shape.Lines {
		newShape.Lines = append(newShape.Lines, AnalyzeLineInWindow(shape.Lines[lineIndex], window)...)
	}
	return newShape
}

func GetShapeSidesVectors(shape types.Shape) []types.LineVector {
	var vectors []types.LineVector
	for _, line := range shape.Lines {
		vectors = append(vectors, CreateNVector(line))
	}

	return vectors
}

func CreateNVector(line types.Line) types.LineVector {
	return types.NewLineVector(line.EndPos.X - line.BeginPos.X, line.EndPos.Y - line.BeginPos.Y)
}

func CreateVectorsForLine(line types.Line) (types.LineVector, types.LineVector) {
	v0 := types.NewLineVector(line.BeginPos.X, line.BeginPos.Y)
	v1 := types.NewLineVector(line.EndPos.X, line.EndPos.Y)

	return v0, v1
}

func GetNextSideNumber(currentSideNumber int, amountOfSides int) int {
	var newSideNumber int
	if currentSideNumber >= amountOfSides - 1 {
		newSideNumber = 0
	} else {
		newSideNumber = currentSideNumber + 1
	}

	return newSideNumber
}

func GetPrevSideNumber(currentSideNumber int, amountOfSides int) int {
	var newSideNumber int
	if currentSideNumber >= amountOfSides - 1 {
		newSideNumber = 0
	} else {
		newSideNumber = currentSideNumber - 1
	}

	return newSideNumber
}

func GetNextSide(currentSideNumber int, sides []types.Line) types.Line {
	return sides[GetNextSideNumber(currentSideNumber, len(sides))]
}

func GetPrevSide(currentSideNumber int, sides []types.Line) types.Line {
	return sides[GetPrevSideNumber(currentSideNumber, len(sides))]
}

func GetLineParameterVersion(line types.Line) func(t float64) types.Point {
	return func(t float64) types.Point {
		return types.Point{
			Y: line.BeginPos.Y + (line.EndPos.Y - line.BeginPos.Y) * t,
			X: line.BeginPos.X + (line.EndPos.X - line.BeginPos.X) * t,
		}
	}
}