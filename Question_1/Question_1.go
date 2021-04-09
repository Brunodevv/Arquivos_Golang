//A questão 1 sugere que dado dois registros que serão gravados em dois arquivos sequencial
//Devemos criar um terceiro arquivo que irá conter os dados em comum dos dois registros

package main

//Importação dos packages necessários para executar o código
import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Declaração do struct "Registro1"
type Registro1 struct {
	Nome     string
	Endereco string
	Telefone string
}

//Declaração do struct "Registro2"
type Registro2 struct {
	Nome      string
	Endereco  string
	Bairro    string
	Cidade    string
	CEP       string
	DataNasci string
}

//Declaração do struct "Registro3"
type Registro3 struct {
	Nome     string
	Endereco string
}

//Função para tratar erros
//Recebe como parâmetro um valor do tipo "error"
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Função para lançar as structs do tipo "Registro1" no "Arquivo1.json"
func LancarRegistro1() {

	//Abri "Arquivo1.json" para leitura e atribuí ele a variável "FileArquivo1"
	FileArquivo1, err := os.Open("Arquivo1.json")
	//Caso ocorra um erro na abertura do arquivo ele será tratado
	Check(err)

	//Adia o fechamento do arquivo até o mesmo ser utilizado
	defer FileArquivo1.Close()

	//Declara um slice do tipo "Registro1"
	var SliceRegistro1 []Registro1

	//Lê o arquivo "Arquivo1.json", criando um array de bytes deste arquivo
	//Atribuí o array de bytes a variável "byteValueRegistro1"
	byteValueRegistro1, err := ioutil.ReadAll(FileArquivo1)
	//Caso ocorra um erro na operação acima ele será reportado pela função "Check"
	Check(err)

	//Converte o array de bytes e atribuí ao slice "SliceRegistro1"
	json.Unmarshal(byteValueRegistro1, &SliceRegistro1)

	//Declaração das variável que serão usadas posteriormente no código
	reader := bufio.NewReader(os.Stdin)
	registro1 := Registro1{}
	Lancar := 0

	//"for" para atribuír dados ao registro do tipo "Registro1"
	for Lancar != 2 {
		fmt.Println("Nome: ")
		registro1.Nome, _ = reader.ReadString('\n')
		fmt.Println("Endereço: ")
		registro1.Endereco, _ = reader.ReadString('\n')
		fmt.Println("Telefone: ")
		registro1.Telefone, _ = reader.ReadString('\n')

		//Acrescenta ao "SliceRegistro1" o atual dado da variável "registro1"
		SliceRegistro1 = append(SliceRegistro1, registro1)

		fmt.Println("Deseja lançar outro registro deste tipo?", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Lancar)
	}

	//Converte os dados do "SliceRegistro1" em bytes
	DataRegistro1, err := json.MarshalIndent(SliceRegistro1, "", "  ")
	//Caso ocorra um erro na operação acima ele será tratado pela função "Check"
	Check(err)

	//Escreve no "Arquivo1.json" os bytes da variável "DataRegistro1"
	//Substítuindo os antigos bytes do arquivo por esses
	err = ioutil.WriteFile("Arquivo1.json", DataRegistro1, 0600)
	//Caso ocorra um erro ele será reportado
	Check(err)
}

//Função para lançar as structs do tipo "Registro2" no "Arquivo2.json"
func LancarRegistro2() {

	//Abri o "Arquivo2.json" para leitura e atribuí ele a variável "FileArquivo2"
	FileArquivo2, err := os.Open("Arquivo2.json")
	//Caso ocorra um erro ao abrir o arquivo ele será tratado pela função "Check"
	Check(err)

	//Adia o fechamento do arquivo até a variável "FileArquivo2" ser utilizada
	defer FileArquivo2.Close()

	//Declaração do slice "SliceRegistro2" do tipo "Registro2"
	var SliceRegistro2 []Registro2

	//Lê o arquivo "Arquivo2.json" e cria um array de bytes com os bytes que nele possuí
	//Atribuí o array de dados a variável "byteValueRegistro2"
	byteValueRegistro2, err := ioutil.ReadAll(FileArquivo2)
	//Caso ocorra um erro na operação acima ele será tratado pela função "Check"
	Check(err)

	//Converte o array de bytes e atribuí os dados a variável "SliceRegistro2"
	json.Unmarshal(byteValueRegistro2, &SliceRegistro2)

	//Declaração das variáveis que serão utilizadas posteriormente na função
	reader := bufio.NewReader(os.Stdin)
	registro2 := Registro2{}
	Lancar := 0

	//"for" para atribuír dados a variável "registro2"
	for Lancar != 2 {
		fmt.Println("Nome: ")
		registro2.Nome, _ = reader.ReadString('\n')
		fmt.Println("Endereço: ")
		registro2.Endereco, _ = reader.ReadString('\n')
		fmt.Println("Bairro: ")
		registro2.Bairro, _ = reader.ReadString('\n')
		fmt.Println("Cidade: ")
		registro2.Cidade, _ = reader.ReadString('\n')
		fmt.Println("CEP: ")
		registro2.CEP, _ = reader.ReadString('\n')
		fmt.Println("Data de nascimento: ")
		registro2.DataNasci, _ = reader.ReadString('\n')

		//Acrescenta ao slice "SliceRegistro2" os dados atuais da variável "registro2"
		SliceRegistro2 = append(SliceRegistro2, registro2)

		fmt.Println("Deseja continuar lançando registros desse tipo?", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Lancar)
	}

	//Converte os dados do slice "SliceRegostro2" em bytes
	DataRegistro2, err := json.MarshalIndent(SliceRegistro2, "", "  ")
	//Caso ocorra um erro na operação acima ele será tratado pela função "Check"
	Check(err)

	//Escreve no "Arquivo2.json" os bytes da variável "DataRegistro2"
	//Substítuindo os antigos bytes
	err = ioutil.WriteFile("Arquivo2.json", DataRegistro2, 0600)

	Check(err)
}

