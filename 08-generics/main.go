package main

import "fmt"

type AppNumber int

// constraint
type Number interface {
	// ~ -> consideramos todos os tipos e subtipos de int
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	mInt := map[string]int{
		"Wesley": 6000,
		"Tadeu":  6000,
		"Zeh":    6000,
		"Pinto":  6000,
	}

	mFloat := map[string]float64{
		"Wesley": 60.00,
		"Tadeu":  55.53,
		"Zeh":    61.84,
		"Pinto":  50.95,
	}

	mCustomNumber := map[string]AppNumber{
		"Wesley": 6000,
		"Tadeu":  55,
		"Zeh":    6184,
		"Pinto":  505,
	}

	fmt.Println(Soma(mInt))
	fmt.Println(Soma(mFloat))
	fmt.Println(Soma(mCustomNumber))
	fmt.Println(Compara("a", "b"))
}
