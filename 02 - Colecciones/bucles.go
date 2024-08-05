package main

import "fmt" 
import	"strings"

	
func main() {

	var suma int = 0

	//Bucle for tradicional
	for i := 0; i < 10; i++ {
		suma = suma + 1
		fmt.Println(suma)
	}

	//bucle forEach no exite en go, su equivalente es el siguiente
	miMapa := map[string]string{
		"Colombia": "Bogotá",
		"Venezuela": "Caracas",
		"Brasil": "Brasilia",
		"Chile": "Santiago",
	}

	for k, v := range miMapa {
		fmt.Println("la capital de " + k + " es: " + v)
	}


	/*
		el bucle while en Go serealza de esta forma:
		se inicializa un bucle for infinito y dentro de el se utiliza un break cuando se cumple
		condicion deseada
	*/
	var fruta string  = ""
	fmt.Println("Ingrese una fruta")
	fmt.Scan(&fruta) //&fruta indica que el valor ingresado se guardara en la variable fruta
	fruta = strings.ToLower(fruta)


	for {
		if(fruta != "naranja"){
			println("No adivinaste la fruta. Ingresa nuevamente una fruta")
			fmt.Scan(&fruta) 
			fruta = strings.ToLower(fruta) 
		}else{
			fmt.Println("Adivinaste la fruta")
			break
		}
	}

}