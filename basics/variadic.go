package basics

func main() {
	result := sum(1, 2, 3, 4, 5)
	println("The sum is:", result)
	result = sum(10, 20)
	println("The sum is:", result)
	result = sum(11, 12, 13)
	println("The sum is:", result)
}

func sum(nums ...int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}