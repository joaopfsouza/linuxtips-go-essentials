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
}
