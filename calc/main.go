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

	div := operations.Div(50, 2)
	fmt.Println(div)

	opdiv := op.Div(50, 2)
	fmt.Println(opdiv)

}
