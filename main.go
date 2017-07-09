package main

import (
	"flag"
	"log"
)

var flagErrorHandling flag.ErrorHandling = flag.ContinueOnError

func main() {
	flag.Usage = usage

	app := &App{
		Config: ParseConfig(),
	}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
