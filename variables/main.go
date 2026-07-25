package main

import "fmt"

// Apenas constantes pode ser criadas e atribuidas fora de função
const PI float64 = 3.14
const TIMESLEEPINSECONDS int = 1

// Variáveis apenas podem ser declaradas
var distance float64 = 152.45

func main() {
	// : declarar variável
	// = atribuir
	// declarar e atribuir em uma linha ":="
	msg := "Hello, World!!!!!!!!!"
	fmt.Printf(msg)

	// Declarar em uma linah com tipo
	var msg2 string
	//Atribuir valor
	msg2 = "Olá Senhor!!!"
	fmt.Printf(msg2)

	// Declarar e atrivuir multiplas variáveis
	var num1, num2, num3 int = 1, 2, 3
	// Simplificanão declaração e atribuição multiplas variáveis
	num4, num5 := 4, 5

	fmt.Println(num1, num2, num3, num4, num5)

	// Linguagem fortemente tipada e estático
	//Pode definir o tipo explicito ou implicito pelo go

	var is_ok bool = true
	println(is_ok)

	print(distance)
}
