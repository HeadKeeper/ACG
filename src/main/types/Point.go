package types

import "math"

type Point struct {
	X float64
	Y float64
}

func (p Point) GetLength(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X - p.X, 2) + math.Pow(p2.Y - p.Y, 2))
}