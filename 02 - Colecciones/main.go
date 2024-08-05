package main
import "fmt"
func main() {

	//Arrays solo guardan datos del mismo tipo
	//Se debe indicar la longitud

	var primerArray [5]string

	primerArray[0] = "Fernando"
	primerArray[1] = "Lucia"
	primerArray[2] = "Nicolas"
	primerArray[3] = "Sofia"
	primerArray[4] = "Ana"
	
	//El bucle for no lleva parentesis para iniciar el contador, condicion y post
	fmt.Println("ARRAY CON BUCLE FOR")
	fmt.Println("==================")
	for i := 0 ; i< len(primerArray)-1; i++{
		fmt.Println(primerArray[i])
	}
	fmt.Println("==================")

	//Éste es el equivalente a foreach en Go
	fmt.Println("ARRAY CON ForEach")
	fmt.Println("==================")
	for _, num := range primerArray{
		fmt.Println(num)
	}
	fmt.Println("==================")

	//SLICE = Coleccion de longitud flexible con elementos del mismo tipo.
	primerSlice := [] int{12,43,210}
	primerSlice = append(primerSlice, 37) // agrego elemento al slice
	fmt.Println("SLICE CON ForEach")
	fmt.Println("==================")
	for _, num := range primerSlice{
		fmt.Println(num)
	}
	fmt.Println("==================")

	//  Para agregar variables de distinto tipo en un array o en un slice
	//  Se debe aplicar una interfaz vacía... La interfaz "envuelve" a los datos
	//  antes de ser almacenados en el array

	segSlice := []interface{}{"Ferrnando", 27,true,4.09}
	fmt.Println(segSlice)
	fmt.Println("==================")

	for _, num := range segSlice{
		fmt.Println(num)
	}


	/* M A P A S */
	//Se definen sin longitud
	fmt.Print("==============MAPAS EN GO==============")


 // Se indica que la variable es del tipo map y dentro de corchetes se escribe el tipo de dato de la clave, 
 // luego se coloca el tipo de dato del "valor" fuera de los corchetes
	mapa := map[string]int{ //clave de tipo string, valor de tipo entero
		"Hello World": 0,
		"Declaracion de variables": 1,
		"Colecciones":2,
	}

	fmt.Println(mapa)

	//de esta forma se agrega un nuevo par clave-valor:
	mapa["Estructura"] = 3
	fmt.Print(mapa)

	//de esta forma se elimina un elemento:

	delete(mapa, "Estructura")
	fmt.Println(mapa)

}