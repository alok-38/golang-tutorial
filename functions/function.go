package main

import (
	"fmt"
	"log"
)

type Operation struct {
	x int
	y int
}

func (op Operation) add() int {
	return op.x + op.y
}

func (op Operation) subtract() int {
	return op.x - op.y
}

func (op Operation) multiply() int {
	return op.x * op.y
}

func (op Operation) divide() (int, error) {
	if op.y == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return op.x / op.y, nil
}

func main() {
	var x, y int
	var choice int

	for {
		fmt.Print("Enter two integers (x and y): ")
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			log.Fatal("Invalid input:", err)
		}

		op := Operation{x: x, y: y}

		fmt.Println("Choose an operation: 1) Add 2) Subtract 3) Multiply 4) Divide 5) Exit")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("%d + %d = %d\n", x, y, op.add())
		case 2:
			fmt.Printf("%d - %d = %d\n", x, y, op.subtract())
		case 3:
			fmt.Printf("%d * %d = %d\n", x, y, op.multiply())
		case 4:
			result, err := op.divide()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%d / %d = %d\n", x, y, result)
			}
		case 5:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

