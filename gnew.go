package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/scgolang/sc"
)

// GNew creates a group.
type GNew struct {
	Action string
	ID     int
	Target int

	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
}

// Run runs the command.
func (gn GNew) Run(args []string) error {
	fs := flag.NewFlagSet("gnew", gn.flagErrorHandling)
	fs.StringVar(&gn.Action, "action", "AddToTail", "add action")
	fs.IntVar(&gn.ID, "id", 1, "group ID")
	fs.IntVar(&gn.Target, "target", 0, "target node ID for the add action")
	if err := fs.Parse(args); err != nil {
		return ErrUsage
	}
	an, err := getActionNumber(gn.Action)
	if err != nil {
		return err
	}
	_, err = gn.scc.Group(int32(gn.ID), an, int32(gn.Target))
	return err
}

// Usage prints a usage message.
func (gn GNew) Usage() {
	fmt.Fprint(os.Stderr, `
scc [GLOBAL OPTIONS] gnew [OPTIONS]

OPTIONS
  -action                         (OPTIONAL) Add action. Default is AddToTail. Possible values are listed here https://git.io/vQ6zA
  -id                             (OPTIONAL) Group ID. Default is 1.
  -target                         (OPTIONAL) Target node for the add action. Default is 0 (root node).
`)
}
