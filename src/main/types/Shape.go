package types

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Shape struct {
	Lines []Line
	Color sdl.Color
	ZIndex int
}

func (thisShape Shape) GetCenter() Point {
	var pivot Point

	var allX float64
	var allY float64
	for _, line := range thisShape.Lines {
		 allX += line.BeginPos.X
		 allY += line.BeginPos.Y
	}

	pivot = Point {
		X: allX / float64(len(thisShape.Lines)),
		Y: allY / float64(len(thisShape.Lines)),
	}

	return pivot
}

func (thisShape Shape) Rotate(angle float64) {
	angle *= math.Pi / 180
	pivot := thisShape.GetCenter()

	for index := range thisShape.Lines {
		thisShape.Lines[index] = Line {
			BeginPos: transformPoint(thisShape.Lines[index].BeginPos, pivot, angle),
			EndPos: transformPoint(thisShape.Lines[index].EndPos, pivot, angle),
		}
	}

}

func transformPoint(point Point, pivot Point, angle float64) Point {
	x := point.X - pivot.X
	y := point.Y - pivot.Y

	x = (x * math.Cos(angle) - y * math.Sin(angle)) + pivot.X
	y = (x * math.Sin(angle) + y * math.Cos(angle)) + pivot.Y

	point = Point {
		X: x,
		Y: y,
	}

	return point
}

func (thisShape Shape) Move(xTransition float64, yTransition float64) {
	for index := range thisShape.Lines {
		thisShape.Lines[index] = Line {
			BeginPos: movePoint(thisShape.Lines[index].BeginPos, xTransition, yTransition),
			EndPos: movePoint(thisShape.Lines[index].EndPos, xTransition, yTransition),
		}
	}
}

func movePoint(point Point, xTransition float64, yTransition float64) Point {
	x := point.X + xTransition
	y := point.Y + yTransition

	point = Point {
		X: x,
		Y: y,
	}

	return point
}