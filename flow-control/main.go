package main

import "fmt"

const MAX int = 1000

func main() {

	fmt.Println("1  - Oi, estou em uma execução 100% linear")
	fmt.Println("2  - Oi, estou em uma execução 100% linear")
	fmt.Println("3  - Oi, estou em uma execução 100% linear")

	var n int
	n = 10

	expr := n <= 100

	if expr {
		fmt.Printf("%d é menor do que o MAX == %d\n", n, MAX)
	}

	for i := 0; i <= MAX; i++ {
		if n%2 == 0 {
			fmt.Printf("%d é Par\n", i)
		} else {

			fmt.Printf("%d é Impar\n", i)
		}
	}

	// Em go não existe while. Apenas loop infinito de for
	for {
		n++
		fmt.Println(n)
		if n > MAX {
			break
		}
	}

	switch n < MAX {
	case true:
		{
			fmt.Printf("o Valor de N = %d é menor que %d", n, MAX)
		}
	case false:
		{
			fmt.Printf("o Valor de N = %d é mauir que %d \n", n, MAX)
		}
	}

	fruta := "banana"
	switch fruta {
	case "banana":
		{
			fmt.Println("Amarela")
		}
	default:
		{
			println("Qualquer outra")
		}
	}

	result := add(3, 4)

	fmt.Printf("Resultado: %d", result)

}

func add(a int, b int) int {
	return a + b
}
