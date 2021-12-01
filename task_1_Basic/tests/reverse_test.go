package tests

import (
	"iTechArt_Lab/task_1_Basic/tasks"
	"testing"
)

func TestReverse(t *testing.T) {

	got := tasks.ReverseParentheses("foo(bar(baz))blim")
	want := "foobazrabblim"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
