package main

import "fmt"

const (
	ADD uint8 = 0x01
	SUB uint8 = 0x02
	DIV uint8 = 0x03
	MUL uint8 = 0x04
)

type Operation struct {
	action uint8
	acc    float32
}

func maquineta(operation Operation, values ...float32) float32 {
	acc := operation.acc

	for _, value := range values {
		switch operation.action {
		case ADD:
			acc += value
		case SUB:
			acc -= value
		case DIV:
			acc /= value
		case MUL:
			acc *= value
		}
	}

	return acc
}

func main() {

	a := maquineta(Operation{action: ADD, acc: 1}, 2, 3)  // 1 + 2 + 3
	b := maquineta(Operation{action: SUB, acc: 9}, 2, 1)  // 9 - 2 - 1
	c := maquineta(Operation{action: DIV, acc: 6}, 2)     // 6 / 2
	d := maquineta(Operation{action: MUL, acc: 10}, 2, 5) // 10 * 2 * 5

	fmt.Println("Soma igual a", a)
	fmt.Println("Subtração igual a", b)
	fmt.Println("Divisão igual a", c)
	fmt.Println("Multiplicação igual a", d)
}
