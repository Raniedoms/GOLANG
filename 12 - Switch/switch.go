package main

import "fmt"

func diaDaSemana (numero int) string{
	switch  numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda Feira"
	case 3:
		return "Terca Feira"
	case 4:
		return "Quinta Feira"
	case 5:
		return "Sexta Feira"
	case 6: 
		return "Sabado"
	case 7:
		return "Domingo"
	default:
		return "Numero invalido"
	}
}

func diaDaSemana2 (numero int) string{

	var diaDaSemana string
	
	switch{
	case numero == 1:
		diaDaSemana = "Segunda Feira"
		//fallthrough - ele joga seu codigo para dentro da proxima condicao
		fallthrough
	case numero == 2:
		diaDaSemana = "Terca Feira"
	case numero == 3:
		diaDaSemana = "Quarta Feira"
	case numero == 4:
		diaDaSemana = "Quinta Feira"
	case numero == 5:
		diaDaSemana = "Sexta Feira"
	case numero == 6:
		diaDaSemana = "Sabado"
	case numero == 7:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Numero Invalido"
	}

	return diaDaSemana
}



func main(){

	fmt.Println("Switch")
	dia := diaDaSemana(5)
	fmt.Println(dia)


	fmt.Println("-------")
	dia2 := diaDaSemana2(4)
	fmt.Println(dia2)

}