//Função para lançar as structs do tipo "Registro3" no "Arquivo3.json"
func LancarRegistro3() {

	//Abri o "Arquivo1.json" para leitura, atríbuindo ele a variável "FileArquivo1"
	FileArquivo1, err := os.Open("Arquivo1.json")
	//Caso ocorra um erro ao abri o arquivo ele será tratado pela função "Check"
	Check(err)

	//Adia o fechamento do arquivo até o mesmo ser utilizado
	defer FileArquivo1.Close()

	//Abri o "Arquivo2.json" para leitura, atríbuindo ele a variável "FileArquivo2"
	FileArquivo2, err := os.Open("Arquivo2.json")
	//Caso ocorra um erro ao abrir o arquivo ele será tratado pela função "Check"
	Check(err)

	//Adia o fechamento do arquivo até o memso ser utilizado
	defer FileArquivo2.Close()

	//Declaração dos slices de cado tipo registro declarado
	var SliceRegistro1 []Registro1
	var SliceRegistro2 []Registro2
	var SliceRegistro3 []Registro3

	//Lê os bytes do "Arquivo1.json" e cria um array com eles
	//Atribuí o array a variável "byteValueRegistro1"
	byteValueRegistro1, err := ioutil.ReadAll(FileArquivo1)
	//Caso ocorra um erro ao ler o arquivo ele será tratado pela função "Check"
	Check(err)
	//Lê os bytes do "Arquivo2.json" e cria um array com eles
	//Atribuí o array a variável "byteValueRegistro2"
	byteValueRegistro2, err := ioutil.ReadAll(FileArquivo2)
	//Caso ocorra um erro ao ler o arquivo ele será tratado pela função "Check"
	Check(err)

	//Converte os bytes da variável "byteValueRegistro1" e atribuí os dados convertidos a variável "SliceRegistro1"
	json.Unmarshal(byteValueRegistro1, &SliceRegistro1)
	//Converte os bytes da variável "byteValueRegistro2" e atribuí os dados convertidos a variável "SliceRegistro2"
	json.Unmarshal(byteValueRegistro2, &SliceRegistro2)

	//Verifica o tamanho do slice "SliceRegistro1" e atribuí ele a variável "TamRegistro1"
	TamRegistro1 := len(SliceRegistro1)
	//Verifica o tamanho do slice "SliceRegistro2" e atribuí ele a variável "tamRegistro2"
	TamRegistro2 := len(SliceRegistro2)
	registro3 := Registro3{}

	//"for" para avançar o slice "SliceRegistro1"
	for i := 0; i < TamRegistro1; i++ {
		//"for" para avançar o slice "SliceRegistro2"
		for j := 0; j < TamRegistro2; j++ {
			//Se o nome do  "SliceRegistro1" que está sendo lido for o mesmo do "SliceRegistro2" e o endereço também for o mesmo o seguinte será feito
			if SliceRegistro1[i].Nome == SliceRegistro2[j].Nome && SliceRegistro1[i].Endereco == SliceRegistro2[j].Endereco {
				//Registro3.nome irá receber o nome do "SliceRegistro1" que está sendo lido
				registro3.Nome = SliceRegistro1[i].Nome
				//Registro3.Endereço irá receber o endereço do "SliceRegistro1" que está sendo lido
				registro3.Endereco = SliceRegistro1[i].Endereco

				//Acrescenta ao "SliceRegistro3" o valor atual do "registro3"
				SliceRegistro3 = append(SliceRegistro3, registro3)
			}
		}
	}

	//Converte os dados da variável "SliceRegistro3" em bytes
	DataRegistro3, err := json.MarshalIndent(SliceRegistro3, "", "  ")
	//Caso ocorra um erro na operação acima ele será tratado pela função "Check"
	Check(err)

	//Escreve os bytes da variável "DataRegistro3" ao arquivo "Arquivo3.json"
	//Substítuindo os bytes que nele havia gravando por os bytes da variável
	err = ioutil.WriteFile("Arquivo3.json", DataRegistro3, 0600)
	//Caso ocorra um erro na operação a seguir ele será tratado pela função "Check"
	Check(err)
	fmt.Println("Arquivos lançados com sucesso!!")
}

func main() {
	consult := 0
	change := 0

	for consult != 2 {
		fmt.Println("O que deseja fazer?", "\n", "1-Lançar registro do tipo 1", "\n", "2-Lançar registro do tipo 2", "\n", "3-Lançar registro do tipo 3")
		fmt.Scanln(&change)

		switch {
		case change == 1:
			LancarRegistro1()
		case change == 2:
			LancarRegistro2()
		case change == 3:
			LancarRegistro3()
		}

		fmt.Println("Deseja fazer outra operação? ", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&consult)
	}
}
