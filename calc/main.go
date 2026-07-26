package main

import (
	"fmt"

	"github.com/joaopfsouza/linuxtips-go-essentials/calc/operations"

	op "github.com/joaopfsouza/linuxtips-go-essentials/calc/operations"
)

func main() {
	add := operations.Sum(5, 2)
	fmt.Println(add)

	sub := operations.Sub(5, 2)
	fmt.Println(sub)

	mult := operations.Mul(5, 2)
	fmt.Println(mult)

	div, err1 := operations.Div(50, 2)
	fmt.Println(div, err1)

	opdiv, err2 := op.Div(50, 0)
	fmt.Println(opdiv, err2)

}
