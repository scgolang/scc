package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/scgolang/sc"
)

// Nfree frees a node.
type Nfree struct {
	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
}

// Run runs the command.
func (nfree Nfree) Run(args []string) error {
	if len(args) == 0 {
		return ErrUsage
	}
	nodeID, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return ErrUsage
	}
	if nodeID <= 0 {
		return ErrUsage
	}
	return nfree.scc.NodeFree(int32(nodeID))
}

// Usage prints a usage message.
func (nfree Nfree) Usage() {
	fmt.Fprint(os.Stderr, `
scc [GLOBAL OPTIONS] nfree NODE_ID
`)
}
