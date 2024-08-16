package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)
	go func() {
		escrever("GoRoutine #01")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("GoRoutine #02")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("GoRoutine #03")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("GoRoutine #04")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // Espera o Add(4) virar Add(0)
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
