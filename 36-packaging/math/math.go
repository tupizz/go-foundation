package math

import "fmt"

type math struct { // struct is not exported, only his constructor is exported
	numberA int // properties are not exported
	numberB int // properties are not exported
}

func NewMath(numberA int, numberB int) math {
	return math{numberA, numberB}
}

func (m math) Add() int {
	m.log("Adding...")
	return m.numberA + m.numberB
}

func (m math) Subtract() int {
	m.log("Subtracting...")
	return m.numberA - m.numberB
}

func (m math) Multiply() int {
	m.log("Multiplying...")
	return m.numberA * m.numberB
}

func (m math) log(s string) {
	fmt.Println(s)
}
