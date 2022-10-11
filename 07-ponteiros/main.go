package main

import "fmt"

func main() {
	var a int = 10
	var ponteiroParaA *int = nil

	ponteiroParaA = &a
	*ponteiroParaA = 10
	fmt.Printf("o valor da variabel ponteiroParaA é %v \n", ponteiroParaA)
	fmt.Printf("o valor da variabel a é %v \n", a)

	var novoPonteiroParaA *int = nil
	novoPonteiroParaA = &a
	*novoPonteiroParaA = 20
	fmt.Printf("o valor da variabel *novoPonteiroParaA é %v \n", *novoPonteiroParaA)
	fmt.Printf("o valor da variabel a é %v \n", a)
}
