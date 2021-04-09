package main

//Importação do packages necessários para execução do código
import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Declaração do struct "Associados"
type Associados struct {
	NumSocio       int
	Nome           string
	Endereco       string
	Bairro         string
	Cidade         string
	Estado         string
	NumDependentes int
	DataNiver      string
}

//Declaração do struct "Mensalidade"
type Mensalidade struct {
	NumSocio       int
	DataVencimento string
	DataPagamento  string
	Valor          float64
}

//Função para tratar erros
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Função para lançar os registro de associados no arquivo "Associados.json" do tipo randômico
func Lancar_Associados() {
	//Abri o arquibo "Associados.json" para leitura e atribuí ele a variável "FileAssociados"
	FileAssociados, err := os.Open("Associados.json")
	//Caso ocorra um err ao abrir o arquivo ele será tratado pela função "Check"
	Check(err)

	//Adia o fechamento do arquivo, para que ele só feche depois de ser utilazados
	defer FileAssociados.Close()

	//Abri o arquivo "Frequentadores.json" para leitura e atribuí ele a variável "FileFrequentadores"
	FileFrequentadores, err := os.Open("Frequentadores.json")
	//Caso ocorra um erro ao abrir o arquivo ele será tratado pela função "Check"
	Check(err)

	//Adia o fechamento do arquivo para que ele só feche depois de ser utilizado
	defer FileFrequentadores.Close()

	//Declaração da variável "NumFrequentadores" que irá receber o número de frequentadores
	NumFrequentadores := 0
	//Declaração do slice "SliceAssociados" do tipo "Associados"
	var SliceAssociados []Associados

	//Lê os bytes do arquibo "Frequentadores.json" e cria um array com eles
	//Atribuí o array a variável "byteValueFrequentadores"
	byteValueFrequentadores, err := ioutil.ReadAll(FileFrequentadores)
	//Caso ocorra um erro na operação acima ele será tratado pela função "Check"
	Check(err)
	//Lê os bytes do arquibo "Associados.json" e cria um array com eles
	//Atribuí a variável "byteValueAssociados" o array gerado
	byteValueAssociados, err := ioutil.ReadAll(FileAssociados)
	//Caso ocorra um erro na operação acima ele será reportado
	Check(err)

	//Converte os bytes da variável "byteValueFrequentadores" e atribuí a variável "NumFrequentadores"
	json.Unmarshal(byteValueFrequentadores, &NumFrequentadores)
	//Converte os bytes da variável "byteValueAssociados" e atrobuó a variável "SliceAssociados"
	json.Unmarshal(byteValueAssociados, &SliceAssociados)

	//Declaração das variáves que serão usadas posteriormente na função
	reader := bufio.NewReader(os.Stdin)
	associados := Associados{}
	Codigo := len(SliceAssociados)
	Lancar := 0

	//"for" para atribuir dados a variável "associados"
	for Lancar != 2 {
		associados.NumSocio = Codigo
		fmt.Println("Nome: ")
		associados.Nome, _ = reader.ReadString('\n')
		fmt.Println("Endereço: ")
		associados.Endereco, _ = reader.ReadString('\n')
		fmt.Println("Bairro: ")
		associados.Bairro, _ = reader.ReadString('\n')
		fmt.Println("Cidade: ")
		associados.Cidade, _ = reader.ReadString('\n')
		fmt.Println("Estado: ")
		associados.Estado, _ = reader.ReadString('\n')
		fmt.Println("Número de dependentes: ")
		fmt.Scanln(&associados.NumDependentes)
		fmt.Println("Data de aniversário: ")
		associados.DataNiver, _ = reader.ReadString('\n')

		//Acrescenta ao "SliceAssociados" os atuais dados da variável "associados"
		SliceAssociados = append(SliceAssociados, associados)

		//Codigo recebe um novo valor que é dado pelo comprimento do slice
		Codigo = len(SliceAssociados)

		//A variável "NumFrequentadores" recebe um novo valor
		//Que é dado pela soma do antigo valor mais o número de dependentes da variável "associados" atual mais 1
		//O valor um é o sócio que não é declarado nos dependentes mas ele também é um frequentador
		NumFrequentadores = NumFrequentadores + associados.NumDependentes + 1

		fmt.Println("Deseja continuar lançando registros de associados? ", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Lancar)
	}

	//Converte o "SliceAssociado" em bytes e atribuí a variável "DataAssociados"
	DataAssociados, err := json.MarshalIndent(SliceAssociados, "", "  ")
	//Caso ocorra um erro na operação acima ele será tratado pela função "Check"
	Check(err)

	//Converte a variável "NumFrequentadores" em bytes e atribuí eles a variável "DataFrequentadores"
	DataFrequentadores, err := json.Marshal(NumFrequentadores)
	//Caso ocorra um erro na operação acima ele será reportado
	Check(err)

	//Escreve no arquivo "Associados.json" os bytes da variável "DataAssociados"
	//Substítuindo os antigos bytes do arquibo pelos da variável
	err = ioutil.WriteFile("Associados.json", DataAssociados, 0600)
	//Caso ocorra um erro ele será tratado pela função "Check"
	Check(err)
	//Escreve no arquivo "Frequentadores.json" os bytes da variável "DataFrequentadores"
	//Substítuindo os antigos bytes do arquibo pelos da variável
	err = ioutil.WriteFile("Frequentadores.json", DataFrequentadores, 0600)
	//Caso ocorra um erro ele será tratado pela função "Check"
	Check(err)
}

