package main;
import "fmt";

func main(){
	var num1, num2 float64
	var operador string
	fmt.Println("Ingrese el primer número")
	fmt.Scan(&num1)
	fmt.Println("Ingrese el segundo número")
	fmt.Scan(&num2)
	fmt.Println("Ingrese operacion (+, -, *, /)")
	fmt.Scan(&operador)

	switch operador{
		case "+": resultado := num1 + num2
		fmt.Println(resultado)
		case "-": resultado := num1 - num2
		fmt.Println(resultado)
		case "*": resultado := num1 * num2
		fmt.Println(resultado)
		case "/":
			if (num2 != 0){
				resultado := num1 / num2
				fmt.Println(resultado)
			}else{
				fmt.Println("No es posible dividir por 0")
			}
	}
}