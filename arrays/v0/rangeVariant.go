package main

import "fmt"

func RangeVariant(numbers [5]int) (sum int) {

	for _, number := range numbers {
		sum += number
	}

	for i, number := range numbers {
		if 5 == number {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	return
}

func main() {
	fmt.Println(RangeVariant([5]int{1, 2, 3, 4, 5}))
}
