package main

import "fmt"

func main() {
	//  Las estructuras se utilizan pararepresentar objetos en Go
	//  Se inicializan con la palabra 'type' seguido del nombre y 
	//  luego entre llaves los valores de la estructura

	type Person struct {
		Nombre   string
		Apellido string
		Edad int
	}
	
	func (p *Person) Saludar() {
		fmt.Println("Hola mi nombre es:", p.Nombre, p.Apellido)
	}
	

	francisco := Person{
		Nombre: "Francisco",
		Apellido: "Nessier",
		Edad: 28,
	}


	francisco.Nombre = "Francisco Pedro" // Modifico un valor de un atributo

	fmt.Println(francisco) // Verifico lo modificado
	francisco.Saludar()


	//  Los  métodos en Go no se definen dentro del struct
	//  Se declaran desde fuera con la siguiente sintaxis

}

