package lab2

import (
	"main/types"
	"main/utils"
)

const (
	AMOUNT_OF_SHAPES = 20
	SCALE float64 = 5
	NU float64 = 0.1
)

var (
	INITIAL_SHAPE__RECT = []types.Side {
		{
			BeginPos: types.Point{
				X: 50,
				Y: 50,
			},
			EndPos: types.Point{
				X: 150,
				Y: 50,
			},
		},
		{
			BeginPos: types.Point{
				X: 150,
				Y: 50,
			},
			EndPos: types.Point{
				X: 150,
				Y: 100,
			},
		},
		{
			BeginPos: types.Point{
				X: 150,
				Y: 100,
			},
			EndPos: types.Point{
				X: 50,
				Y: 100,
			},
		},
		{
			BeginPos: types.Point{
				X: 50,
				Y: 100,
			},
			EndPos: types.Point{
				X: 50,
				Y: 50,
			},
		},
	}

	INITIAL_SHAPE__TRIANGLE = []types.Side {
		{
			BeginPos: types.Point{
				X: 50,
				Y: 50,
			},
			EndPos: types.Point{
				X: 75,
				Y: 100,
			},
		},
		{
			BeginPos: types.Point{
				X: 75,
				Y: 100,
			},
			EndPos: types.Point{
				X: 100,
				Y: 50,
			},
		},
		{
			BeginPos: types.Point{
				X: 100,
				Y: 50,
			},
			EndPos: types.Point{
				X: 50,
				Y: 50,
			},
		},
	}

	INITIAL_SHAPE__FIVEANGLE = []types.Side {
		{
			BeginPos: types.Point{
				X: 25,
				Y: 60,
			},
			EndPos: types.Point{
				X: 75,
				Y: 80,
			},
		},
		{
			BeginPos: types.Point{
				X: 75,
				Y: 80,
			},
			EndPos: types.Point{
				X: 125,
				Y: 60,
			},
		},
		{
			BeginPos: types.Point{
				X: 125,
				Y: 60,
			},
			EndPos: types.Point{
				X: 100,
				Y: 40,
			},
		},
		{
			BeginPos: types.Point{
				X: 100,
				Y: 40,
			},
			EndPos: types.Point{
				X: 50,
				Y: 40,
			},
		},
		{
			BeginPos: types.Point{
				X: 50,
				Y: 40,
			},
			EndPos: types.Point{
				X: 25,
				Y: 60,
			},
		},
	}

	INITIAL_SHAPE__SIXANGLE = []types.Side {
		{
			BeginPos: types.Point{
				X: 25,
				Y: 60,
			},
			EndPos: types.Point{
				X: 50,
				Y: 80,
			},
		},
		{
			BeginPos: types.Point{
				X: 50,
				Y: 80,
			},
			EndPos: types.Point{
				X: 75,
				Y: 80,
			},
		},
		{
			BeginPos: types.Point{
				X: 75,
				Y: 80,
			},
			EndPos: types.Point{
				X: 100,
				Y: 60,
			},
		},
		{
			BeginPos: types.Point{
				X: 100,
				Y: 60,
			},
			EndPos: types.Point{
				X: 75,
				Y: 40,
			},
		},
		{
			BeginPos: types.Point{
				X: 75,
				Y: 40,
			},
			EndPos: types.Point{
				X: 50,
				Y: 40,
			},
		},
		{
			BeginPos: types.Point{
				X: 50,
				Y: 40,
			},
			EndPos: types.Point{
				X: 25,
				Y: 60,
			},
		},
	}
)

func Perform() {
	PerformWithScale(SCALE)
}

func PerformWithScale(scale float64) {
	var shapes []types.Shape
	shapes = append(shapes, CreateInitialShape(INITIAL_SHAPE__RECT))
	for index := 1; index <= AMOUNT_OF_SHAPES; index++ {
		shapes = append(shapes, CreateNewShape(shapes[index - 1]))
	}

	utils.Draw(shapes, scale)
}