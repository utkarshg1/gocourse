package intermediate

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	go printLetters()
	go printNumbers()
	time.Sleep(2 * time.Second)
}

func sayHello() {
	time.Sleep(time.Second)
	fmt.Println("Hello from GO!")
}

func printNumbers() {
	for i := 0 ; i <= 5 ; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde"{
		fmt.Println(string(letter))
		time.Sleep(200 * time.Millisecond)
	}
}