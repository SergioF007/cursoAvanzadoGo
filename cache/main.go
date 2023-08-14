package main

import (
	"fmt"
	"sync"
	"time"
)

func Fibonacci(n int) int {

	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)

}

type Memory struct {
	f     Function
	cache map[int]FunctionResult // el cahce me va almacenar todas los Key  y sus resultados que se calculen en la funcion Function
	lock  sync.Mutex
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

// Constructor del Memory
func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

// Metodo Get de Memory
// va a recibirnos el entero y no va a devolver si ya existe un resultado para este entero en el cache
// si no existe lo va a calcular.

func (m *Memory) GetMemory(key int) (interface{}, error) {

	m.lock.Lock()
	result, exists := m.cache[key] // consulta en el cache
	m.lock.Unlock()

	if !exists {
		m.lock.Lock()
		result.value, result.err = m.f(key) // calcula el fibonacci
		m.cache[key] = result               // lo almaceno en el cache
		m.lock.Unlock()
	}

	return result.value, result.err
}

// usamos el interface{} para indicar que vamos a tener un respueta genrica
// osea que no sabemos con exactitud el tipo de resultado que vamos a obtener
func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {

	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38}

	var wg sync.WaitGroup

	for _, n := range fibo {

		wg.Add(1)

		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.GetMemory(index)

			if err != nil {

				fmt.Println(err)

			}
			fmt.Printf("%d, %s, %d\n", index, time.Since(start), value)
		}(n)

	}

	wg.Wait()

}
