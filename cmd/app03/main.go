package main

import "fmt"

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

func sum(a, b int) int {
	return a + b
}

func showSum(a, b int) {
	fmt.Println(sum(a, b))
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func isPasswordValid(password string) bool {
	if len(password) < 8 {
		return false
	}
	if len(password) >= 16 {
		return false
	}
	return true
}

func calc(num int, f func(int) int) int {
	return f(num)
}

func showThreeTimes[T any](num T) {
	fmt.Println(num)
	fmt.Println(num)
}

func main() {
	fmt.Println(triple(10))
	fmt.Println(sum(3, 7))
	showSum(3, 7)

	quotient, remainder := divide(10, 3)
	fmt.Println(quotient, remainder)

	fmt.Println(isPasswordValid("abc"))
	fmt.Println(isPasswordValid("helloworld"))

	f1 := triple
	fmt.Println(f1(5))

	f2 := func(num int) int {
		return num * 3
	}
	fmt.Println(f2(5))

	fmt.Println(calc(3, double))
	fmt.Println(calc(3, triple))

	showThreeTimes(3)
	showThreeTimes(5.2)
}
