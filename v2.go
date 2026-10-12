package main

import (
	"fmt"
	"sync"
	"time"
)

const philosophers = 5
const rounds = 5

const time_to_think = 2 * time.Second
const time_to_eat = 3 * time.Second

var forks [philosophers]sync.Mutex
var mutex sync.Mutex

func Philosopher(id, iteration int, wg *sync.WaitGroup) {
	defer wg.Done()

	// think
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteration)
	time.Sleep(time_to_think)

	// mutex global garante que só um filósofo pega garfos por vez
	mutex.Lock()
	left := id
	right := (id + 1) % philosophers
	forks[left].Lock()
	forks[right].Lock()
	mutex.Unlock()

	// eat
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" E%d\n", iteration)
	time.Sleep(time_to_eat)

	forks[right].Unlock()
	forks[left].Unlock()
}

func main() {
	fmt.Println("\n[P1] [P2] [P3] [P4] [P5]")

	start := time.Now()

	for i := 1; i <= rounds; i++ {
		var wg sync.WaitGroup
		for j := 0; j < philosophers; j++ {
			wg.Add(1)
			go Philosopher(j, i, &wg)
		}
		wg.Wait()
	}

	elapsed := time.Since(start)
	fmt.Printf("\nParallel Dinner took %s\n\n", elapsed)
}
