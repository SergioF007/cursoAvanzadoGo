package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFibonacci(n int) int {

	fmt.Printf("Calculate Expensive Fibonacci for %d\n", n)
	time.Sleep(5 * time.Second)

	return n
}

type Service struct {
	InProgress map[int]bool       //almacena los numeros de los caules vamos a calcular la serie fibonacci he idicar en que estado estan
	IdPending  map[int][]chan int // vamos a mapiar las llaves de tipo entero a un slice de canales
	Lock       sync.RWMutex
}

func (s *Service) Work(job int) {

	s.Lock.RLock()
	exists := s.InProgress[job] // esto nos va a indicar si el job esta siendo procesado o no
	if exists {

		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)

		s.Lock.Lock()
		s.IdPending[job] = append(s.IdPending[job], response)
		// asi que por el canal de response en por donde se le va a cuminicar
		// al worket exists que ya se a terminado de calcular la serie de fibonacci
		s.Lock.Unlock()

		fmt.Printf("Waiting for Response job: %d\n", job)
		resp := <-response
		fmt.Printf("Response Done, received %d\n", resp)

		return
	}

	// cuando aun no esta en progreso, Generamos el bloqueo para empezar el proceso
	s.Lock.RUnlock()

	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Calculate Fibonacci for %d\n", job)
	result := ExpensiveFibonacci(job)

	// ya que calculamos la serie fibonacci guadamos su resultado y lo almacenamos en result
	// para consultar los datos vamos, bloquiamos y habilitamos la lectura
	s.Lock.RLock()
	// traemos pos Workers que estaba esperando los resultados que la funcion a calculado y si existen o no
	pendingWorkers, exists := s.IdPending[job]
	s.Lock.RUnlock()

	if exists {
		for _, pendingWorker := range pendingWorkers {
			pendingWorker <- result // le notificamos a todos los workers que estan pendiente que su resultado ya a sido calculado y se les esta enviando
		}

		fmt.Printf("Result sent - all peding workers ready job: %d\n", job)

	}

	s.Lock.Lock()
	// los estoy configurando a su estado inicial
	s.InProgress[job] = false // va hacer falso por que ya fue calculado
	s.IdPending[job] = make([]chan int, 0)
	s.Lock.Unlock()

}

// Creamos el constructor
func NewService() *Service {

	return &Service{
		InProgress: make(map[int]bool),
		IdPending:  make(map[int][]chan int),
	}
}

func main() {

	service := NewService()
	jobs := []int{3, 4, 5, 5, 4, 8, 8, 8}
	var wg sync.WaitGroup
	wg.Add(len(jobs)) // le agrego al contador la longitud de los jobs que tenemos
	for _, n := range jobs {
		go func(job int) {

			defer wg.Done()
			service.Work(job)

		}(n)
	}

	wg.Wait()

}
