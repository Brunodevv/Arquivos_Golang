package main

//Declaração dos packages necessários para execução do código
import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Declaração do tipo "Cliente" de struct
type Cliente struct {
	Codigo   int
	Telefone string
	Nome     string
	Endereco string
	RG       string
	CPF      string
}

//Declaração do tipo "Fita" de struct
type Fita struct {
	Codigo     int
	Oscar      string
	Titulo     string
	Assunto    string
	DataCompra string
	Preco      float64
}

//Declaração do tipo "Movimento" de struct
type Movimento struct {
	CodigoFita         int
	CodigoCliente      int
	QuantidadeDiasFora int
	Preco              float64
}

//Função para tratar erros no código
//Recebe como parâmetro uma variável do tipo "error"
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Função para lançar clientes
func LancarCliente() {
	FileClientes, err := os.Open("Clientes.json")
	Check(err)

	defer FileClientes.Close()

	var SliceCliente []Cliente

	byteValueCliente, err := ioutil.ReadAll(FileClientes)
	Check(err)

	json.Unmarshal(byteValueCliente, &SliceCliente)

	reader := bufio.NewReader(os.Stdin)
	Codigo := len(SliceCliente)
	cliente := Cliente{}
	Lancar := 0

	for Lancar != 2 {
		cliente.Codigo = Codigo
		fmt.Println("Nome: ")
		cliente.Nome, _ = reader.ReadString('\n')
		fmt.Println("Telefone: ")
		cliente.Telefone, _ = reader.ReadString('\n')
		fmt.Println("Endereço: ")
		cliente.Endereco, _ = reader.ReadString('\n')
		fmt.Println("RG: ")
		cliente.RG, _ = reader.ReadString('\n')
		fmt.Println("CPF: ")
		cliente.CPF, _ = reader.ReadString('\n')

		SliceCliente = append(SliceCliente, cliente)
		Codigo = len(SliceCliente)

		fmt.Println("Deseja continuar lançando clientes? ", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Lancar)

	}

	DataCliente, err := json.MarshalIndent(SliceCliente, "", "  ")
	Check(err)

	err = ioutil.WriteFile("Clientes.json", DataCliente, 0600)
	Check(err)
}

//Função para lançar fitas
func LancarFita() {
	FileFita, err := os.Open("Fitas.json")
	Check(err)

	defer FileFita.Close()

	var SliceFita []Fita

	byteValueFita, err := ioutil.ReadAll(FileFita)
	Check(err)

	json.Unmarshal(byteValueFita, &SliceFita)

	reader := bufio.NewReader(os.Stdin)
	fita := Fita{}
	Lancar := 0
	Codigo := len(SliceFita)
	Oscar := 0

	for Lancar != 2 {
		fita.Codigo = Codigo
		fmt.Println("Título: ")
		fita.Titulo, _ = reader.ReadString('\n')
		fmt.Println("Oscar", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Oscar)
		if Oscar == 1 {
			fita.Oscar = "Sim"
		} else {
			fita.Oscar = "Não"
		}
		fmt.Println("Assunto: ")
		fita.Assunto, _ = reader.ReadString('\n')
		fmt.Println("Data de compra: ")
		fita.DataCompra, _ = reader.ReadString('\n')
		fmt.Println("Preço: ")
		fmt.Scanln(&fita.Preco)

		SliceFita = append(SliceFita, fita)
		Codigo = len(SliceFita)

		fmt.Println("Deseja continuar lançando registro de fita? ", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Lancar)
	}

	DataFita, err := json.MarshalIndent(SliceFita, "", "  ")
	Check(err)

	err = ioutil.WriteFile("Fitas.json", DataFita, 0600)
	Check(err)
}

//Função para lançar registro de movimento
func LancarMovimento() {
	FileMovimento, err := os.Open("Movimento.json")
	Check(err)

	defer FileMovimento.Close()

	var SliceMovimento []Movimento

	byteValueMovimento, err := ioutil.ReadAll(FileMovimento)
	Check(err)

	json.Unmarshal(byteValueMovimento, &SliceMovimento)

	movimento := Movimento{}
	Lancar := 0

	for Lancar != 2 {
		fmt.Println("Código da fita: ")
		fmt.Scanln(&movimento.CodigoFita)
		fmt.Println("Código do cliente: ")
		fmt.Scanln(&movimento.CodigoCliente)
		fmt.Println("Quantidade de dias fora: ")
		fmt.Scanln(&movimento.QuantidadeDiasFora)
		fmt.Println("Preço: ")
		fmt.Scanln(&movimento.Preco)

		SliceMovimento = append(SliceMovimento, movimento)

		fmt.Println("Deseja lançar mais registro de movimento?", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Lancar)
	}

	DataMovimento, err := json.MarshalIndent(SliceMovimento, "", "  ")
	Check(err)

	err = ioutil.WriteFile("Movimento.json", DataMovimento, 0600)
	Check(err)
}

