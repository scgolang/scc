package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/scgolang/sc"
)

// ReadBuf reads a buffer.
type ReadBuf struct {
	Num int

	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
}

// Run runs the command.
func (rb ReadBuf) Run(args []string) error {
	fs := flag.NewFlagSet("readbuf", rb.flagErrorHandling)
	fs.IntVar(&rb.Num, "num", 0, "buffer number")
	if err := fs.Parse(args); err != nil {
		return ErrUsage
	}
	if len(fs.Args()) == 0 {
		return ErrUsage
	}
	p, err := filepath.Abs(fs.Args()[0])
	if err != nil {
		return errors.Wrap(err, "making path absolute")
	}
	// Check that the file exists and is readable.
	f, err := os.Open(p)
	if err != nil {
		return errors.Wrap(err, "opening file for reading")
	}
	_ = f.Close()

	// Read the buffer.
	_, err = rb.scc.ReadBuffer(p, int32(rb.Num))
	return errors.Wrap(err, "reading buffer")
}

// Usage prints a usage message.
func (rb ReadBuf) Usage() {
	fmt.Fprint(os.Stderr, `
scc [GLOBAL OPTIONS] readbuf [OPTIONS] FILE

OPTIONS
  -num                         (REQUIRED) Buffer number.
`)
}
