//A questão 5 diz que temos que converter um arquivo de acesso direto em um arquivo de acesso sequêncial

package main

//Importação dos pacotes necessários para execução do código
import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Aluno struct {
	Nome string
	//Ao tirar está variável do struct não será mais possível acessar o registro diretamente
	//NumAluno  	int
	N1, N2, N3, N4 float64
}

//Função para tratar erros no algoritmo
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Função para lançar as notas
func LancarNota() {
	FileSala, err := os.Open("Sala.json")
	Check(err)

	defer FileSala.Close()

	var SliceAlunos []Aluno

	byteValueAlunos, err := ioutil.ReadAll(FileSala)
	Check(err)

	json.Unmarshal(byteValueAlunos, &SliceAlunos)

	reader := bufio.NewReader(os.Stdin)
	aluno := Aluno{}
	Lancar := 0

	for Lancar != 2 {
		fmt.Println("Nome do aluno: ")
		aluno.Nome, _ = reader.ReadString('\n')
		fmt.Println("Nota 1: ")
		fmt.Scanln(&aluno.N1)
		fmt.Println("Nota 2: ")
		fmt.Scanln(&aluno.N2)
		fmt.Println("Nota 3: ")
		fmt.Scanln(&aluno.N3)
		fmt.Println("Nota 4: ")
		fmt.Scanln(&aluno.N4)

		SliceAlunos = append(SliceAlunos, aluno)

		fmt.Println("Deseja lançar mais registros de aluno no arquivo? ", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Lancar)
	}

	DataAluno, err := json.MarshalIndent(SliceAlunos, "", "  ")
	Check(err)

	err = ioutil.WriteFile("Sala.json", DataAluno, 0600)
	Check(err)
}

func main() {
	consult := 0
	change := 0

	for consult != 2 {
		fmt.Println("Deseja lançar registros de alunos", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&change)

		switch {
		case change == 1:
			LancarNota()
		}
	}
}
