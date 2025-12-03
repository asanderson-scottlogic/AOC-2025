package main

import (
	"testing"
)

func TestDial(t *testing.T) {
	t.Run("Example from the website", func(t *testing.T) {
		part2total = 0
		dial := Dial{50}
		want := 6

		dial.Turn(left, 68)
		dial.Turn(left, 30)
		dial.Turn(right, 48)
		dial.Turn(left, 5)
		dial.Turn(right, 60)
		dial.Turn(left, 55)
		dial.Turn(left, 1)
		dial.Turn(left, 99)
		dial.Turn(right, 14)
		dial.Turn(left, 82)

		if part2total != want {
			t.Errorf("want %v got %v", want, part2total)
		}
	})
}
