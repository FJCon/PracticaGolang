package main

import "fmt"

func main()  {

	//EN GO NO ES NECESARIO UTILIZAR ;

	//PRIMERA FORMA DE DECLARAR
	//se utiliza var + nombre de variable + tipo de dato + valor de la variable
	var nombre string = "Fernando"
	var apellido string = "Conil"
	var edad int = 26
	fmt.Println(nombre, apellido, edad)

	//SEGUNDA FORMA DE DECLARAR
	//No se utiliza var, escribimos e nombre de la variable seguido de := y el vaor de la variable
	nombre01 := "Lucia"
	apellido02 := "Schmidt"
	edad02 := 27
	fmt.Println(nombre01, apellido02, edad02)
}