// task 1
package main

import "fmt"

func printtt(cur, n int) {
	if cur > n {
		return
	}
	fmt.Println(cur)
	printtt(cur+1, n)
}

// task 2

func printrange(a, b int) {
	fmt.Println(a)
	if a == b {
		return
	}
	if a < b {
		printrange(a+1, b)
	} else {
		printrange(a-1, b)
	}
}

// task 3

func sumdigits(n int) int {
	if n == 0 {
		return 0
	}
	return (n % 10) + sumdigits(n/10)
}

// task 4

func printprfactors(n, d int) {
	if d*d > n {
		if n > 1 {
			fmt.Println(n)
		}
		return
	}
	if n%d == 0 {
		fmt.Println(d)
		printprfactors(n/d, d)
	} else {
		printprfactors(n, d+1)
	}
}

func main() {

	fmt.Println("task 1:")
	printtt(1, 5)

	fmt.Println("task 2.1:")
	printrange(3, 7)

	fmt.Println("task 2.2")
	printrange(7, 3)

	fmt.Println("task 3")
	fmt.Println("Сумма цифр:", sumdigits(12345))

	fmt.Println("task 4")
	printprfactors(84, 2)
}
