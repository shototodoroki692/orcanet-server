package main

import (
	"fmt"
	"orcanet-server/api"
)

func main() {
	fmt.Printf("exécution du serveur Orca Network\n")

	apiServer := api.NewAPIServer(":3001")

	apiServer.Run()
}