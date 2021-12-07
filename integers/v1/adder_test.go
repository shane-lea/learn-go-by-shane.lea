package integers

import (
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	excepted := 4
	if sum != excepted {
		t.Errorf("except:'%d', sum:'%d'", sum, excepted)
	}
}
