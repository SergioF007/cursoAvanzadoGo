package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

// Vamos a aplicar composicion para representar una herencia de un clase que tiene relacion con Computer

type Laptop struct {
	Computer
}

// ya que estamos construyendo la herencia con composicion, y asi solo me permite compartir los atributos
// para representar la herencia de metodos en go vamos a crear un funcion que me retorne la interfaz que tambien implementad Computer.
// Por lo que tenemos que aplicar composicion para poder representar la herencia de metodos
func newLaptop() IProduct {

	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

// Creamos otra subclase

type Desktop struct {
	Computer
}

func newDesktop() IProduct {

	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {

	if computerType == "laptop" {
		return newLaptop(), nil
	}

	if computerType == "desktop" {
		return newDesktop(), nil
	}

	return nil, fmt.Errorf("Invalid computer type")
}

// Creo la funcion que me va imprimir la informacion de la supclase de interes.
func printNameAndStock(p IProduct) {
	fmt.Printf("Product name: %s, with stock %d\n", p.getName(), p.getStock())
}

func main() {

	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")

	printNameAndStock(laptop)
	printNameAndStock(desktop)
}