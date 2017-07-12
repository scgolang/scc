package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/scgolang/sc"
)

// Common errors.
var (
	ErrUsage = errors.New("usage error")
)

// App defines the application's behavior.
type App struct {
	Config

	Commands map[string]Command
}

// Initialize initializes the application.
func (app *App) Initialize() error {
	scc, err := sc.NewClient("udp", "127.0.0.1:0", app.ScsynthAddr, 5*time.Second)
	if err != nil {
		return errors.Wrap(err, "creating client")
	}
	if _, err := scc.AddDefaultGroup(); err != nil {
		return err
	}
	app.Commands = map[string]Command{
		"gnew":     &GNew{flagErrorHandling: flagErrorHandling, scc: scc},
		"gquery":   &GQuery{flagErrorHandling: flagErrorHandling, scc: scc},
		"nfree":    &Nfree{flagErrorHandling: flagErrorHandling, scc: scc},
		"nmap":     &Nmap{flagErrorHandling: flagErrorHandling, scc: scc},
		"nmapa":    &Nmapa{flagErrorHandling: flagErrorHandling, scc: scc},
		"nset":     &Nset{flagErrorHandling: flagErrorHandling, scc: scc},
		"querybuf": &QueryBuf{flagErrorHandling: flagErrorHandling, scc: scc},
		"readbuf":  &ReadBuf{flagErrorHandling: flagErrorHandling, scc: scc},
		"senddefs": &SendDefs{flagErrorHandling: flagErrorHandling, scc: scc},
		"synth":    &Synth{flagErrorHandling: flagErrorHandling, scc: scc},
	}
	app.Commands["help"] = Help{Commands: app.Commands}
	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, `
scc [GLOBAL OPTIONS] command [COMMAND OPTIONS]

GLOBAL OPTIONS
  -scsynth                     Remote address of scsynth.

COMMANDS
  gnew                         Create a new group.
  gquery                       Query the SuperCollider node graph.
  help                         Print this help message.
  nfree                        Free a node.
  nmap                         Map control bus(es) to a node's control value(s).
  nmapa                        Map audio bus(es) to a node's control value(s).
  nset                         Set a node's control value(s).
  querybuf                     Query for information about a buffer.
  readbuf                      Read a buffer.
  senddefs                     Send all the synthdefs that are registered in the sc pkg.
  synth                        Create a synth node.

For help with a particular command, "scc help COMMAND"
`)
}

// Run runs the application.
func (app *App) Run() error {
	if err := app.Initialize(); err != nil {
		return err
	}
	args := flag.Args()

	if len(args) == 0 {
		return errors.New("missing command")
	}
	cmd, ok := app.Commands[args[0]]
	if !ok {
		return errors.New("unrecognized command: " + args[0])
	}
	err := cmd.Run(args[1:])
	if err == ErrUsage {
		cmd.Usage()
		return nil
	}
	return err
}

// Command is a command.
type Command interface {
	Run(args []string) error
	Usage()
}
