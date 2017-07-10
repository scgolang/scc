package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/scgolang/sc"
)

// Nset sets a node's control value.
type Nset struct {
	Controls map[string]float32
	NodeID   int

	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
}

// Run runs the command.
func (nset Nset) Run(args []string) error {
	fs := flag.NewFlagSet("nset", flagErrorHandling)
	fs.IntVar(&nset.NodeID, "id", 0, "node ID")
	if err := fs.Parse(args); err != nil {
		return ErrUsage
	}
	if nset.NodeID <= 0 {
		return ErrUsage
	}
	if len(fs.Args()) == 0 {
		return ErrUsage // Print a usage message if there are no kv pairs.
	}
	ctls, ok := getSynthControls(fs.Args())
	if !ok {
		return ErrUsage
	}
	return nset.scc.NodeSet(int32(nset.NodeID), ctls)
}

// Usage prints a usage message.
func (nset Nset) Usage() {
	fmt.Fprint(os.Stderr, `
scc [GLOBAL OPTIONS] nset [OPTIONS] CTL=VALUE [... CTL=VALUE]

OPTIONS
  -id    (REQUIRED) Node ID.

CTL must be a string and VALUE must be a float.
`)
}
