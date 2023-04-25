package main

import "fmt"

func main(){
	suma := func(a float64,b float64) float64{
		resultado := a+b
		return resultado
	}

	fmt.Println(suma(1,3))

	resta := func(a float64,b float64)float64{
		resultado := a - b
		return resultado
	}
	
	mult := func(a float64, b float64)float64{
		resultado := a * b
		return resultado
	}
	div := func(a float64, b float64) float64{
		resultado := a / b
		return resultado
	}

	fmt.Println("Resultado suma:")
	fmt.Println(suma(3,5))
	fmt.Println("Resultado resta:")
	fmt.Println(resta(3,9))
	fmt.Println("Resultado multiplicación:")
	fmt.Println(mult(3,-1))
	fmt.Println("Resultado división:")
	fmt.Println(div(1,2))

}