package intermediate

import (
	"errors"
	"fmt"
	"math"
)

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error : %s", e.message)
}

func eprocess() error {
	return &myError{"Something went wrong"}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Square root of negative number is imaginary")
	}
	return math.Sqrt(x), nil
}

func main() {
	result, err := sqrt(12)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	} else {
		fmt.Printf("Result: %.4f\n", result)
	}

	err1 := eprocess()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
}