package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// se conecta a  -> host:port
// Escribir al -> host:port
// Leer -> host:port
//ej: > [Hola] -> host:port -> [Hola]

var (
	port = flag.Int("p", 3090, "port")
	host = flag.String("h", "localhost", "host")
)

func main() {
	flag.Parse()
	conexion, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port)) //net.Dial se utiliza es para conectarce al servidor
	if err != nil {
		log.Fatal(err)
	}
	// lo definimos asi porque no nos interesa lo que nos va a estar tranportando
	// ya que solo sirve como canal de control
	done := make(chan struct{})
	go func() { // va a leer todo lo que se esta recibiendo atravez de la conexion  y lo va a escribir en la consola
		io.Copy(os.Stdout, conexion)
		done <- struct{}{}
	}()

	CopyContent(conexion, os.Stdin)
	conexion.Close()
	<-done
}

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
