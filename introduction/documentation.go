package main

import (
	"fmt"
)

func main() {

	// Hello, World
	fmt.Println("Hello, World!")

	// 1. Declaración de variables -----------
	var a int = 10
	var b float64 = 3.14
	var c string = "Hello, Golang!"
	var d bool = true

	// Declaración corta (infiriendo el tipo)
	e := "Inferido como string"

	fmt.Println(a, b, c, d, e)

	// 2. Arreglo (tamaño fijo) --------------
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Arreglo:", arr)

	// Slice (sin tamaño fijo, dinámico)
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	slice = append(slice, 8) // Añadir elemento

	fmt.Println("Slice:", slice)

	// 3. Mapas -------------------------------
	m := make(map[string]int)
	m["uno"] = 1
	m["dos"] = 2

	fmt.Println("Mapa:", m)
	fmt.Println("Valor de 'uno':", m["uno"])

	// 4. Condicionales -----------------------
	x := 10

	// if-else
	if x > 5 {
		fmt.Println("x es mayor que 5")
	} else if x == 5 {
		fmt.Println("x es igual que 5")
	} else {
		fmt.Println("x es menor que 5")
	}

	// Switch
	day := "lunes"
	switch day {
	case "lunes":
		fmt.Println("Es el primer día de la semana")
	case "martes":
		fmt.Println("Es el segundo día de la semana")
	case "miercoles":
		fmt.Println("Es el tercer día de la semana")
	default:
		fmt.Println("No está dentro de los primeros 3 días de la semana")
	}

	// 5. Bucles -------------------------------
	// Bucle for tradicional
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Bucle for como while
	x = 0
	for x < 5 {
		fmt.Println("x:", x)
		x++
	}

	// Iterar sobre un slice
	sli := []string{"a", "b", "c"}
	for i, v := range sli {
		fmt.Printf("Indice: %d, Valor: %s\n", i, v)
	}

	// 6 Funciones ------------------------------
	resultado := sumar(10, 5)
	cociente, resto := dividir(10, 3)

	fmt.Println("Suma:", resultado)
	fmt.Println("Cociente:", cociente, "Resto:", resto)

	// 7. Estructuras y métodos (struct)
	persona := Persona{"Juan", 30}
	persona.saludar()

	// 8. Punteros
	a = 10
	incrementar(&a)
	fmt.Println("Valor incrementado:", a) // Output: 11
}

// 6 Funciones ------------------------------
// Simple
func sumar(a int, b int) int {
	return a + b
}

// Función con múltiples valores de retorno
func dividir(a int, b int) (int, int) {
	return a / b, a % b
}

// 7. Estructuras y métodos (struct)
// Definición
type Persona struct {
	nombre string
	edad   int
}

// Método asociado a la estructura Persona
func (p Persona) saludar() {
	fmt.Printf("Hola, soy %s y tengo %d años\n", p.nombre, p.edad)
}

// 8. Punteros
// Función que modifica el valor mediante un puntero
func incrementar(p *int) {
	*p = *p + 1
}
