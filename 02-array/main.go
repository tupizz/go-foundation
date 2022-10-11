package main

import "fmt"

func main() {
	var array [3]int
	array[0] = 11
	array[1] = 99
	array[2] = 28

	fmt.Printf("último indice do array é %v \n", array[len(array)-1])

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	for i, v := range array {
		message := fmt.Sprintf("indice: %d, valor: %v", i, v)
		fmt.Println(message)
	}

}
