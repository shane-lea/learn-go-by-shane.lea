package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		except := 15
		if got != except {
			t.Errorf("got:'%d', want:'%d'", got, except)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		got := Sum(numbers)
		except := 6
		if got != except {

		}
	})

}
