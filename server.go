package main

import (
	"flag"

	"github.com/pkg/errors"
	"github.com/scgolang/sc"
)

// Server represents a running instance of scsynth.
type Server struct {
	Commands map[string]Command

	cmd               Command
	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
	server            *sc.Server
}

// Initialize initializes an instance of scsynth.
func (s *Server) Initialize(args []string) error {
	if len(args) < 1 {
		return errors.New("missing server command")
	}
	s.Commands = map[string]Command{
		"start": &serverStart{flagErrorHandling: s.flagErrorHandling, scc: s.scc},
		"stop":  &serverStop{flagErrorHandling: s.flagErrorHandling, scc: s.scc},
	}
	cmd, ok := s.Commands[args[0]]
	if !ok {
		return errors.Errorf("unknown command %s", args[0])
	}
	s.cmd = cmd

	return nil
}

// Run runs the command.
func (s *Server) Run(args []string) error {
	if err := s.Initialize(args); err != nil {
		return err
	}
	return nil
}

// Usage prints out a usage message.
func (s *Server) Usage() {
}

type serverStart struct {
	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
}

func (s *serverStart) Initialize(args []string) error {
	return nil
}

func (s *serverStart) Run(args []string) error {
	if err := s.Initialize(args); err != nil {
		return err
	}
	return nil
}

func (s *serverStart) Usage() {
}

type serverStop struct {
	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
}

func (s *serverStop) Initialize(args []string) error {
	return nil
}

func (s *serverStop) Run(args []string) error {
	if err := s.Initialize(args); err != nil {
		return err
	}
	return nil
}

func (s *serverStop) Usage() {
}
