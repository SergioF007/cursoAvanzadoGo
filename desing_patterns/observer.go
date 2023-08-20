package main

import "fmt"

type Topic interface {
	register(observer Observer)
	// esto notifica a los observadores que ha ocurrido el evento que se estaba
	//esperando que ocurriera
	broadCast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

// Item -> No Disponible en primera instacia
// Item -> Avise -> HAY ITEM

// este Item struct es el objeto que representa el estado del observable
type Item struct {
	observers []Observer // es la variable que se crea para que este observando la disponibilidad del producto
	name      string
	available bool // esta variable indica cuando hay o no hay disponinlidad del item
}

// constructor de Item
func NewItem(name string) *Item {

	return &Item{
		name: name,
	}
}

// esta funcion me va a actualiar la disponibilidad del Item para poder ejecutar el
func (i *Item) UpdateAvailable() {
	fmt.Printf("Items %s is available\n", i.name)
	i.available = true
	i.broadCast()
}

// implementamos los metodos de la interfaces para que el struc de Item
func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) broadCast() {

	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

// creamos un notificador de tipo email, este va hacer un tipo de observer.
type EmailClient struct {
	id string
}

func (eC *EmailClient) updateValue(value string) {

	fmt.Printf("Sending Email - %s available from client %s\n", value, eC.id)

}

func (eC *EmailClient) getId() string {
	return eC.id
}

func main() {

	nvidiaItem := NewItem("RTX 3080")

	// vamos a crear dos observadores
	firstObserver := &EmailClient{
		id: "12ab",
	}

	secunObserver := &EmailClient{
		id: "34dc",
	}

	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secunObserver)
	nvidiaItem.UpdateAvailable()
}
