package operations

import (
	"errors"
	"fmt"
)

func Sum(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func Sub(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func Mul(num1 float64, num2 float64) float64 {
	return num1 * num2
}

func Div(num1 float64, num2 float64) (float64, error) {

	if num2 != 0 {
		return num1 / num2, nil
	} else {
		fmt.Println("Não pode dividir por zero")
		err := errors.New("Não pode dividir por zero")
		return 0, err
	}
}
