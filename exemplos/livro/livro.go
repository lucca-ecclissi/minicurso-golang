package main

import (
	"fmt"
)

type Livro struct {
	titulo string
	preco float64
	numeroPaginas int
}

func (l *Livro) Digitar(){
	fmt.Print("Entre com o Titulo: ")
	fmt.Scanf("%s", &l.titulo)
	
	fmt.Print("Entre com o Preço: ")
	fmt.Scanf("%f", &l.preco)
	
	fmt.Print("Entre com o Numero de Páginas: ")
	fmt.Scanf("%d", &l.numeroPaginas)
}

func (l Livro) Mostrar(){
	fmt.Printf("Titulo: %v ",l.titulo)
	fmt.Printf(" Preço: %v ",l.preco)
	fmt.Printf(" Numero de Páginas: %v ",l.numeroPaginas)
}

func main(){
	defer fmt.Print("\n\t FIM DA EXECUÇÃO")

	var rotate string
	
	livro1 := Livro{}
	
	livro1.Digitar()
	
	if livro1.numeroPaginas >= 100 {
		fmt.Printf("Eita quantas Páginas \n")
	}
	
	for _, c := range livro1.titulo {
		rotate = string(c) + rotate
	}

	fmt.Printf("Seu titulo ao contrário é %s \n",rotate)
	
	livro1.Mostrar()
}	