//Função para lançar os registro de mensalidade no arquivo "Mensalidade.json" do tipo sequencial
func Lancar_Mensalidade() {
	//Abri o arquivo "Mensalidade.json" para leitura e atribuí ele a variável "FileMensalidade"
	FileMensalidade, err := os.Open("Mensalidade.json")
	//Caso ocorra um erro ao abrir o arquivo ele será reportado
	Check(err)

	//Adia o fechamento do arquivo para que ele só feche após ser utilizado
	defer FileMensalidade.Close()

	//Declaração do slice "SliceMensalidade" do tipo "Mensalidade"
	var SliceMensalidade []Mensalidade

	//Lê os bytes do arquibo "Mensalidade.json" e cria uma array com eles
	//Atribuí o array a variável "byteValueMensalidade"
	byteValueMensalidade, err := ioutil.ReadAll(FileMensalidade)
	//Caso ocorra um erro na operação acima ele será tratado pela função "Check"
	Check(err)

	//Converte os bytes da variável "byteValueMensalidade" e atribuí o valor convertido a variável "SliceMensalidade"
	json.Unmarshal(byteValueMensalidade, &SliceMensalidade)

	//Declaração das variáveis que serão usados posteriormente na função
	reader := bufio.NewReader(os.Stdin)
	mensalidade := Mensalidade{}
	Lancar := 0

	//"for" para atribuir dados a variável "mensalidade"
	for Lancar != 2 {
		fmt.Println("Informe o número do associado: ")
		fmt.Scanln(&mensalidade.NumSocio)
		fmt.Println("Data de vencimento: ")
		mensalidade.DataVencimento, _ = reader.ReadString('\n')
		fmt.Println("Valor: ")
		fmt.Scanln(&mensalidade.Valor)
		fmt.Println("Data de pagamento: ")
		mensalidade.DataPagamento, _ = reader.ReadString('\n')

		//Acrescenta ao "SliceMensalidade" o valor atual da variável "mensalidade"
		SliceMensalidade = append(SliceMensalidade, mensalidade)

		fmt.Println("Deseja lançar mais um registro de associado? ", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Lancar)

	}

	//Converte em bytes os dados do "SliceMensalidade" e atribuí eles a variável "DataMensalidade"
	DataMensalidade, err := json.MarshalIndent(SliceMensalidade, "", "  ")
	//Caso ocorra um erro na operação acima ele será tratado pela função "Check"
	Check(err)

	//Escreve no arquivo "Mensalidade.json" os bytes da variável "DataMensalidade"
	//Substítuindo os antigos bytes pelos novos
	err = ioutil.WriteFile("Mensalidade.json", DataMensalidade, 0600)
	//Caso ocorra um erro ele será tratado pela função "Check"
	Check(err)
}

