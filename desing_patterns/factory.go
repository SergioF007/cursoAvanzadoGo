package main

import "fmt"

// Creo el contrato de metodos que tiene que implemetar Computer
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

// Se cumple con el contrato al implemetar todos los metodos por lo que
// el struc Computer esta usando de manera implisita la interfaz
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
} // es fue un ejemplo como Lapto puede implementar los atributos de Computer

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

// Aqui instaciamos nuestros objetos(Subclases)
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
// aqui estamos aplicando un metodo polimorfo en nuestras objetos
// usea, se puede implementar a ambos objetos ya que su logica en este caso se quiere que sea compartida
func printNameAndStock(p IProduct) {
	fmt.Printf("Product name: %s, with stock %d\n", p.getName(), p.getStock())
}

func main() {

	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")

	printNameAndStock(laptop)
	printNameAndStock(desktop)
}
