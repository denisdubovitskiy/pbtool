package main

import (
	"log"
	"os"

	"github.com/denisdubovitskiy/pbtool/internal/application"
)

func main() {
	app, err := application.New()
	if err != nil {
		log.Fatalf("unable to create application: %v\n", err)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
