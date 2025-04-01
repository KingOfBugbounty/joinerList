package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/schollz/progressbar/v3"
)

// Função para ler um arquivo e retornar uma lista de strings
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines, scanner.Err()
}

// Função para processar as combinações e escrever no arquivo
func processCombinations(wordlist, subdomains []string, results chan<- string, wg *sync.WaitGroup) {
	for _, word := range wordlist {
		for _, sub := range subdomains {
			results <- fmt.Sprintf("%s.%s", word, sub)
		}
	}
	wg.Done()
}

func main() {
	// Definição das flags
	wordlistFile := flag.String("wordlist", "", "Arquivo contendo a wordlist (ex: wordlist.txt)")
	subdomainsFile := flag.String("subdomains", "", "Arquivo contendo a lista de subdomínios (ex: subdomains.txt)")
	flag.Parse()

	// Verificação de parâmetros obrigatórios
	if *wordlistFile == "" || *subdomainsFile == "" {
		fmt.Println("Uso: ./uniqwordlist -wordlist wordlist.txt -subdomains subdomains.txt")
		os.Exit(1)
	}

	// Lendo as listas
	wordlist, err := readLines(*wordlistFile)
	if err != nil {
		fmt.Println("Erro ao ler", *wordlistFile, ":", err)
		return
	}

	subdomains, err := readLines(*subdomainsFile)
	if err != nil {
		fmt.Println("Erro ao ler", *subdomainsFile, ":", err)
		return
	}

	// Criando arquivo de saída
	outputFile, err := os.Create("final.txt")
	if err != nil {
		fmt.Println("Erro ao criar final.txt:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	// Definindo número de workers (goroutines controladas)
	const numWorkers = 10
	results := make(chan string, 1000)
	var wg sync.WaitGroup

	// Criando a barra de progresso
	total := int64(len(wordlist) * len(subdomains))
	bar := progressbar.Default(total)

	// Dividindo o trabalho entre os workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processCombinations(wordlist, subdomains, results, &wg)
		}()
	}

	// Goroutine para escrever no arquivo
	go func() {
		for res := range results {
			writer.WriteString(res + "\n")
			bar.Add(1)
		}
	}()

	// Esperando todos os workers terminarem
	wg.Wait()
	close(results)

	// Garantindo que todos os dados foram escritos no arquivo
	writer.Flush()

	fmt.Printf("\n[✔] Processamento concluído! %d subdomínios gerados e salvos em final.txt.\n", total)
}
