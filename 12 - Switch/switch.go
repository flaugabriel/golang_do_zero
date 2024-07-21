package main

import(
	"fmt"
)

// no switch não tem brake ele já sai altomaticamente
func main() {

	dia := diaDaSemana(3)
	fmt.Println(dia)
	dia2 := diaDaSemana(1)
	fmt.Println(dia2)
}

func diaDaSemana(numero int) string {
	var diaDasemana string

	switch numero {
	case 1: 
		diaDasemana = "Domingo"
		fallthrough // joga a prá proxima clausula
	case 2:
		diaDasemana = "Segunda-Feira"
	case 3: 
	diaDasemana = "Terça-Feira"
	case 4:
		diaDasemana = "Quarta-Feira"
	case 5:
		diaDasemana = "Quinta-Feira"
	case 6:
		diaDasemana = "sexta-Feira"
	case 7:
		diaDasemana = "Sabado"
	default:
		return "Número Inválido"
	}
	return diaDasemana
}

func diaDaSemana2(numero int) string {
	switch {
		case numero == 1: 
			return "Domingo"
		case numero == 2:
			return "Segunda-Feira"
		case numero == 3: 
			return "Terça-Feira"
		case numero == 4:
			return "Quarta-Feira"
		case numero == 5:
			return "Quinta-Feira"
		case numero == 6:
			return "sexta-Feira"
		case numero == 7:
			return "Sabado"
		default:
			return "Número Inválido"
		}
}