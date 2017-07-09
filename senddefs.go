package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/scgolang/sc"
)

// SendDefs sends some synthdefs.
type SendDefs struct {
	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
}

// Run runs the command.
func (sd SendDefs) Run(args []string) error {
	return sd.scc.SendAllDefs()
}

// Usage prints a usage message.
func (sd SendDefs) Usage() {
	fmt.Fprint(os.Stderr, `
scc [GLOBAL OPTIONS] senddefs
`)
}
