package main

import (
	"fmt"
)

func main() {
	valor := 0
	for i := 0; i < 10; i++ {
		valor++
		fmt.Printf("valor +1 = %v\r\n", valor)
	}

	for {
		valor--
		fmt.Printf("valor -1 = %v\r\n", valor)

		if valor == 0 {
			break
		}
	}

	
	potato := "Batata"
	for indice, letra := range potato {
		fmt.Printf("potato[%v] = %q\r\n", indice, letra)
	}
	
	//Pode-se testar uma condicao, simulando o while(true) de outras linguagens
	teste := true
	for teste {
		fmt.Println("Vamos imprimir mais uma linha nesse for...")
		valor++
		if valor == 10 {
			teste = false
		}
	}
}