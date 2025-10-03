package main

import (
	"fmt"
	"math/big"
	"runtime"
	"time"
)


func calcularFatorial(n int) {
	resultado := big.NewInt(1)
	for i := 2; i <= n; i++ {
		resultado.Mul(resultado, big.NewInt(int64(i)))
	}
}

func main() {
	tarefas := []int{45000, 46000, 47000, 48000} 

	fmt.Println("CPUs disponíveis:", runtime.NumCPU())

	// Teste com apenas 1 CPU
	runtime.GOMAXPROCS(1)
	inicio := time.Now()
	for _, t := range tarefas {
		calcularFatorial(t)
	}
	fmt.Println("Tempo com 1 CPU:", time.Since(inicio))

	// Teste com todas as CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())
	inicio = time.Now()
	done := make(chan bool)

	for _, t := range tarefas {
		go func(n int) {
			calcularFatorial(n)
			done <- true
		}(t)
	}

	// Esperar todas as goroutines terminarem
	for range tarefas {
		<-done
	}
	fmt.Println("Tempo com todas as CPUs:", time.Since(inicio))
}
