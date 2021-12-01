package solution

import (
	"fmt"
	"log"
	"math"
)

type Parallelogramm struct {
	A, B, C, D, Center Dot
	S                  float64
}

func (p *Parallelogramm) CreateFourthDot(a, b, c Dot, pole Pole) *Parallelogramm {

	var d Dot
	var color Color

	p.A = a
	p.B = b
	p.C = c

	centerX := (a.X + c.X) / 2
	centerY := (a.Y + c.Y) / 2

	if (2*centerX-b.X) >= 0 && (2*centerY-b.Y) >= 0 && (2*centerX-b.X) < pole.size && (2*centerY-b.Y) < pole.size {
		d.X = 2*centerX - b.X
		d.Y = 2*centerY - b.Y
		p.D = d
	} else {
		centerX = (b.X + c.X) / 2
		centerY = (b.Y + c.Y) / 2
		if (2*centerX-a.X) >= 0 && (2*centerY-a.Y) >= 0 && (2*centerX-a.X) < pole.size && (2*centerY-a.Y) < pole.size {
			d.X = 2*centerX - a.X
			d.Y = 2*centerY - a.Y
			p.D = d
		} else {
			centerX = (a.X + b.X) / 2
			centerY = (a.Y + b.Y) / 2
			if (2*centerX-c.X) >= 0 && (2*centerY-c.Y) >= 0 && (2*centerX-c.X) < pole.size && (2*centerY-c.Y) < pole.size {
				d.X = 2*centerX - c.X
				d.Y = 2*centerY - c.Y
				p.D = d
			}
		}

	}

	color = Color((d.X + d.Y) / 2)
	if color > 8 {
		color = color / 2
	}
	d.SetColor(color)
	p.D = d

	return p
}

func (p *Parallelogramm) MakeParam(pole Pole) (err error) {
	if len(pole.P) <= p.D.X || len(pole.P) <= p.D.Y || p.D.X < 0 || p.D.Y < 0 {
		fmt.Println()
		log.Fatal("Index of D.x or D.y is bigger or smaller than size of field, error was taken from MakeParam function, please, change coordinates")
	}
	d := p.D
	pole.P[d.X][d.Y] = int(d.Color)
	return
}

func (p *Parallelogramm) FindSquare(a, b, c Dot) float64 {

	firstStep := a.X*(b.Y-c.Y) - a.Y*(b.X-c.X) + b.X*c.Y - b.Y*c.X
	S := math.Abs(float64(firstStep))
	p.S = S
	return p.S
}

func (p *Parallelogramm) FindCenter(a, c Dot) Dot {

	var center Dot
	color := center.Color

	center.X = (a.X + c.X) / 2
	center.Y = (a.Y + c.Y) / 2

	color = Color((center.X + center.Y) / 2)
	center.SetColor(color)
	p.Center = center

	return p.Center
}
