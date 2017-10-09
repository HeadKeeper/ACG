package types

import (
	"fmt"
	"math"
)

type LineVector struct {
	X float64
	Y float64
}

func NewLineVector(x, y float64) LineVector {
	return LineVector{x, y}
}

func FromScalar(v float64) LineVector {
	return LineVector{v, v}
}

func Zero() LineVector {
	return LineVector{0, 0}
}

func Unit() LineVector {
	return LineVector{1, 1}
}

func (v LineVector) Copy() LineVector {
	return LineVector{v.X, v.Y}
}

func (v LineVector) Add(v2 LineVector) LineVector {
	return LineVector{v.X + v2.X, v.Y + v2.Y}
}

func (v LineVector) Subtract(v2 LineVector) LineVector {
	return LineVector{v.X - v2.X, v.Y - v2.Y}
}

func (v LineVector) Multiply(v2 LineVector) LineVector {
	return LineVector{v.X * v2.X, v.Y * v2.Y}
}

func (v LineVector) Divide(v2 LineVector) LineVector {
	return LineVector{v.X / v2.X, v.Y / v2.Y}
}

func (v LineVector) MultiplyScalar(s float64) LineVector {
	return LineVector{v.X * s, v.Y * s}
}

func (v LineVector) MultiplyScalarOnVector(v2 LineVector) float64 {
	return v.GetLength() * v2.GetLength() * v.GetCosAngleBetweenVector(v2)
}

func (v LineVector) DivideScalar(s float64) LineVector {
	return LineVector{v.X / s, v.Y / s}
}

func (v LineVector) GetLength() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v LineVector) GetCosAngleBetweenVector(v2 LineVector) float64 {
	return (v.X * v2.X + v.Y * v2.Y) / (v.GetLength() * v2.GetLength())
}

func (v LineVector) String() string {
	return fmt.Sprintf("%v:%v", v.X, v.Y)
}
