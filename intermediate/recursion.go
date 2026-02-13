package intermdiate

import "fmt"

func main() {
	fmt.Println(factorial(10))
	fmt.Println(sumOfDigits(9))
	fmt.Println(sumOfDigits(35))
	fmt.Println(sumOfDigits(145))
	fmt.Println(sumOfDigits(4512))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}