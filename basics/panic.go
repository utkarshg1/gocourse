package basics

import "fmt"

func main() {
	process(10)
	process(-3)
}

func process(i int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if i < 0 {
		panic("Negative input is not allowed")
	}
	fmt.Println("Processing value:", i)
}