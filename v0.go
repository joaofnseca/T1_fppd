package main

import (
	"fmt"
	"time"
)

func Filosofo(id, iteracao int) {
     tempo_pra_pensar := 2
	tempo_pra_comer := 3

	// pensa
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteracao) // Pensando
	time.Sleep(time.Duration(tempo_pra_pensar) * time.Millisecond)

	// eat

	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" E%d\n", iteracao) // COmendo
	time.Sleep(time.Duration(tempo_pra_comer) * time.Millisecond)
}

func main() {
	filosofos := 5
	rounds := 5

	fmt.Println("\n[P1] [P2] [P3] [P4] [P5]\n")

	// roda filosofos

	comeco := time.Now()

	for i := 1; i <= rounds; i++ {
		for j := 0; j < filosofos; j++ {
			Filosofo(j, i)
		}
	}
	fim := time.Since(comeco)

	fmt.Printf("\nJantar Sequencial levou %s\n\n", fim)
}