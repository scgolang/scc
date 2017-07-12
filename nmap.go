package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/scgolang/sc"
)

// Nmap sets a node's control value.
type Nmap struct {
	Controls map[string]float32
	NodeID   int

	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
}

// Run runs the command.
func (nmap Nmap) Run(args []string) error {
	fs := flag.NewFlagSet("nmap", flagErrorHandling)
	fs.IntVar(&nmap.NodeID, "id", 0, "node ID")
	if err := fs.Parse(args); err != nil {
		return ErrUsage
	}
	if nmap.NodeID <= 0 {
		return ErrUsage
	}
	if len(fs.Args()) == 0 {
		return ErrUsage // Print a usage message if there are no kv pairs.
	}
	busMapping, ok := getBusMapping(fs.Args())
	if !ok {
		return ErrUsage
	}
	return nmap.scc.NodeMap(int32(nmap.NodeID), busMapping)
}

// Usage prints a usage message.
func (nmap Nmap) Usage() {
	fmt.Fprint(os.Stderr, `
scc [GLOBAL OPTIONS] nmap [OPTIONS] CTL=VALUE [... CTL=VALUE]

OPTIONS
  -id    (REQUIRED) Node ID.

CTL must be a string and VALUE must be a control bus index (integer).
`)
}

func getBusMapping(args []string) (map[string]int32, bool) {
	if len(args) == 0 {
		return nil, true
	}
	m := map[string]int32{}

	for _, arg := range args {
		pieces := strings.Split(arg, "=")

		if len(pieces) != 2 {
			return nil, false
		}
		val, err := strconv.ParseInt(pieces[1], 10, 32)
		if err != nil {
			return nil, false
		}
		m[pieces[0]] = int32(val)
	}
	return m, true
}
