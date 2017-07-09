package main

import "errors"

// Help prints out help messages for other commands.
type Help struct {
	Commands map[string]Command
}

// Run runs the command.
func (h Help) Run(args []string) error {
	if len(args) == 0 {
		usage()
		return nil
	}
	cmd, ok := h.Commands[args[0]]
	if !ok {
		return errors.New("unrecognized command: " + args[0])
	}
	cmd.Usage()
	return nil
}

// Usage prints a usage message for the command.
func (h Help) Usage() {
	usage()
}
