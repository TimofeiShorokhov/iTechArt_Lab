package tests

import (
	"fmt"
	"iTechArt_Lab/task_2_Geometry/solution"
	"testing"
)

func TestParallelogramm_CreateFourthDot(t *testing.T) {
	var a = solution.Dot{X: 1, Y: 2, Color: solution.Yellow}
	var b = solution.Dot{X: 2, Y: 3, Color: solution.Brown}
	var c = solution.Dot{X: 1, Y: 6, Color: solution.Blue}
	var d = solution.Dot{X: 0, Y: 5, Color: solution.Yellow}
	var center = solution.Dot{X: 1, Y: 4, Color: solution.Yellow}

	var pole solution.Pole
	pole.CreateArray(8)

	var paral solution.Parallelogramm
	var paral1 = solution.Parallelogramm{A: a, B: b, C: c, D: d, Center: center, S: 4}

	pole.SetThreeDots(a, b, c)
	paral.CreateFourthDot(a, b, c, pole)
	paral.FindCenter(a, c)
	paral.FindSquare(a, b, c)

	fmt.Println(paral)
	fmt.Println()
	fmt.Println(paral1)

	if paral != paral1 {
		t.Errorf("Something wrong")
	}

}
