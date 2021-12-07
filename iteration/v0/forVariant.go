package main

import "fmt"

func forVariantOne() {
	i := 0
	for i < 3 {
		fmt.Println("forVariantOne >>", i)
		i++
	}
}

func forVariantTwo() {
	for i := 5; i < 9; i++ {
		fmt.Print("forVariantTwo >>")
		fmt.Println(i)
	}
}

func main() {
	forVariantOne()
	forVariantTwo()

	for {
		fmt.Println("break >>  loop")
		break
	}

	for i := 0; i <= 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print("continue >>")
		fmt.Println(i)
	}
}
