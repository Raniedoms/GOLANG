package main
import "fmt"
func main(){

	fmt.Println("Estrutura de Controle")
	fmt.Println("_____________________")


	numero := 20


	if numero > 15 {
		fmt.Println("Maior que 15")
	}else{
		fmt.Println("Menor que 15")
	}

	fmt.Println("_____________________")

	//quando voce cria uma variavel no ifInit voce limita ela ao escopo do IF
	
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero e maior que 0")
	} else if numero < -10{
		fmt.Println("Numero e menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}	
}