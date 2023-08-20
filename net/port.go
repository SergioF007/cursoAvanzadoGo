package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "urL to scan")

func main() {
	flag.Parse() // tomamos los parametros que se implmentan con flag y los vamos a dejar disponibles dentro de la variable site
	var wg sync.WaitGroup

	for port := 1; port < 100; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conexion, err := net.Dial("tcp", fmt.Sprintf("%s: %d", *site, port))
			if err != nil {
				return
			}
			conexion.Close()
			fmt.Printf("Port %d is open\n", port)

		}(port)
	}
	wg.Wait()
}
