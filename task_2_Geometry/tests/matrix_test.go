package tests

import (
	"iTechArt_Lab/task_2_Geometry/solution"
	"reflect"
	"testing"
)

func TestMatrix_CreateArray(t *testing.T) {
	var pole solution.Pole

	pole.CreateArray(8)

	matrix := make([][]int, 8)
	for i := range matrix {
		matrix[i] = make([]int, 8)
	}
	bo := reflect.DeepEqual(pole.P, matrix)

	if bo != true {
		t.Errorf("Something wrong")
	}
}

func TestMatrix_SetThreeDots(t *testing.T) {
	var pole solution.Pole

	var a = solution.Dot{X: 1, Y: 2, Color: solution.Yellow}
	var b = solution.Dot{X: 2, Y: 3, Color: solution.Brown}
	var c = solution.Dot{X: 1, Y: 6, Color: solution.Blue}

	pole.CreateArray(8)
	pole.SetThreeDots(a, b, c)

	matrix := make([][]int, 8)
	for i := range matrix {
		matrix[i] = make([]int, 8)
	}
	matrix[1][2] = 2
	matrix[2][3] = 7
	matrix[1][6] = 3

	bo := reflect.DeepEqual(pole.P, matrix)

	if bo != true {
		t.Errorf("Something wrong")
	}

}

func TestMatrix_ChangeOneDot(t *testing.T) {
	var pole solution.Pole
	var paral solution.Parallelogramm
	var cir solution.Circle

	pole.CreateArray(8)

	var a = solution.Dot{X: 1, Y: 2, Color: solution.Yellow}
	var b = solution.Dot{X: 2, Y: 3, Color: solution.Brown}
	var c = solution.Dot{X: 1, Y: 6, Color: solution.Blue}

	var extraDot = solution.Dot{X: 1, Y: 0, Color: solution.Green}

	pole.SetThreeDots(a, b, c)
	paral.CreateFourthDot(a, b, c, pole)
	paral.MakeParam(pole)
	paral.FindSquare(a, b, c)
	paral.FindCenter(a, c)
	cir.FindRadius(paral)
	cir.FindLength()
	cir.MakeCircle(paral, pole)

	pole.ChangeOneDot(extraDot, paral, cir)

	if pole.P[1][0] != 1 {
		t.Errorf("Something wrong")
	}

}
