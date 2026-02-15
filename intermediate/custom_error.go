package intermediate

import (
	"errors"
	"fmt"
)

type customError struct {
	code    int
	message string
	err     error
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error : %d : %s : %v", e.code, e.message, e.err)
}

func doSomething() error {
	return &customError{
		code:    404,
		message: "Resource not found",
		err:     errors.New("Underlying error message"),
	}
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
	}

}