//Função para verificar os nomes e o assuntos dos filmes que um determinado cliente locou
func Nome_Assunto() {
	FileCliente, err := os.Open("Clientes.json")
	Check(err)

	defer FileCliente.Close()

	FileMovimento, err := os.Open("Movimento.json")
	Check(err)

	defer FileMovimento.Close()

	FileFita, err := os.Open("Fitas.json")
	Check(err)

	defer FileFita.Close()

	var SliceCliente []Cliente
	var SliceMovimento []Movimento
	var SliceFita []Fita

	byteValueCliente, err := ioutil.ReadAll(FileCliente)
	Check(err)
	byteValueMovimento, err := ioutil.ReadAll(FileMovimento)
	Check(err)
	byteValueFita, err := ioutil.ReadAll(FileFita)
	Check(err)

	json.Unmarshal(byteValueCliente, &SliceCliente)
	json.Unmarshal(byteValueMovimento, &SliceMovimento)
	json.Unmarshal(byteValueFita, &SliceFita)

	TamSliceMovimento := len(SliceMovimento)
	CodigoCliente := 0
	Titulos := ""
	Assuntos := ""

	fmt.Println("Informe o código do cliente: ")
	fmt.Scanln(&CodigoCliente)

	//"for" para avançar o slice "SliceMovimento"
	for i := 0; i < TamSliceMovimento; i++ {
		//Caso o codigo do cliente que está sendo lido for igual ao informado
		//Novos valores vão ser atribuídos as variável "Títulos" e "Assuntos"
		if SliceMovimento[i].CodigoCliente == CodigoCliente {
			Titulos = Titulos + SliceFita[SliceMovimento[i].CodigoFita].Titulo
			Assuntos = Assuntos + SliceFita[SliceMovimento[i].CodigoFita].Assunto
		}
	}

	fmt.Print("Cliente ", SliceCliente[CodigoCliente].Nome, "já locou os respectivos títulos: ", "\n", Titulos, "E os respectivos assuntos: ", "\n", Assuntos)
}

//Função para verifiar o nome e telefone dos clientes que locaram uma determinada fita
func Nome_Telefone() {
	FileCliente, err := os.Open("Clientes.json")
	Check(err)

	defer FileCliente.Close()

	FileMovimento, err := os.Open("Movimento.json")
	Check(err)

	defer FileMovimento.Close()

	var SliceCliente []Cliente
	var SliceMovimento []Movimento

	byteValueCliente, err := ioutil.ReadAll(FileCliente)
	Check(err)
	byteValueMovimento, err := ioutil.ReadAll(FileMovimento)
	Check(err)

	json.Unmarshal(byteValueCliente, &SliceCliente)
	json.Unmarshal(byteValueMovimento, &SliceMovimento)

	TamSliceMovimento := len(SliceMovimento)
	CodigoFita := 0

	fmt.Println("Informe o código da fita: ")
	fmt.Scanln(&CodigoFita)

	//"for" para avançar o slice "SliceMovimento"
	for i := 0; i < TamSliceMovimento; i++ {
		//Se o código da fita que está sendo lida for igual ao informado o cliente que locou essa fita será apresentado
		//Juntamente com seu telefone
		if SliceMovimento[i].CodigoFita == CodigoFita {
			fmt.Println("Cliente", SliceCliente[SliceMovimento[i].CodigoCliente].Nome, "Telefone: ", SliceCliente[SliceMovimento[i].CodigoCliente].Telefone, "locou está fita")
		}
	}
}

//Função para verificar os gastos de cada sistema
func Gasto_Cliente() {
	FileCliente, err := os.Open("Clientes.json")
	Check(err)

	FileMovimento, err := os.Open("Movimento.json")
	Check(err)

	byteValueCliente, err := ioutil.ReadAll(FileCliente)
	Check(err)
	byteValueMovimento, err := ioutil.ReadAll(FileMovimento)
	Check(err)

	var SliceCliente []Cliente
	var SliceMovimento []Movimento

	json.Unmarshal(byteValueCliente, &SliceCliente)
	json.Unmarshal(byteValueMovimento, &SliceMovimento)

	TamSliceCliente := len(SliceCliente)
	TamSliceMovimento := len(SliceMovimento)

	//"for" para avançar o slice "SliceCliente"
	for i := 0; i < TamSliceCliente; i++ {
		//Zera a variável gasto toda vez que o slice "SliceCliente" avançar
		Gastos := 0.0
		//"for" para avançar o slice "SliceMovimento"
		for j := 0; j < TamSliceMovimento; j++ {
			//Se o codigo do cliente que está sendo lido pelo "SliceMovimento" for o mesmo do cliente do "SliceCliente"
			//A variável gasto vai somar o valor da fita locada pelo cliente
			if SliceMovimento[j].CodigoCliente == SliceCliente[i].Codigo {
				Gastos = Gastos + SliceMovimento[j].Preco
			}
		}

		fmt.Println("Cliente ", SliceCliente[i].Nome, "tem gasto equivalente a: R$", Gastos)
	}
}

