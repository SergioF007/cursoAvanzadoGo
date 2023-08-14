package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock() // en el momento donde alguien este escribiendo datos en esta funcion, vamos a bloquiar el acceso a las mismas dado a que esta en un proceso.
	b := balance
	balance = b + amount
	lock.Unlock() // y se desbloque despues de ejecutar lo que ocurrer en las dos lineas anteriores.

}

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance   // solo estamos leyendo balance. por lo contrario en lo que pasa en Deposit.
	lock.RUnlock() // garantizo que mientras el valor termine de ejecutarse va aun mantener el estado anterior.
	return b
}

func main() {

	// vamos a representar diferentes depositos que estan ocurriendo en el balance en el mismo periodo de tiempo.

	var wg sync.WaitGroup
	var lock sync.RWMutex

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)

	}

	wg.Wait()
	fmt.Println(Balance(&lock))

}
