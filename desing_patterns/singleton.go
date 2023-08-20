package main

import (
	"fmt"
	"sync"
	"time"
)

type DataBase struct{}

// Vamos a crear la funcion de conexion de la DataBase
func (DataBase) CreateSingleConnection() {
	fmt.Println("Creating Singleton for DataBase")
	time.Sleep(2 * time.Second)
	fmt.Println("Creation Done")
}

var db *DataBase
var lock sync.Mutex

func getDataBaseIntance() *DataBase {
	// implementamos estas dos lineas para corregir el error que se causa en las goroutines al no
	// garantizar que se alla terminado un proceso para comenzar el otro, ya que solo necesitamos que con la primera ya pueda crear la conexion
	lock.Lock()
	defer lock.Unlock()
	// estamos validadno si ya hay una instancia creada, por lo que si no la hay, sera igual a nil y toca crearla
	if db == nil {
		fmt.Println("Creating DB Connection")
		db = &DataBase{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB Already Created")
	}

	return db

}

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDataBaseIntance()
		}()
	}
	wg.Wait()
}
