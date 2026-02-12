package basics

import "fmt"

func main() {
	process(10)
}

func process(i int) {
	defer fmt.Println("Deferred i value ", i)
	defer fmt.Println("First defer statement")
	defer fmt.Println("Second defer statement")
	defer fmt.Println("Third defer statement")
	fmt.Println("This is normal statement")
	i++
	fmt.Println("Value of i : ", i)
}