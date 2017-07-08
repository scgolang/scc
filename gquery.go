package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/scgolang/sc"
)

// GQuery queries the SuperCollider node graph.
type GQuery struct {
	GroupID int

	scc *sc.Client
}

// Run runs the command.
func (gq *GQuery) Run(args []string) error {
	fs := flag.NewFlagSet("gquery", flagErrorHandling)
	fs.IntVar(&gq.GroupID, "group", 0, "Group ID")
	if err := fs.Parse(args); err != nil {
		return ErrUsage
	}
	g, err := gq.scc.QueryGroup(int32(gq.GroupID))
	if err != nil {
		return errors.Wrap(err, "querying group")
	}
	printGroup(g)
	return nil
}

// Usage prints a usage message.
func (gq *GQuery) Usage() {
	fmt.Fprint(os.Stderr, `
scc gquery [OPTIONS]

OPTIONS
  -group                         (OPTIONAL) Group ID. Default is 0 (root group).
`)
}

func printGroup(g *sc.GroupNode) error {
	err := printGroupP(g, "")
	fmt.Println("")
	return err
}

func printGroupP(g *sc.GroupNode, prefix string) error {
	fmt.Printf("%sgroup(id=%d)", prefix, g.ID())

	for i, child := range g.Children {
		if i == len(g.Children)-1 {
			fmt.Printf(prefix + "\u2514\u2500\u2500 ")
		} else {
			fmt.Printf(prefix + "\u251c\u2500\u2500 ")
		}
		var childPrefix string

		if i == len(g.Children)-1 {
			childPrefix = prefix + "    "
		} else {
			childPrefix = prefix + "\u2502   "
		}
		switch c := child.(type) {
		case *sc.SynthNode:
			printSynthP(c, childPrefix)
		case *sc.GroupNode:
			printGroupP(c, childPrefix)
		default:
			return errors.Errorf("unrecognized node type: %T", c)
		}
	}
	return nil
}

func printSynthP(s *sc.SynthNode, prefix string) {
}
