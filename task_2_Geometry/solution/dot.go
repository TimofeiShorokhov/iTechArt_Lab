package solution

type Color int

const (
	Red   Color = 1
	Green       = iota
	Yellow
	Blue
	Purple
	Black
	White
	Brown
)

type Dot struct {
	X     int
	Y     int
	Color Color
}

func (d *Dot) SetColor(c Color) {
	d.Color = c
}
