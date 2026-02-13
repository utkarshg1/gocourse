package intermediate

import "fmt"

func main() {
	var ptr *int

	var a int = 100

	ptr = &a

	fmt.Println("Address of a :", &a)
	fmt.Println("Value of a :",a)
	fmt.Println("Address of ptr :", ptr)
	fmt.Println("Value of ptr :", *ptr)

	modifyValue(ptr)
	fmt.Println("Modified value of a:", a)
}

func modifyValue(ptr *int) {
	*ptr++
}