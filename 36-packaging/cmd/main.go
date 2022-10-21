package main

import (
	"36-packaging/math" // import the package using the module name that we created with go mod init
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	m := math.NewMath(10, 5)
	fmt.Println(m.Add())
	fmt.Println(m.Subtract())
	fmt.Println(m.Multiply())
	//m.log("Hello, World!") // this will not work because log is not exported
}
