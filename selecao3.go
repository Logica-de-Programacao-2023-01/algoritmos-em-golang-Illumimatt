package main

import "fmt"

func main() {
	diasDaSemana := []string{"Domingo", "Segunda-feira", "Terça-feira", "Quarta-feira", "Quinta-feira", "Sexta-feira", "Sábado"}

	var num int
	valido := false

	for !valido {
		fmt.Print("Digite um número inteiro entre 1 e 7: ")
		fmt.Scanln(&num)

		if num >= 1 && num <= 7 {
			valido = true
		} else {
			fmt.Println("Erro: número inválido.")
		}
	}

	fmt.Println(diasDaSemana[num-1])
}
