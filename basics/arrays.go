package basics

import "fmt"

func main() {
	var numbers [5]int;
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Cherry", "Date"}
	fmt.Println(fruits)
}