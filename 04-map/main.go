package main

import "fmt"

func main() {
	salarios := map[string]float64{
		"Wesley": 1000,
		"João":   2000.49,
		"Maria":  9823.9,
	}

	//mapinit := map[string]string{}

	fmt.Println(salarios["Wesley"])

	delete(salarios, "Wesley")
	salarios["Wes"] = 5000
	fmt.Println(salarios["Wes"])
	fmt.Println(salarios["nasalkfjds"])

}