//Função para verificar os aniversariantes do mês
func Aniversariante_do_Mes() {

	//Abri o arquivo "Associados.json" para leitura e atribuí ele a variável "FileAssociados"
	FileAssociados, err := os.Open("Associados.json")
	//Caso ocorra um erro ele será tratado pela função "Check"
	Check(err)

	//Adia o fechamento do arquivo "Associados.json" para que ele só feche após ser utilizado
	defer FileAssociados.Close()

	//Declaração do slice "SliceAssociados" do tipo Associados
	var SliceAssociados []Associados

	//Lê os bytes do arquivo "Associados.json" e cria um array com eles
	//Atribuí o array a variável "byteValueAssociados"
	byteValueAssociados, err := ioutil.ReadAll(FileAssociados)
	//Caso ocorra um erro na operação acima ele será tratado pela função "Check"
	Check(err)

	//Converte os bytes da variável "byteValueAssociados" e atribuí o valor a variável "SliceAssociados"
	json.Unmarshal(byteValueAssociados, &SliceAssociados)

	//Declaração das variáveis que serão usadas posteriormente na função
	mes := ""
	TamSliceAssociados := len(SliceAssociados)
	Aniversariantes := ""

	fmt.Println("Informe o mês atual: ")
	fmt.Scanln(&mes)

	for i := 0; i < TamSliceAssociados; i++ {
		//Se o mês da data de nascimento do registro que está sendo lido for igual ao mês inserido o seguinte vai acontecer
		if SliceAssociados[i].DataNiver[3] == mes[0] && SliceAssociados[i].DataNiver[4] == mes[1] {
			//A variável aniveráriantes irá acumular o nome do associados que faz aniverário naquele mês
			Aniversariantes = Aniversariantes + SliceAssociados[i].Nome
		}
	}

	fmt.Println("Os aniversáriantes do mês são: ", Aniversariantes)
}

