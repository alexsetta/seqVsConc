package main

import (
	"fmt" 
	"net/http"
	"time"
)

// <-chan - canal somente-leitura
func siteStatusCodeConc(urls ...string) <-chan string {
	// cria um canal de string
	c := make(chan string)

	// laço para cada url passada por parâmetro
	for _, url := range urls {
		// chama uma goroutine que verifica se o site está online ou não
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			c <- url + " - " + resp.Status
		}(url)
	}
	return c
}

// <-chan - canal somente-leitura
func siteStatusCodeSeq(urls ...string) []string{
	// cria um slice vazio para retorno
	result := make([]string, 0)
	
	// laço para cada url passada por parâmetro
	for _, url := range urls {
		// chama uma função que verifica se o site está online ou não
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		result = append(result, url + " - " + resp.Status)
	}
	return result
}

func main() {
	startTime := time.Now()
	resultSeq := siteStatusCodeSeq("https://www.amazon.com", "https://www.correios.com.br", "https://www.google.com", "https://www.youtube.com")
	fmt.Println("Primeiro:", resultSeq[0])
	fmt.Println("Segundo:", resultSeq[1])
	fmt.Println("Terceiro:", resultSeq[2])
	fmt.Println("Quarto:", resultSeq[3])
	
    diff := time.Now().Sub(startTime)
    fmt.Printf("Tempo total sequencial: %.2f seconds\n\n", diff.Seconds())

	startTime = time.Now()
	resultConc := siteStatusCodeConc("https://www.amazon.com", "https://www.correios.com.br", "https://www.google.com", "https://www.youtube.com")
	fmt.Println("Primeiro:", <-resultConc)
	fmt.Println("Segundo:", <-resultConc)
	fmt.Println("Terceiro:", <-resultConc)
	fmt.Println("Quarto:", <-resultConc)
	
	diff = time.Now().Sub(startTime)
	fmt.Printf("Tempo total concorrente: %.2f seconds\n\n", diff.Seconds())
}
