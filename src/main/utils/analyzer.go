package utils

import (
	"main/types"
	"fmt"
)

// Cyrus & Beck algorithm
func GetInsideLinePart(line types.Line, shape types.Shape) []types.Line {
	//tIn := float64(0)
	//tOut := float64(1)

	vectorBegin, vectorEnd := CreateVectorsForLine(line)
	vectorD := vectorEnd.Subtract(vectorBegin)

	//var vectorsP []types.LineVector
	var vectorsN []types.LineVector = GetShapeSidesVectors(shape)

	for sideNumber, edge := range shape.Lines {
		valueP := vectorsN[sideNumber].MultiplyScalarOnVector(vectorD)
		vectorF, _ := CreateVectorsForLine(edge)
		t := vectorsN[sideNumber].MultiplyScalarOnVector(vectorBegin.Subtract(vectorF)) /
			(vectorsN[sideNumber].MultiplyScalarOnVector(vectorEnd.Subtract(vectorBegin))) * -1
		fmt.Println("P : ", valueP)
		fmt.Println("t: ", t)
		//fun := AnalyzePValue(valueP)

	}

	return nil
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
