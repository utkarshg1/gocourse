package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")

	if err != nil {
		fmt.Println("Error opening the file :", err)
		return
	}
	defer func() {
		fmt.Println("Closing file")
		file.Close()
	}()
	fmt.Println(file)
	fmt.Println("File Opened successfully")

	// data := make([]byte, 256)
	// _, err = file.Read(data)
	// if err != nil {
	// 	fmt.Println("Error occured during read file :", err)
	// 	return
	// }
	// fmt.Println("File Content :", string(data))

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading line :", err)
		return
	}
}