package main

import (
	"fmt"
)

func main() {
	// Paso 1: Encabezado decorativo [cite: 98]
	fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")

	// Declaración de variables para entrada del usuario [cite: 103, 126]
	var n1, n2 float64
	var op string

	// Captura de datos [cite: 121]
	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&n1) // Usamos & porque Scan necesita la dirección de memoria [cite: 127]

	fmt.Print("Ingresa la operación (+, -, *, /, ^, !): ")
	fmt.Scan(&op)

	// El factorial solo usa un número, el resto usa dos [cite: 141]
	if op != "!" {
		fmt.Print("Ingresa el segundo número: ")
		fmt.Scan(&n2)
	}

	// Estructura Switch para el control de flujo [cite: 128]
	switch op {
	case "+":
		fmt.Printf("Resultado: %.2f + %.2f = %.2f\n", n1, n2, n1+n2)
	case "-":
		fmt.Printf("Resultado: %.2f - %.2f = %.2f\n", n1, n2, n1-n2)
	case "*":
		fmt.Printf("Resultado: %.2f * %.2f = %.2f\n", n1, n2, n1*n2)
	case "/":
		// Manejo defensivo: evitar que el programa crashee [cite: 12, 131]
		if n2 != 0 {
			fmt.Printf("Resultado: %.2f / %.2f = %.2f\n", n1, n2, n1/n2)
		} else {
			fmt.Println("Error: no se puede dividir entre cero")
		}
	case "^":
		// Potencia manual con bucle for [cite: 135, 139]
		resultado := 1.0
		for i := 0; i < int(n2); i++ {
			resultado *= n1
		}
		fmt.Printf("Resultado: %.2f ^ %.2f = %.2f\n", n1, n2, resultado)
	case "!":
		// Factorial manual con bucle for [cite: 135, 143]
		resultado := 1.0
		num := int(n1)
		if num < 0 {
			fmt.Println("Error: No existe factorial de números negativos")
		} else {
			for i := 1; i <= num; i++ {
				resultado *= float64(i)
			}
			fmt.Printf("Resultado: %d! = %.0f\n", num, resultado) // Caso 0! dará 1 [cite: 142]
		}
	default:
		// Manejo de operación no válida [cite: 132]
		fmt.Println("Error: operación no reconocida")
	}
}
