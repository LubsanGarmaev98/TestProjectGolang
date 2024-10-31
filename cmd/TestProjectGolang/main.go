package main

import (
	app2 "github.com/LubsanGarmaev98/TestProjectGolang/internal/pkg/app"
	"log"
)

func main() {
	app, err := app2.New()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
