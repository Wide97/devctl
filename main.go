package main

import (
	"devctl/cmd"
	"fmt"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Errore")
		os.Exit(1)
	}
}

//aggiungo os exit e gestione errore per il momento
