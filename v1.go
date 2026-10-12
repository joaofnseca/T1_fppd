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

func Philosopher(id, iteration int, wg *sync.WaitGroup) {
	defer wg.Done()

	// think
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteration)
	time.Sleep(time_to_think)

	// pick up left fork, then right fork (deadlock-prone order)
	left := id
	right := (id + 1) % philosophers

	forks[left].Lock()
	fmt.Printf("Philosopher %d picked up fork %d (left), waiting for fork %d (right)\n", id, left, right)
	forks[right].Lock()

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
	fmt.Println("\n[P1] [P2] [P3] [P4] [P5]\n")

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
