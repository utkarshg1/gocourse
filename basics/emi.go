package basics

import (
	"errors"
	"fmt"
	"math"
)

func main(){
	p, err := input("Enter Principal : ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	n, err := input("Enter Number of years : ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	r, err := input("Enter Rate of interest : ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	emi, amount, interest := calculateEMI(p, n, r)
	fmt.Printf("EMI : %.2f\n", emi)
	fmt.Printf("Amount : %.2f\n", amount)
	fmt.Printf("Interest : %.2f\n", interest)
}

func input(prompt string) (float64, error) {
	var value float64
	fmt.Print(prompt)
	_, err := fmt.Scanln(&value)
	if err != nil {
		return 0, err
	}
	if value <= 0 {
		return 0, errors.New("Value must be greater than zero")
	}
	return value, nil
}

func calculateEMI(p, n, r float64) (float64, float64, float64) {
	n = n*12
	r = r/1200
	x := math.Pow(1+r, n)
	emi := p*r*x/(x-1)
	amount := emi*n
	interest := amount - p
	return emi, amount, interest
}