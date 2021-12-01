package tests

import (
	"iTechArt_Lab/task_2_Geometry/solution"
	"testing"
)

func TestDot_SetColor(t *testing.T) {
	var dot solution.Dot

	dot.SetColor(solution.Red)
	want := solution.Red
	if dot.Color != want {
		t.Errorf("got %v, wanted %v", dot.Color, want)
	}
}
