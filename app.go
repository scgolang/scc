package main

import (
	"flag"
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
	app.Commands = map[string]Command{
		"gquery": GQuery{scc: scc},
	}
	return nil
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