//Função para alterar um registro de Associado
func Alterar_Registro() {
	//Abri o arquivo "Associados.json" para leitura e atribuí ele a variável "FileAssociados"
	FileAssociados, err := os.Open("Associados.json")
	//Caso ocorra um erro ao abrir o arquivo ele será tratado pela função "Check"
	Check(err)

	//Adia o fechamento do arquivo "Associados.json" para que ele só feche após ser utilizado
	defer FileAssociados.Close()

	//Abri o arquivo "Frequentadores.json" para leitura e atribuí ele a variável "FileFrequentadores"
	FileFrequentadores, err := os.Open("Frequentadores.json")
	//Caso ocorra um erro ao abrir o arquivo ele será tratado pela função "Check"
	Check(err)

	//Adia o fechamento do arquivo "Frequentadore.json" para que ele só feche após ser utilizado
	defer FileFrequentadores.Close()

	//Declararção da variável "frequentadores" que receberá a quantidade atual de frequentadores do clube
	frequentadores := 0
	//Declaração do slice "SliceAssociados" do tipo Associados
	var SliceAssociados []Associados

	//Lê os bytes do arquibo "Frequentadores.json" e cria um array desses bytes
	//Atribuí o array de bytes a variável "byteValueFrequentadores"
	byteValueFrequentadores, err := ioutil.ReadAll(FileFrequentadores)
	//Caso ocorra um erro na operação acima ele será tratado pela função "Check"
	Check(err)
	//Lê os bytes do arquivo "Associados.json" e cria um array desses bytes
	//Atribuí o array de bytes a variável "byteValueAssociados"
	byteValueAssociados, err := ioutil.ReadAll(FileAssociados)
	//Caso ocorra um erro na operação acima ele será tratado pela função "Check"
	Check(err)

	//Converte os bytes da variável "byteValueFrequentadores" e atribuí o valor a variável "frequentadores"
	json.Unmarshal(byteValueFrequentadores, &frequentadores)
	//Converte os bytes da variável "byteValieAssociados" e atribuí o valor ao slice "SliceAssociados"
	json.Unmarshal(byteValueAssociados, &SliceAssociados)

	//Declaração das variáveis que serão usadas posteriormente na função
	reader := bufio.NewReader(os.Stdin)
	associados := Associados{}
	CodigoAssociado := 0
	Change := 0
	choice := 0

	fmt.Println("Informe o código do úsuario que deseja alterar: ")
	fmt.Scanln(&CodigoAssociado)

	//"for" para o úsuario escolher se deseja continuar alterando o registro específicado acima
	for Change != 2 {
		fmt.Println("O que deseja alterar?", "\n", "1-Nome", "\n", "2-Endereço", "\n", "3-Bairro", "\n", "4-Cidade", "\n", "5-Estado", "\n", "6-Número de dependentes", "\n", "7-Data de aniversário ")
		//Escolhe o que deseja alterar no registro
		fmt.Scanln(&choice)

		switch {
		case choice == 1:
			fmt.Println("Informe o nome: ")
			associados.Nome, _ = reader.ReadString('\n')
			SliceAssociados[CodigoAssociado].Nome = associados.Nome
			fmt.Println("Registro alterado com sucesso!!")
		case choice == 2:
			fmt.Println("Informe o endereço: ")
			associados.Endereco, _ = reader.ReadString('\n')
			SliceAssociados[CodigoAssociado].Endereco = associados.Endereco
			fmt.Println("Registro alterado com sucesso!!")
		case choice == 3:
			fmt.Println("Informe o bairro: ")
			associados.Bairro, _ = reader.ReadString('\n')
			SliceAssociados[CodigoAssociado].Bairro = associados.Bairro
			fmt.Println("Registro alterado com sucesso!!")
		case choice == 4:
			fmt.Println("Informe a cidade: ")
			associados.Cidade, _ = reader.ReadString('\n')
			SliceAssociados[CodigoAssociado].Cidade = associados.Cidade
			fmt.Println("Registro alterado com sucesso!!")
		case choice == 5:
			fmt.Println("Informe o estado: ")
			associados.Estado, _ = reader.ReadString('\n')
			SliceAssociados[CodigoAssociado].Estado = associados.Estado
			fmt.Println("Registro alterado com sucesso!!")
		case choice == 6:
			//Subtrái na variável frequentadores a quantidade que tinha no registro específicado
			frequentadores = frequentadores - SliceAssociados[CodigoAssociado].NumDependentes - 1
			fmt.Println("Informe o número de dependentes")
			fmt.Scanln(&associados.NumDependentes)
			//Soma a variável o novo valor de associados do registro específicado
			frequentadores = frequentadores + associados.NumDependentes + 1
			SliceAssociados[CodigoAssociado].NumDependentes = associados.NumDependentes
			fmt.Println("Registro alterado com sucesso!!")
		case choice == 7:
			fmt.Println("Informe a data de aniversário: ")
			associados.DataNiver, _ = reader.ReadString('\n')
			SliceAssociados[CodigoAssociado].DataNiver = associados.DataNiver
			fmt.Println("Registro alterado com sucesso!!")
		}

		fmt.Println("Deseja alterar mais alguma informação desse associado? ", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Change)
	}

	//Converte os valores do "SliceAssociados" em bytes e atribuí a variável "DataAssociado"
	DataAssociados, err := json.MarshalIndent(SliceAssociados, "", "  ")
	//Caso ocorra um erro na operação ele será tratado na função "Check"
	Check(err)
	//Converte os valores da variável "frequentadores" em bytes e atribuí a variável "DataFrequentadores"
	DataFrequentadores, err := json.Marshal(frequentadores)
	//Caso ocorra um erro ele será tratado na função "Check"
	Check(err)

	//Escreve no arquivo "Associados.json" os bytes da variável "DataAssociados"
	//Substítuindo os bytes anteriores do arquivo
	err = ioutil.WriteFile("Associados.json", DataAssociados, 0600)
	//Caso ocorra um erro na operação acima ele será tratado na função "Check"
	Check(err)
	//Escreve no arquivo "Frequentadores.json" os bytes da variável "DataFrequentadores"
	//Substítuindo os bytes anteriores do arquivo
	err = ioutil.WriteFile("Frequentadores.json", DataFrequentadores, 0600)
	//Caso ocorra um erro ele será tratado na função "Check"
	Check(err)
}

