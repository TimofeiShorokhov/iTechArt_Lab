package solution

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

type Circle struct {
	Color
	Center Dot
	Radius float64
	Square float64
	Length float64
}

func (c *Circle) SetColor(col Color) {
	c.Color = col
}

func (c *Circle) FindRadius(p Parallelogramm) float64 {
	c.Radius = math.Sqrt(p.S / math.Pi)
	return c.Radius
}

func (c *Circle) FindLength() float64 {
	c.Length = 2 * (math.Pi) * c.Radius
	return c.Length
}

func (c *Circle) MakeCircle(p Parallelogramm, pole Pole) {
	var color Color

	c.Square = p.S
	c.Center = p.Center
	c.FindRadius(p)
	c.FindLength()
	n := float64(pole.size)
	color = Color((c.Square / n) * 8)
	if color > 8 {
		color = Color(rand.Intn(8-1) + 1)
	}
	c.SetColor(color)
	if len(pole.P) <= c.Center.X || len(pole.P) <= c.Center.Y || c.Center.X < 0 || c.Center.Y < 0 {
		fmt.Println()
		log.Fatal("Index of D.x or D.y is bigger or smaller than size of field, error was taken from MakeParam function, please, change coordinates")
	}
	pole.P[c.Center.X][c.Center.Y] = int(c.Color)
}
