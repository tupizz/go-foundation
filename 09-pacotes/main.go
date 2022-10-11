package main

import (
	"09-pacotes/matematica"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	fmt.Printf("My id is %v \n\n", id)

	soma := matematica.Soma(10, 20)
	fmt.Println(soma)

	i := 5
	for i > 0 {
		fmt.Println(i)
		i--
	}
}