//Função para verificar quais fitas um cliente locou mais de uma vez
func Locou_Mais_De_Uma_Vez() {
	FileCliente, err := os.Open("Clientes.json")
	Check(err)

	FileMovimento, err := os.Open("Movimento.json")
	Check(err)

	byteValueCliente, err := ioutil.ReadAll(FileCliente)
	Check(err)
	byteValueMovimento, err := ioutil.ReadAll(FileMovimento)
	Check(err)

	var SliceCliente []Cliente
	var SliceMovimento []Movimento

	json.Unmarshal(byteValueCliente, &SliceCliente)
	json.Unmarshal(byteValueMovimento, &SliceMovimento)

	TamSliceCliente := len(SliceCliente)
	TamSliceMovimento := len(SliceMovimento)

	//"for" para avançar o slice "SliceCliente"
	for i := 0; i < TamSliceCliente; i++ {
		//Zera o slice "SliceFitaCodigo" toda vez que o "SliceCliente" avançar
		var SliceFitaCodigo []int

		//"for" para avançar o slice "SliceMovimento"
		for j := 0; j < TamSliceMovimento; j++ {
			//Se o codigo do cliente lido pelo "SliceMovimento" for o mesmo que o do "SliceCliente"
			//O slice "SliceFitaCodigo" irá acrescentar o código da fita que está vinculada ao codigo do cliente
			if SliceMovimento[j].CodigoCliente == SliceCliente[i].Codigo {
				SliceFitaCodigo = append(SliceFitaCodigo, SliceMovimento[j].CodigoFita)
			}
		}

		//Se o comprimento da fita for maior do que um
		//A função "Organiza_Fita_Codigo" recebe como parâmetro o slice "SliceFitaCodigo"
		//Nesta função as fitas serão organizadas, contadas quantas vezes se repetiram e apresentadas
		if len(SliceFitaCodigo) > 1 {
			fmt.Println("\n", "Cliente", SliceCliente[i].Nome)
			Organiza_Fita_Codigo(SliceFitaCodigo)
		}

	}
}

//Função que orgazia um slice que contém os codigos das fitas repetidas que um cliente locou
//em ordem crescente, facilitando assim a contagem de quantas vezes certa fita foi locada por um cliente
func Organiza_Fita_Codigo(SliceCodigoFita []int) {
	FileFita, err := os.Open("Fitas.json")
	Check(err)

	defer FileFita.Close()

	var SliceFitas []Fita

	byteValueFita, err := ioutil.ReadAll(FileFita)
	Check(err)

	json.Unmarshal(byteValueFita, &SliceFitas)

	TamSliceFitaCodigo := len(SliceCodigoFita)
	QuantidadeVezes := 0

	//Algoritmo Insertion Sort para orgaziar o "SliceCodigoFita" em ordem crescente
	for i := 0; i < TamSliceFitaCodigo; i++ {
		k := SliceCodigoFita[i]
		j := i - 1

		for j >= 0 && k < SliceCodigoFita[j] {
			SliceCodigoFita[j+1] = SliceCodigoFita[j]
			j--
		}
		SliceCodigoFita[j+1] = k
	}

	//Variável verification irá receber o código que está na posição 0 do "SliceCodigoFita"
	verification := SliceCodigoFita[0]
	QuantidadeVezes = 1
	//"for" para verificar quantas vezes a fita se repetiu e imprimila
	for i := 1; i < TamSliceFitaCodigo; i++ {

		if SliceCodigoFita[i] != verification {

			if QuantidadeVezes > 1 {
				fmt.Println("Locou: ", SliceFitas[verification].Titulo, QuantidadeVezes, "vezes")
			}

			verification = SliceCodigoFita[i]
			QuantidadeVezes = 1

			//Irá somar a variável "QuantidadeVezes" toda vez que a fita se repetir
		} else {
			QuantidadeVezes++
		}
	}
	//Como os últimos valores não passar pela estrutura "if" que diz que se ele for diferente é para imprimir o nome da fita
	// e a quantidade de vezes, nós criamos um if do lado de fora do for para verificar se a quantidade de vezes é maior do 1 e
	//assim imprimir o nome da fita e a quantidade de vezes que se repetiu
	if QuantidadeVezes > 1 {
		fmt.Println("Locou: ", SliceFitas[verification])
		fmt.Println(QuantidadeVezes, "vezes")
	}
}

