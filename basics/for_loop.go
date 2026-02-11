package basics

import "fmt"

func main() {
	for i := 1 ; i <=5 ; i++ {
		fmt.Println(i, "Hello!")
	}

	numbers := [] int {5, 6, 7, 8, 9}
	for index, value:= range numbers {
		fmt.Printf("Index : %d , Value : %d\n", index, value)
	}

	var sum = 0
	for i:=1; i<=1000; i++{
		sum += i*i
	}
	fmt.Println("Sum of square numbers from 1 to 1000:", sum)
}