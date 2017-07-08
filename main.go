package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Usage = usage

	app := &App{
		Config: ParseConfig(),
	}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `
scc [GLOBAL OPTIONS] command [COMMAND OPTIONS]

GLOBAL OPTIONS
  -scsynth                     Remote address of scsynth.

COMMANDS
  gquery                       Query the SuperCollider node graph.
  help                         Print this help message.

For help with a particular command, "scc help COMMAND"
`)
}