//Função para separar as fitas por assuntos e também apresentar as fitas com oscar
func Fitas_Assuntos_Fitas_Oscar() {
	FileFita, err := os.Open("Fitas.json")
	Check(err)

	defer FileFita.Close()

	var SliceFitas []Fita
	var SliceAssunto []string

	byteValueFita, err := ioutil.ReadAll(FileFita)
	Check(err)

	json.Unmarshal(byteValueFita, &SliceFitas)

	TamSliceFita := len(SliceFitas)

	//"for" que percorre o "SliceFita" pegando todos os assuntos que nele possuí mesmo que repetido
	for i := 0; i < TamSliceFita; i++ {
		SliceAssunto = append(SliceAssunto, SliceFitas[i].Assunto)
	}

	//Fornece como parâmetro para a função "Organiza_Fita_Assunto" o "SliceAssunto"
	//E recebe como retorno da função o mesmo só que organizado e sem os assuntos que se repetiram
	SliceAssunto = Organiza_Fita_Assunto(SliceAssunto)
	tamSliceAssunto := len(SliceAssunto)
	Oscar := ""

	//"for" para avançar o "SliceAssunto"
	for i := 0; i < tamSliceAssunto; i++ {
		//Variável para acumular os títulos de um determinado assunto
		//Irá ser zerado toda vez que o "SliceAssunto" avançar
		Titulo := ""
		//"for" para avançar o "SliceFita"
		for j := 0; j < TamSliceFita; j++ {
			//Se o assunto do filme que está sendo lido do "SliceFita" for o mesmo da variavém "Assunto"
			//A variável "Titulo" irá acrescentar o título do filme
			if SliceFitas[j].Assunto == SliceAssunto[i] {
				Titulo = Titulo + SliceFitas[j].Titulo
			}
			//Quando tiver no último valor do "SliceAssunto" irá ser verificado quais filmes possuem oscar
			if i == tamSliceAssunto-1 && SliceFitas[j].Oscar == "Sim" {
				Oscar = Oscar + SliceFitas[j].Titulo
			}
		}
		fmt.Println("\n", "Os filmes com o assunto: ", SliceAssunto[i], "São: ", Titulo)
	}
	fmt.Println("Os filmes prêmiados com oscar são: ", Oscar)
}

//Função que recebe um slice que contém todos os assuntos, elimina os assuntos que se repetiram
//Retorna um slice que contém todos os assuntos sem repetições
func Organiza_Fita_Assunto(SliceAssuntos []string) (Sliceassunto []string) {
	TamSliceAssunto := len(SliceAssuntos)

	//Algoritmo Isertion Sort para organizar o vetor em ordem alfabética
	//Agrupando assim os assuntos que forem iguais um atrás dos outros
	for i := 0; i < TamSliceAssunto; i++ {
		k := SliceAssuntos[i]
		j := i - 1

		for j >= 0 && k < SliceAssuntos[j] {
			SliceAssuntos[j+1] = SliceAssuntos[j]
			j--
		}
		SliceAssuntos[j+1] = k
	}

	var SliceAssuntosFinal []string

	//Variável "verification" irá receber o primeiro valor do "SliceAssuntos"
	verification := SliceAssuntos[0]

	//"for" para eliminar os assuntos que se repetirem
	for i := 0; i < TamSliceAssunto; i++ {
		if SliceAssuntos[i] != verification {
			SliceAssuntosFinal = append(SliceAssuntosFinal, verification)
			verification = SliceAssuntos[i]
		}
	}

	//retorna os valores que estão na variável "SliceFinal"
	return SliceAssuntosFinal
}

