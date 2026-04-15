package main

import (
	"fmt"
)

func main() {
	fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")

	var n1, n2 float64
	var op string

	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&n1)

	fmt.Print("Ingresa la operación (+, -, *, /, ^, !): ")
	fmt.Scan(&op)

	if op != "!" {
		fmt.Print("Ingresa el segundo número: ")
		fmt.Scan(&n2)
	}

	switch op {
	case "+":
		fmt.Printf("Resultado: %.2f + %.2f = %.2f\n", n1, n2, n1+n2)
	case "-":
		fmt.Printf("Resultado: %.2f - %.2f = %.2f\n", n1, n2, n1-n2)
	case "*":
		fmt.Printf("Resultado: %.2f * %.2f = %.2f\n", n1, n2, n1*n2)
	case "/":
		if n2 != 0 {
			fmt.Printf("Resultado: %.2f / %.2f = %.2f\n", n1, n2, n1/n2)
		} else {
			fmt.Println("Error: no se puede dividir entre cero")
		}
	case "^":
		resultado := 1.0
		for i := 0; i < int(n2); i++ {
			resultado *= n1
		}
		fmt.Printf("Resultado: %.2f ^ %.2f = %.2f\n", n1, n2, resultado)
	case "!":
		resultado := 1.0
		num := int(n1)
		if num < 0 {
			fmt.Println("Error: No existe factorial de números negativos")
		} else {
			for i := 1; i <= num; i++ {
				resultado *= float64(i)
			}
			fmt.Printf("Resultado: %d! = %.0f\n", num, resultado)
		}
	default:
		fmt.Println("Error: operación no reconocida")
	}
}
