package main

import (
	"fmt"
	"sync"
	"time"
)

const filosofos = 5
const rounds = 5

const tempo_pra_pensar = 2 * time.Second
const tempo_pra_comer = 3 * time.Second

var garfos [filosofos]sync.Mutex

func Filosofo(id, iteracao int, wg *sync.WaitGroup) {
	defer wg.Done()

	// pensa
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteracao)
	time.Sleep(tempo_pra_pensar)

	// ordenação global: sempre pega o garfo de menor índice primeiro
	esquerda := id
	direita := id + 1

	// se true, completamos a volta na mesa
	if direita == filosofos {
		direita = 0 // Garfo da direita passa a ser o primeiro garfo da mesa
		esquerda, direita = direita, esquerda // inverte os garfos de mão para travar o menor primeiro
											  // e quebrar o ciclo
	}

	garfos[esquerda].Lock()
	garfos[direita].Lock()

	// come
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" E%d\n", iteracao)
	time.Sleep(tempo_pra_comer)

	garfos[direita].Unlock()
	garfos[esquerda].Unlock()
}

func main() {
	fmt.Println("\n[P1] [P2] [P3] [P4] [P5]")

	comeco := time.Now()

	for i := 1; i <= rounds; i++ {
		var wg sync.WaitGroup
		for j := 0; j < filosofos; j++ {
			wg.Add(1)
			go Filosofo(j, i, &wg)
		}
		wg.Wait()
	}

	fim := time.Since(comeco)
	fmt.Printf("\nJantar Paralelo levou %s\n\n", fim)
}
