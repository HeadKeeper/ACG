package utils

import (
	"image/color"
	"golang.org/x/image/colornames"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

func CreateRandomColor() sdl.Color {
	var colors []color.RGBA
	for _, value := range colornames.Map {
		colors = append(colors, value)
	}
	rgba := colors[random(0, 146)]
	return sdl.Color{
		R: rgba.R,
		G: rgba.G,
		B: rgba.B,
		A: rgba.A,
	}
}

func random(min, max int) int {
	rand.Seed(rand.Int63())
	return rand.Intn(max - min) + min
}
