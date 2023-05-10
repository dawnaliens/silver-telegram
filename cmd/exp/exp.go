package main

import (
	"fmt"
	"strings"
)

func main() {
	Demo()
	Demo(1)
	Demo(1, 2, 3)
	fmt.Println(Sum())
	fmt.Println(Sum(4))
	fmt.Println(Sum(4, 5, 6))

	fib := []int{1, 1, 2, 3, 5, 8}
	Demo(fib...)
	fmt.Println(Sum(fib...))

	strings := []string{"the", "quick", "brown", "fox"}
	fmt.Println(Join(strings...))
}

func Demo(numbers ...int) {
	for _, number := range numbers {
		fmt.Print(number, " ")
	}
	fmt.Println()
}

func Sum(numbers ...int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func Join(vals ...string) string {
	var sb strings.Builder
	for i, s := range vals {
		sb.WriteString(s)
		if i < len(vals)-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}
