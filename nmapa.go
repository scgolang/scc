package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/scgolang/sc"
)

// Nmapa sets a node's control value.
type Nmapa struct {
	Controls map[string]float32
	NodeID   int

	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
}

// Run runs the command.
func (nmapa Nmapa) Run(args []string) error {
	fs := flag.NewFlagSet("nmapa", flagErrorHandling)
	fs.IntVar(&nmapa.NodeID, "id", 0, "node ID")
	if err := fs.Parse(args); err != nil {
		return ErrUsage
	}
	if nmapa.NodeID <= 0 {
		return ErrUsage
	}
	if len(fs.Args()) == 0 {
		return ErrUsage // Print a usage message if there are no kv pairs.
	}
	busMapping, ok := getBusMapping(fs.Args())
	if !ok {
		return ErrUsage
	}
	return nmapa.scc.NodeMapa(int32(nmapa.NodeID), busMapping)
}

// Usage prints a usage message.
func (nmapa Nmapa) Usage() {
	fmt.Fprint(os.Stderr, `
scc [GLOBAL OPTIONS] nmapa [OPTIONS] CTL=VALUE [... CTL=VALUE]

OPTIONS
  -id    (REQUIRED) Node ID.

CTL must be a string and VALUE must be an audio bus index (integer).
`)
}
