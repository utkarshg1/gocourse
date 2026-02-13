package basics

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Deffered statement")
	fmt.Println("Starting main function")

	os.Exit(1)

	fmt.Println("This line will not be executed")
}