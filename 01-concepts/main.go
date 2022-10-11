package main

import "fmt"

type ID int

func main() {
	const variavelString string = "tadeu"
	const variavelInteira float64 = 56.43
	const variavelNewType ID = 1

	// formas de printar valores e tipos com a lib fmt
	fmt.Printf("o tipo de %s é %T \n", variavelString, variavelString)   // o tipo de tadeu é string
	fmt.Printf("o tipo de %v é %T \n", variavelInteira, variavelInteira) // o tipo de 56.43 é float64
	fmt.Printf("o tipo de %v é %T \n", variavelNewType, variavelNewType) // o tipo de 1 é main.ID
}
