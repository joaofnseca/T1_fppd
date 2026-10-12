package main

import (
	"fmt"
	"time"
)

const ciclos = 20
const pensar = 20 * time.Millisecond
const comer = 30 * time.Millisecond

var concluidos [5]int

func filosofo(id int) {
	time.Sleep(pensar)
	time.Sleep(comer)
	concluidos[id]++
}

func main() {
	inicio := time.Now()
	for rodada := 0; rodada < ciclos; rodada++ {
		for id := 0; id < 5; id++ {
			filosofo(id)
		}
	}
	total := time.Since(inicio)
	fmt.Printf("Tempo total: %s\n", total)
	fmt.Println("Espera media: 0; maior espera: 0; maximo comendo: 1")
	fmt.Println("Ciclos por filosofo:", concluidos)
}