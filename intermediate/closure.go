package intermediate

import "fmt"

func main() {
	// seq := adder()

	// fmt.Println(seq())
	// fmt.Println(seq())
	// fmt.Println(seq())
	// fmt.Println(seq())
	// fmt.Println(seq())

	// seq2 := adder()
	// fmt.Println(seq2())

	subtractor := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()
	
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(2))
	fmt.Println(subtractor(3))
	fmt.Println(subtractor(4))

}

func adder() func() int {
	i := 0
	fmt.Println("Current value of i :", i)
	return func() int {
		i++
		fmt.Println("Adding 1 to i")
		return i
	}
}