//Função que verifica quais fitas com oscar um determinado cliente locou
func Fita_Oscar_Cliente() {
	FileFita, err := os.Open("Fitas.json")
	Check(err)

	defer FileFita.Close()

	FileMovimento, err := os.Open("Movimento.json")
	Check(err)

	defer FileMovimento.Close()

	var SliceMovimento []Movimento
	var SliceFita []Fita

	byteValueMovimento, err := ioutil.ReadAll(FileMovimento)
	Check(err)
	byteValueFita, err := ioutil.ReadAll(FileFita)
	Check(err)

	json.Unmarshal(byteValueMovimento, &SliceMovimento)
	json.Unmarshal(byteValueFita, &SliceFita)

	TamSliceMovimento := len(SliceMovimento)
	CodigoCliente := 0
	Oscar := ""

	fmt.Println("Informe o código do cliente: ")
	fmt.Scanln(&CodigoCliente)

	//"for" para percorrer o "SliceMovimento"
	for i := 0; i < TamSliceMovimento; i++ {
		//Se o código de cliente do "SliceMovimento[i]" for igual ao que foi informado e a fita tiver um oscar
		//O título da fita será aramzenado na variável "Oscar"
		if SliceMovimento[i].CodigoCliente == CodigoCliente && SliceFita[SliceMovimento[i].CodigoFita].Oscar == "Sim" {
			Oscar = Oscar + SliceFita[SliceMovimento[i].CodigoFita].Titulo
		}
	}

	if Oscar != "" {
		fmt.Println("Cliente locou as seguintes fitas prêmiadas com oscar: ", Oscar)
	} else {
		fmt.Println("Cliente não locou nenhum filme prêmiado com o Oscar")
	}
}

//Função que verifica o tempo de locação de cada fita e também quais fitas já se pagaram
func Fitas_Tempo_de_Locação(Change int) {
	FileFita, err := os.Open("Fitas.json")
	Check(err)

	defer FileFita.Close()

	FileMovimento, err := os.Open("Movimento.json")
	Check(err)

	defer FileMovimento.Close()

	var SliceFita []Fita
	var SliceMovimento []Movimento

	byteValueFita, err := ioutil.ReadAll(FileFita)
	Check(err)
	byteValueMovimento, err := ioutil.ReadAll(FileMovimento)
	Check(err)

	json.Unmarshal(byteValueFita, &SliceFita)
	json.Unmarshal(byteValueMovimento, &SliceMovimento)

	TamSliceFita := len(SliceFita)
	TamSliceMovimento := len(SliceMovimento)
	Fita := ""

	//"for" para avançar o "SliceFita"
	for i := 0; i < TamSliceFita; i++ {
		//Variáveis para armazenar o tempo e o valor de cada fita
		//Serão zeradas toda vez que o "SliceFita" avançar
		Tempo := 0
		Valor := 0.0

		//"for" para avançar o "SliceMovimento"
		for j := 0; j < TamSliceMovimento; j++ {

			//Se o código do "SliceFita[i]" for igual ao código do "SliceMovimento[j]"
			//Valores serão acrescentados as seguintes variáveis
			if SliceFita[i].Codigo == SliceMovimento[j].CodigoFita {
				Fita = SliceFita[i].Titulo
				Tempo = Tempo + SliceMovimento[j].QuantidadeDiasFora
				Valor = Valor + SliceMovimento[j].Preco
			}

		}

		switch {
		case Change == 8:
			if SliceFita[i].Titulo == Fita {

				fmt.Println("Fita: ", Fita, "ficou", Tempo, "dias locada")
			}
		case Change == 9:
			if Valor > SliceFita[i].Preco {
				fmt.Println("A fita", Fita, "já se pagou")
			}
		}

	}
}

func main() {
	consult := 0
	change := 0

	for consult != 2 {
		fmt.Println("O que deseja fazer? ", "\n", "1-Lançar Cliente", "\n", "2-Lançar fita", "\n", "3-Lançar registro de movimento")
		fmt.Println(" 4-Nomes e assuntos dos filmes que um determinado cliente locou", "\n", "5-Nome e telefone dos clientes que locaram determinada fita", "\n", "6-Relatório", "\n", "7-Fitas prêmiadas com oscar que um cliente locou")
		fmt.Println(" 8-Quais fitas foram locadas e tempo de locação", "\n", "9-Fitas que já se pagaram")
		fmt.Scanln(&change)

		switch {
		case change == 1:
			LancarCliente()
		case change == 2:
			LancarFita()
		case change == 3:
			LancarMovimento()
		case change == 4:
			Nome_Assunto()
		case change == 5:
			Nome_Telefone()
		case change == 6:
			Gasto_Cliente()
			Locou_Mais_De_Uma_Vez()
			Fitas_Assuntos_Fitas_Oscar()
		case change == 7:
			Fita_Oscar_Cliente()
		case change == 8:
			Fitas_Tempo_de_Locação(change)
		case change == 9:
			Fitas_Tempo_de_Locação(change)
		}

		fmt.Println("\n", "Deseja realizar mais alguma operação? ", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&consult)
	}
}
