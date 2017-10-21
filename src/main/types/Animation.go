package types

import "main/values"

type Animation struct {
	DX         float64
	DY         float64
	DAngle     float64
	StepNumber int
	Playing    bool
	Shape      Shape
}

func (thisAnimation Animation) isPlaying() bool {
	return thisAnimation.Playing
}

func (thisAnimation Animation) Setup(shape Shape) {
	pivot := shape.GetCenter()
	thisAnimation.DX = (pivot.X - values.X_DISTANCE) / ( values.NUMBER_STEPS * 2)
	thisAnimation.DY = (pivot.Y - values.Y_DISTANCE) / ( values.NUMBER_STEPS * 2)
	thisAnimation.Shape = shape
	thisAnimation.StepNumber = 0
	thisAnimation.DAngle = 90 / float64(thisAnimation.StepNumber)
}

func (thisAnimation Animation) Play() {
	thisAnimation.Playing = true

	thisAnimation.Shape.Move(thisAnimation.DX, thisAnimation.DY)
	thisAnimation.Shape.Rotate(thisAnimation.DAngle)

	thisAnimation.StepNumber += 1
	if thisAnimation.StepNumber >= values.NUMBER_STEPS {
		thisAnimation.Stop()
	}
}

func (thisAnimation Animation) Stop() {
	thisAnimation.Playing = false
}
