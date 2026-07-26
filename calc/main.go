package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joaopfsouza/linuxtips-go-essentials/calc/operations"

	op "github.com/joaopfsouza/linuxtips-go-essentials/calc/operations"
)

func main() {

	// slice (array) uma lista de valores
	//args := os.Args

	ops := os.Args[1]
	s1 := os.Args[2]
	s2 := os.Args[3]

	n1, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		panic("Valor invalido")
	}

	n2, err := strconv.ParseFloat(s2, 64)
	fmt.Println(n2)
	if err != nil {
		panic("Valor invalido")
	}

	fmt.Printf("Executando a operação %s com valores %f e %f\n", ops, n1, n2)
	fmt.Println(ops)
	fmt.Println(n1)
	fmt.Println(n2)

	var result float64

	switch ops {
	case "+":
		{
			result = op.Sum(n1, n2)

		}
	case "-":
		{
			result = op.Sub(n1, n2)
		}
	case "*":
		{
			result = op.Mul(n1, n2)

		}
	case "/":
		{
			result, err = op.Div(n1, n2)

			if err != nil {
				fmt.Println(err)
				panic("Valor inválido")
			}

		}
	default:
		{
			fmt.Println("Operador inválido")
			err = errors.New("Operador inválido")
		}
	}

	if err != nil {
		fmt.Printf("O operador %s é inválido\n", ops)
	} else {

		fmt.Printf("O resultado de %f %s %f = %f\n", n1, ops, n2, result)
	}

	sayHello("Joao", "Paulo", "Ferreira")

}

func sayHello(names ...string) {
	for _, name := range names {
		fmt.Printf("Hello %s \n", name)
	}
}

func calc() {
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
