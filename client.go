package main

import (
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
	//On vérifie qu'on a bien un argument qui doit être le numéro de port
	if len(os.Args) != 2 {
		fmt.Printf("Please use: go run client.go 'PortNumber' \n")
		os.Exit(1)
	} else {
		fmt.Printf(os.Args[1])
		port, err := strconv.Atoi(os.Args[1])
		gestionErr(err)
		return port
	}
}

func main() {
	// Recupération du numéro de port et création port string
	port := getArgs()
	portStr := fmt.Sprintf("127.0.0.1:%s", strconv.Itoa(port))
	// Connexion
	conn, err := net.Listen("tcp", portStr)
	gestionErr(err)

	defer conn.Close()
	// reader + envoie de la matrice au serveur ...
}
