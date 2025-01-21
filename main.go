package main

import (
	"command-line-application/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting point.")

	aplicacao := app.Generate()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
