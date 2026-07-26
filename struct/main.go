package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Cidade    string
}

func (p Pessoa) Ola() string {
	return "Olá, meu nome é " + p.Nome + "!"
}

func (p *Pessoa) Niver() {
	p.Idade += 1
}

func NovaPessoa(nome string, idade int, cidade string) Pessoa {
	return Pessoa{
		Nome:   nome,
		Idade:  idade,
		Cidade: cidade,
	}
}

func main() {
	var p1 Pessoa
	p1.Nome = "João"
	p1.Sobrenome = "Souza"
	p1.Idade = 36
	p1.Cidade = "Itapecerica da Serra"

	fmt.Println(p1)

	p2 := Pessoa{
		"Ricardo",
		"Almeida",
		56,
		"São Paulo"}

	fmt.Println(p2)

	p3 := Pessoa{
		Nome:   "Richard",
		Idade:  29,
		Cidade: "cidade"}

	p3.Sobrenome = "Cadorso"
	fmt.Println(p3)

	fmt.Println(p1.Ola())
	fmt.Println(p2.Ola())
	fmt.Println(p3.Ola())

	p1.Niver()
	fmt.Printf("%s agora tem %d anos.\n", p1.Nome, p1.Idade)

	p2.Niver()
	fmt.Printf("%s agora tem %d anos.\n", p2.Nome, p2.Idade)

	p3.Niver()
	fmt.Printf("%s agora tem %d anos.\n", p3.Nome, p3.Idade)

	pessoa := NovaPessoa("José", 29, "Recife")
	fmt.Println(pessoa)
}
