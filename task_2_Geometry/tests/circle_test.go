package tests

import (
	"iTechArt_Lab/task_2_Geometry/solution"
	"testing"
)

func TestCircle_SetColor(t *testing.T) {
	var circle solution.Circle

	circle.SetColor(solution.Red)
	want := solution.Red
	if circle.Color != want {
		t.Errorf("got %v, wanted %v", circle.Color, want)
	}
}

func TestCircle_FindRadius(t *testing.T) {
	var par = solution.Parallelogramm{S: 15}
	var cir solution.Circle
	cir.FindRadius(par)
	var want float64
	want = 2.1850968611841584
	if cir.Radius != want {
		t.Errorf("got %v, wanted %v", cir.Radius, want)
	}
}

func TestCircle_FindLength(t *testing.T) {

	var cir = solution.Circle{Radius: 15}
	cir.FindLength()
	var want float64
	want = 94.24777960769379
	if cir.Length != want {
		t.Errorf("got %v, wanted %v", cir.Length, want)
	}
}

func TestCircle_MakeCircle(t *testing.T) {
	var par = solution.Parallelogramm{A: solution.Dot{X: 2, Y: 3, Color: 2},
		B:      solution.Dot{X: 10, Y: 3, Color: solution.White},
		C:      solution.Dot{X: 5, Y: 2, Color: solution.Yellow},
		Center: solution.Dot{X: 2, Y: 4, Color: solution.Purple},
		S:      10}
	var pole solution.Pole
	pole.CreateArray(15)
	var cir solution.Circle
	cir.MakeCircle(par, pole)
	wantRes := 5
	if pole.P[cir.Center.X][cir.Center.Y] != wantRes {
		t.Errorf("got %v, wanted %v", pole.P[cir.Center.X][cir.Center.Y], wantRes)

	}
}
