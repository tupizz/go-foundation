package main

import "fmt"

func main() {
	// slices tem um ponteiro para o array, um tamanho e uma capacidade
	// por debaixo dos panos estamos trabalhando com arrays
	// uma forma dinâmica de trabalhar com arrays

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d valor=%v", len(slice), cap(slice), slice)
}
