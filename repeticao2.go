package main

import "fmt"

//Faça um algoritmo que imprima os múltiplos de 3 de 0 a 30.

func main() {
	for num := 0; num < 31; num++ {
		if num%3 == 0 {
			fmt.Println(num, "é divisivel por 3")
		} else {
			fmt.Println(num, "não é divisivel por 3")
		}
	}
}
