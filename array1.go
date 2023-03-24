package main

//Crie um Array de inteiros com 5 elementos e imprima cada elemento em uma linha
//separada.
import "fmt"

func main() {
	inteiros := [5]int{1, 2, 3, 4, 5}

	for _, inteiros := range inteiros {
		fmt.Println(inteiros)
	}

}
