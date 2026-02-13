package intermediate

import "fmt"

func main() {
	message := "Hello \nGo!"
	fmt.Println(message)
	rawMessage := `Hello\nMani Meow!`
	fmt.Println(rawMessage)

	for i, c := range message{
		fmt.Printf("Index: %d, Character: %c\n", i, c)
	}

	var ch rune = 'A'
	fmt.Println(ch)
}