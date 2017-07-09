package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/scgolang/sc"
)

// DefaultSynthID is the default ID for synth nodes.
const DefaultSynthID = 1000

// Synth adds synth nodes to the SuperCollider node graph.
type Synth struct {
	Action   string
	Controls map[string]float32
	Def      string
	Group    int
	ID       int

	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
}

func (s *Synth) Initialize(args []string) error {
	fs := flag.NewFlagSet("synth", s.flagErrorHandling)
	fs.StringVar(&s.Action, "action", "AddToTail", "add action")
	fs.StringVar(&s.Def, "def", "", "synthdef name")
	fs.IntVar(&s.Group, "group", 0, "group ID")
	fs.IntVar(&s.ID, "id", 0, "synth node ID")
	if err := fs.Parse(args); err != nil {
		return ErrUsage
	}
	if len(s.Def) == 0 {
		return ErrUsage
	}
	ctls, ok := getSynthControls(fs.Args())
	if !ok {
		return ErrUsage
	}
	s.Controls = ctls
	return nil
}

// Run runs the command.
func (s *Synth) Run(args []string) error {
	if err := s.Initialize(args); err != nil {
		return err
	}
	an, err := getActionNumber(s.Action)
	if err != nil {
		return err
	}
	_, err = s.scc.Synth(s.Def, int32(s.ID), an, int32(s.Group), s.Controls)
	return err
}

// Usage prints out a usage message.
func (s *Synth) Usage() {
	fmt.Fprint(os.Stderr, `
scc [GLOBAL OPTIONS] synth [OPTIONS] [PARAMETERS]

OPTIONS
  -action                         (OPTIONAL) Add action. Default is AddToTail. Possible values are listed here https://git.io/vQ6zA
  -def                            (REQUIRED) Synthdef name.
  -group                          (OPTIONAL) Group ID. Default is 0 (root group).
  -id                             (OPTIONAL) Synth node ID. Default is 1000.

PARAMETERS
  The synth parameters are passed as key-value pairs separated by "=".
  For example, freq=440 gain=0.5
  The key must be a string and the value will be parse with this https://golang.org/pkg/strconv/#ParseFloat.
`)
}

var actions = map[string]int32{
	"AddToHead":  sc.AddToHead,
	"AddToTail":  sc.AddToTail,
	"AddBefore":  sc.AddBefore,
	"AddAfter":   sc.AddAfter,
	"AddReplace": sc.AddReplace,
}

func getActionNumber(s string) (int32, error) {
	a, ok := actions[s]
	if !ok {
		return 0, ErrUsage
	}
	return a, nil
}

func getSynthControls(args []string) (map[string]float32, bool) {
	if len(args) == 0 {
		return nil, true
	}
	m := map[string]float32{}

	for _, arg := range args {
		pieces := strings.Split(arg, "=")

		if len(pieces) != 2 {
			return nil, false
		}
		val, err := strconv.ParseFloat(pieces[1], 64)
		if err != nil {
			return nil, false
		}
		m[pieces[0]] = float32(val)
	}
	return m, true
}
