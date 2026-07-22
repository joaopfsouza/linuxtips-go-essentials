# Compilling GO code

1 - GO é uma linguagem COMPILADA -  muito pareceido com C!

CÓDIGO -> COMPILADOR DO CÒDIGO -> LINGUAGEM de MÁquina -> EXECUTÁVEL (PROGRAMA)

main.go -> compilador do GO -> executável main

- Linguagens de programação podem ser interpretadas ou compiladas!
- (Não é sempre) linguagem de programação compiladas tendem a ser mais performáticas
- GO é uma linguagem multi-plataforma :
    - Facilita o compartilhanmento do programa
    - CI: Compilar
    - CD: Coloca o binário para rodar na máquina
    - Máquina não precisa ter o código fonte

Pros do go:
- Binário GO é autocontido, ou seja, quaalquer biblioteca de sistema etc, já vem dentro do binário.
- Código compilado em C nem sempre tem dentro desse código as bibliotecas (.h) que o seu código precisa para rodar dentro do executável

- `go run main.go` -> compila e roda o arquivo main.go
- `go build main.go` -> gera o binário

Arquiteturas 

- GOOS
- GOARCH

```bash

GOOS=windows GOARCH=amd64 go build  main.go
```

Create modulo

```bash
go mod init github.com/joaopfsouza/linuxtips-go-essentials/compilling-go-code
```