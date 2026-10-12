package main

import (
	"fmt"
	"sync"
	"time"
)

const ciclos = 20
const pensar = 20 * time.Millisecond
const comer = 30 * time.Millisecond

var garfos [5]sync.Mutex
var pares = [5][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 0}}
var grupo, barreira sync.WaitGroup

func filosofo(id int) {
	primeiro, segundo := pares[id][0], pares[id][1]
	for rodada := 0; rodada < ciclos; rodada++ {
		time.Sleep(pensar)
		garfos[primeiro].Lock()
		fmt.Printf("F%d possui G%d e espera G%d\n", id, primeiro, segundo)
		barreira.Done()
		barreira.Wait()
		garfos[segundo].Lock()

		time.Sleep(comer)
		garfos[segundo].Unlock()
		garfos[primeiro].Unlock()
	}
	grupo.Done()
}

func main() {
	barreira.Add(5)
	grupo.Add(5)
	for id := 0; id < 5; id++ {
		go filosofo(id)
	}
	grupo.Wait()
}