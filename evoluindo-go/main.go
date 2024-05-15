package main

import (
	"fmt"
	"time"
)

func contador(x int) {
	for i := 0; i < x; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func worker(wokerId int, data chan int) {
	for x := range data { //leitura do canal
		fmt.Printf("Worker %d está processando o valor %d\n", wokerId, x)
		time.Sleep(time.Second)
	}
}

func main() {

	//Demora 30 segundos para terminar
	//contador(10)
	//contador(10)
	//contador(10)

	//Usando goroutines
	//go contador(10)
	//go contador(10)
	//go contador(10)

	//contador(10)

	//Usando goroutines
	//canal := make(chan string)
	//go func() {
	//	canal <- "Olá Mundo"
	//}()

	//mensagem := <-canal
	//fmt.Println(mensagem)

	// Melhorando o código
	canal := make(chan int)

	//go worker(1, canal) //T1 - worker 1
	//go worker(2, canal) //T2 - worker 2
	//go worker(3, canal) //T3 - worker 3
	//go worker(4, canal) //T4 - worker 4
	//Obs. Worker 1, 2, 3 e 4 estão concorrendo para pegar os dados do canal, por isso a ordem de execução é aleatória,
	//cada worker pega um valor do canal e processa, e assim por diante. Worker = Tarefa, que é uma goroutine, que é uma thread.
	qtdWorkers := 10000
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, canal)
	}

	for i := 0; i < 100000; i++ { //produzindo dados
		canal <- i
	}

}
