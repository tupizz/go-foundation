package main

import (
	"bufio"
	"fmt"
	"os"
)

func EscreverArquivo() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Hello world! Estamos aqui escrevendo dados no arquivo arquivo.txt :D"))
	//tamanho, err := f.WriteString("hello world!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("arquivo criado com sucesso tamno %d bytes\n", tamanho)

	defer f.Close()
}

func LerArquivo() {
	conteudoArquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\no conteudo do arquivo é -> \n\t%s\n\n", string(conteudoArquivo))
}

func LeituraMemorySafe() {
	arquivo, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// precisamos criar o reader desse arquivo
	reader := bufio.NewReader(arquivo)

	// de quanto em quanto vamos ler o arquivo?
	buffer := make([]byte, 10)

	for {
		readPos, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:readPos]))
	}
}

func RemoverArquivo() {
	err := os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("\n\n------ arquivo removido =D ------\n\n")
}

func main() {
	EscreverArquivo()
	LerArquivo()
	LeituraMemorySafe()
	RemoverArquivo()
}
