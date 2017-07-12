package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/scgolang/sc"
)

// Nfree frees a node.
type Nfree struct {
	Controls map[string]float32
	NodeID   int

	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
}

// Run runs the command.
func (nfree Nfree) Run(args []string) error {
	fs := flag.NewFlagSet("nfree", flagErrorHandling)
	fs.IntVar(&nfree.NodeID, "id", 0, "node ID")
	if err := fs.Parse(args); err != nil {
		return ErrUsage
	}
	if nfree.NodeID <= 0 {
		return ErrUsage
	}
	return nfree.scc.NodeFree(int32(nfree.NodeID))
}

// Usage prints a usage message.
func (nfree Nfree) Usage() {
	fmt.Fprint(os.Stderr, `
scc [GLOBAL OPTIONS] nfree [OPTIONS]

OPTIONS
  -id    (REQUIRED) Node ID.
`)
}
