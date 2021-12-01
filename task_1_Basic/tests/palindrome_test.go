package tests

import (
	"iTechArt_Lab/task_1_Basic/tasks"
	"testing"
)

func TestPalindrome(t *testing.T) {

	got := tasks.IsPalindrome("abba")
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
