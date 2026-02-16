package intermediate

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("output.txt")

	if err != nil {
		fmt.Println("Error creating file :", err)
		return
	}
	defer file.Close()

	// write data to file
	data := []byte("Welcome to Go writing a file - Utkarsh Gaikwad\n")
	n, err := file.Write(data)

	if err != nil {
		fmt.Println("Error occured while writing to file :", err)
		return
	}

	fmt.Println(n, "Bytes written to output.txt")
}