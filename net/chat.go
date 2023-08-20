package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	inconmigClients = make(chan Client) // esto es un canal de canles para los clientes que se estan conectando
	leavingClients  = make(chan Client) // esta canal de canales es para los clientes que se estan deconectando
	messages        = make(chan string) // los mensajes que va ha estar siendo transmitidos
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port")
)

// Client1 -> Server -> HandleConnection(Client1)

func HandleConnection(conexion net.Conn) {
	defer conexion.Close()
	message := make(chan string)

	go MessageWrite(conexion, message)

	// vamos a darle un nombre a cada cliente que se connecta
	//lo siguiente es la representacion del puerto que se esta conectando a esta conexion
	// Ej: Client1:2560   ó  si e estuviese conectamdo a Platzi.com, por el puerto 38
	// platzi.com:38
	clientName := conexion.RemoteAddr().String()

	// mandamos un mensaje al cliente que se acaba de conectar
	message <- fmt.Sprintf("Welcome to the server, your name %s\n", clientName)
	// utilizo el canal de mensajes del sistema para que le transmita a los demas cuando se a conectado un cliente nuevo
	messages <- fmt.Sprintf("New client is here, name %s\n", clientName)

	inconmigClients <- message
	// vamos a empezar a leer todo lo que se va a estar escribiendo en esta conexion
	inputMessage := bufio.NewScanner(conexion)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())

	}

	leavingClients <- message
	messages <- fmt.Sprintf("%s saig goodbye", clientName)

}

func MessageWrite(conexion net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprintln(conexion, message)
	}
}

func Broadcast() {
	clients := make(map[Client]bool)
	for {
		select {
		case message := <-messages:
			for client := range clients {
				client <- message
			}
		case newClient := <-inconmigClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}

	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port)) // el net.Listen se utilizar para crear nuestro propio servidor
	if err != nil {
		log.Fatal(err)
	}

	go Broadcast()

	for {
		conexion, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConnection(conexion)
	}

}
