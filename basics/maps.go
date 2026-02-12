package basics

import "fmt"

func main() {
	myMap := make(map[string]int)
	myMap["Age"] = 30
	myMap["Score"] = 95
	myMap["Year"] = 2026
	fmt.Println(myMap)
	fmt.Println(myMap["Age"])
	fmt.Println(myMap["Year"])
	fmt.Println(myMap["test"])
}