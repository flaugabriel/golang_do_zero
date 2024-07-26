package main

import "fmt"

func main(){
	//defer funcao1()
	// funcao2()
	fmt.Println(alunoEstaAprovado(7,8))
}

func funcao1(){
	fmt.Println("função1")
}

func funcao2(){
	fmt.Println("função2")
}

func alunoEstaAprovado(n1, n2 float32) bool{
	defer fmt.Println("Média calculada. resultado será retornado")
	fmt.Println("Função aluno aprovado?")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

// esse comando defer e importante para lidar com banco de dados alua 23. Defer
// e bom para fecha conexão com banco de dados ou exectua uma consulta 
// em vez de fechar a conexão em todos os lugares so basta usar uma vez ao fechar a conexão