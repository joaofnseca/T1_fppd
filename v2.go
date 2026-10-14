package main

import (
	"fmt"
	"sync"
	"time"
)

const filosofos = 5
const rounds = 5

const tempo_pra_pensar = 20 * time.Millisecond
const tempo_pra_comer = 30 * time.Millisecond

var garfos [filosofos]sync.Mutex
var mutex sync.Mutex

func Filosofo(id, iteracao int, wg *sync.WaitGroup) {
	defer wg.Done()

	// pensa
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteracao)
	time.Sleep(tempo_pra_pensar)

	// mutex global garante que só um filósofo pega garfos por vez
	mutex.Lock()
	esquerda := id
	direita := id + 1
	if direita == filosofos { 
		direita = 0
	}
	garfos[esquerda].Lock()
	garfos[direita].Lock()
	mutex.Unlock()

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
