package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	numbersPlus := [...]int{1, 2, 3, 4, 6}
	got := Sum(numbers)
	gotPlus := Sum(numbersPlus)
	want := 15
	if got != want {
		t.Errorf("got:'%d', want:'%d'", got, want)
	}
	if gotPlus != want {
		t.Errorf("gotPlus:'%d', want:'%d', numbersPlus:'%v'", gotPlus, want, numbersPlus)
	}
}
