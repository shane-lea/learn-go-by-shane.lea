package repeat

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	repeat := Repeat("s", 10)
	expect := "ssssssssss"

	if repeat != expect {
		t.Errorf("except:'%q', repect:'%q'", expect, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("s", 5)
	}
}

func ExampleRepeat() {
	character := Repeat("a", 10)
	fmt.Println(character)
	//Output aaaaaaaaaa
}
