package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func gestionErr(err) {
	if err != nil {
		fmt.Printf("ERREUR") // à preciser
		os.Exit(1)
	}
}

func getArgs() int {
	if len(os.Args) != 2 {
		fmt.Printf("Please use: go run server.go 'PortNumber' \n")
		os.Exit(1)
	} else {
		fmt.Printf(os.Args[1])
		port, err := strconv.Atoi(os.Args[1])
		gestionErr(err)
		return port
	}
}

func main() {
	port := getArgs()
	portString := fmt.Sprintf(":%s", strconv.Itoa(port))
	listener, err := net.Listen("tcp", portString)
	gestionErr(err)

	nbconn := 0
	for {
		fmt.Printf("New connection \n")
		conn, err := listener.Accept()
		gestionErr(err)
		go gestionConnexion(conn, nbconn)
		nbconn += 1
	}

}

func gestionConnexion(conn net.Conn, nbconn int) {
	defer conn.Close()
	connReader := bufio.NewReader(conn)
	for {
		ligne, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Printf("Erreur client", nbconn)
			break
		}
		// go dijkstra(ligne)
	}
}
