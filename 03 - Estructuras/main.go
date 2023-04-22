package main

import "fmt"

func main() {
	//  Las estructuras se utilizan pararepresentar objetos en Go
	//  Se inicializan con la palabra 'type' seguido del nombre y 
	//  luego entre llaves los valores de la estructura

	type Person struct{
		nombre string
		apellido string
	}

	francisco := Person{
		nombre: "Francisco",
		apellido: "Nessier",
	}
	
	francisco.nombre = "Francisco Pedro" // Modifico un valor de un atributo

	fmt.Println(francisco) // Verifico lo modificado

}
