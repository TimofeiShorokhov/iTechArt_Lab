package tests

import (
	"iTechArt_Lab/task_1_Basic/tasks"
	"testing"
)

func TestMatrixSum(t *testing.T) {

	got := tasks.MatrixElementsSum([][]int{{1, 1, 1, 0}, {0, 5, 0, 1}, {2, 1, 3, 10}})
	want := 9

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