//Função para excluir um registro de associado
func Excluir_Registro() {
	//Abri o arquivo "Associados.json" para leitura e atribuí ele a "FileAssociados"
	FileAssociados, err := os.Open("Associados.json")
	//Caso ocorra um erro ao abrir o arquivo ele será tratado pela função "Check"
	Check(err)

	//Adia o fechamento do arquivo "Associados.josn" para que ele feche só após ser utilizado
	defer FileAssociados.Close()

	//Abri o arquivo "Frequentadores.json" para leitura e atribuí ele a "FileFrequentadores"
	FileFrequentadores, err := os.Open("Frequentadores.json")
	//Caso ocorra um erro ao abrir o arquivo ele será tratado pela função "Check"
	Check(err)

	//Adia o fechamento do arquivo "Frequentadores.json" para que ele feche só após ser utlizado
	defer FileFrequentadores.Close()

	//Declaração da variável "Frequentadores" que receberá o valor atual de frequentadores
	Frequentadores := 0
	//Declaração do slice "SliceAssociados" do tipo "Associados" que receberá os registros atuais do arquivo "Associados.json"
	var SliceAssociados []Associados
	//Declaração do slice "SliceAssociadosNovo" do tipo "Associados" que receberá os registros após exclusão
	var SliceAssociadosNovo []Associados

	//Lê os bytes do arquivo "Frequentadores.json" e atribuí eles a variável "byteValueFrequentadores"
	byteValueFrequentadores, err := ioutil.ReadAll(FileFrequentadores)
	//Caso ocorra um erro na operação acima ele será tratado pela função "Check"
	Check(err)
	//Lê os bytes do arquivo "Associados.josn" e atribuí eles a variável "byteValueAssociados"
	byteValueAssociados, err := ioutil.ReadAll(FileAssociados)
	//Caso ocorra um erro na operação acima ele será tratado pela função "Check"
	Check(err)

	//Converte os bytes da variável "byteValueFrequentadores" e atribuí o valor a variável "Frequentadores"
	json.Unmarshal(byteValueFrequentadores, &Frequentadores)
	//Converte os bytes da variável "byteValueFrequentadores" e atribuí o valor ao slice "SliceAssociados"
	json.Unmarshal(byteValueAssociados, &SliceAssociados)

	//Declaração das variáveis que serão usadas posteriormente na função
	TamSliceAssociados := len(SliceAssociados)
	CodigoAssociado := 0
	CodigoNovo := 0

	fmt.Println("Informe o código do associado que deseja remover: ")
	//Informa o codigo do associado que deseja excluir
	fmt.Scanln(&CodigoAssociado)

	for i := 0; i < TamSliceAssociados; i++ {
		//Todos os registros que tiverem o código diferente do informado será atribuído ao slice "SliceAssociadosNovo"
		if SliceAssociados[i].NumSocio != CodigoAssociado {

			CodigoNovo = len(SliceAssociadosNovo)
			//Receberá um novo código de acesso já que a posição do registro foi alterada
			SliceAssociados[i].NumSocio = CodigoNovo
			SliceAssociadosNovo = append(SliceAssociadosNovo, SliceAssociados[i])

		} else {
			//Tira a quantidade de frequentadores que havia no registro do arquibo "Frequentadores.json"
			Frequentadores = Frequentadores - SliceAssociados[i].NumDependentes - 1
		}
	}

	DataAssociadosNovo, err := json.MarshalIndent(SliceAssociadosNovo, "", "  ")
	Check(err)
	DataFrequentadores, err := json.Marshal(Frequentadores)
	Check(err)

	err = ioutil.WriteFile("Associados.json", DataAssociadosNovo, 0600)
	Check(err)
	err = ioutil.WriteFile("Frequentadores.json", DataFrequentadores, 0600)
	Check(err)
}

