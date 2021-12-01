package solution

import "fmt"

type Pole struct {
	size int
	P    [][]int
}

func (p *Pole) CreateArray(n int) *Pole {
	p.size = n
	p.P = make([][]int, n)
	for i := range p.P {
		p.P[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			p.P[i][j] = 0
		}
	}
	return p
}

func (p Pole) PrintPole() {

	for i := 0; i < p.size; i++ {
		for j := 0; j < p.size; j++ {
			fmt.Print(p.P[i][j], " ")
		}
		fmt.Println("")
	}

}

func (p *Pole) DeleteDots() {
	for i := 0; i < len(p.P); i++ {
		for j := 0; j < len(p.P); j++ {
			p.P[i][j] = 0
		}
	}
}

func (p *Pole) SetThreeDots(a Dot, b Dot, c Dot) {

	p.P[a.X][a.Y] = int(a.Color)
	p.P[b.X][b.Y] = int(b.Color)
	p.P[c.X][c.Y] = int(c.Color)
}

func PrintInfo(par Parallelogramm, cir Circle) {
	fmt.Printf("Paralellogram dots:\nA: x:%v, y:%v, color:%v\n"+
		"B: x:%v, y:%v, color:%v\n"+
		"C: x:%v, y:%v, color:%v\n"+
		"D: x:%v, y:%v, color:%v\n"+
		"Square of Parallelogram: %v\n"+
		"Center of Paralellogram: %v\n"+
		"Length of Circle: %v\n"+
		"Center of circle: %v\n"+
		"Radius of Circle: %v\n",
		par.A.X, par.A.Y, par.A.Color,
		par.B.X, par.B.Y, par.B.Color,
		par.C.X, par.C.Y, par.C.Color,
		par.D.X, par.D.Y, par.D.Color,
		par.S, par.Center, cir.Length, cir.Center, cir.Radius)
}

func (p *Pole) ChangeOneDot(a Dot, par Parallelogramm, cir Circle) {
	p.DeleteDots()
	par.A = a
	p.SetThreeDots(a, par.B, par.C)
	par.CreateFourthDot(a, par.B, par.C, *p)
	par.MakeParam(*p)
	par.FindSquare(a, par.B, par.C)
	par.FindCenter(a, par.C)
	cir.FindRadius(par)
	cir.MakeCircle(par, *p)
	PrintInfo(par, cir)
}
