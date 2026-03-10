package main

import (
	"fmt"
)

func climbStairs(n int) int {
	// if n == 1 {
	// 	return 1
	// }

	// if n == 2 {
	// 	return 2
	// }
	// return climbStairs(n-1) + climbStairs(n-2)

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	steps_1 := 1
	steps_2 := 2

	for i := 2; i < n; i++ {
		steps := steps_2 + steps_1
		steps_1 = steps_2
		steps_2 = steps
	}
	return steps_2
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))

}