//Função para verificar os associados inadimplentes
func Associados_Inadimplentes() {

	FileAssociados, err := os.Open("Associados.json")
	Check(err)

	defer FileAssociados.Close()

	FileMensalidade, err := os.Open("Mensalidade.json")
	Check(err)

	defer FileMensalidade.Close()

	var SliceAssociados []Associados
	var SliceMensalidade []Mensalidade

	byteValueAssociados, err := ioutil.ReadAll(FileAssociados)
	Check(err)
	byteValueMensalidade, err := ioutil.ReadAll(FileMensalidade)
	Check(err)

	json.Unmarshal(byteValueAssociados, &SliceAssociados)
	json.Unmarshal(byteValueMensalidade, &SliceMensalidade)

	TamSliceMensalidade := len(SliceMensalidade)
	Dia := 0
	Mes := 0
	ConsultDia := 0
	ConsultMes := 0
	Valor := 0.0
	Inadimplentes := ""

	fmt.Println("Informe o dia atual: ")
	fmt.Scanln(&ConsultDia)
	fmt.Println("Informe o mês atual: ")
	fmt.Scanln(&ConsultMes)

	for i := 0; i < TamSliceMensalidade; i++ {

		//Converte o dia que está no formato de string em valores inteiros
		if SliceMensalidade[i].DataVencimento[0] == 49 {
			Dia = 10 + int(SliceMensalidade[i].DataVencimento[1]) - 48
		} else if SliceMensalidade[i].DataVencimento[0] == 50 {
			Dia = 20 + int(SliceMensalidade[i].DataVencimento[1]) - 48
		} else if SliceMensalidade[i].DataVencimento[0] == 51 {
			Dia = 20 + int(SliceMensalidade[i].DataVencimento[1]) - 48
		} else {
			Dia = int(SliceMensalidade[i].DataVencimento[1]) - 48
		}

		//Converte o mês que está em formato de string em valores inteiros
		if SliceMensalidade[i].DataVencimento[3] == 49 {
			Mes = 10 + int(SliceMensalidade[i].DataVencimento[4]) - 48
		} else {
			Mes = int(SliceMensalidade[i].DataVencimento[4]) - 48
		}

		//Se o mês informado for maior que o mês de pagamento e o dia for maior ou igual e não tiver data de pagamento, o associado está inadimplente
		if Mes < ConsultMes && ConsultDia >= Dia && SliceMensalidade[i].DataPagamento == "\r\n" {
			//Acumula os nomes dos associados inadimplentes
			Inadimplentes = Inadimplentes + SliceAssociados[SliceMensalidade[i].NumSocio].Nome
			//Acumula o valor da dívida do associado
			Valor = Valor + SliceMensalidade[i].Valor
		}
	}

	fmt.Println("Os assocaidos inadimplentes são: ", Inadimplentes)
	fmt.Print("O valor total da divída é de: R$", Valor)
}

func main() {
	Consult := 0
	Change := 0

	for Consult != 2 {
		fmt.Println("O que deseja fazer? ", "\n", "1-Lançar registro de associado", "\n", "2-Lançar registro de mensalidade", "\n", "3-Ver número total de frequentadores", "\n", "4-Aniversáriantes do mês")
		fmt.Println(" 5-Incluir, Alterar ou Excluir um registro", "\n", "6-Associados Inadimplentes")
		fmt.Scanln(&Change)

		switch {
		case Change == 1:
			Lancar_Associados()
		case Change == 2:
			Lancar_Mensalidade()
		case Change == 3:
			frequentadores, err := ioutil.ReadFile("Frequentadores.json")
			Check(err)
			fmt.Print("O número total de frequentadores do clube é de: ", string(frequentadores), " Pessoas")
		case Change == 4:
			Aniversariante_do_Mes()
		case Change == 5:
			fmt.Println("1-Incluir", "\n", "2-Alterar", "\n", "3-Excluir")
			fmt.Scanln(&Change)
			switch {
			case Change == 1:
				Lancar_Associados()
			case Change == 2:
				Alterar_Registro()
			case Change == 3:
				Excluir_Registro()
			}
		case Change == 6:
			Associados_Inadimplentes()
		}

		fmt.Println("\n", "Deseja realizar mais alguma operação? ", "\n", "1-Sim", "\n", "2-Não")
		fmt.Scanln(&Consult)
	}
}
