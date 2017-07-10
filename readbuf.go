package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"
	"github.com/scgolang/sc"
)

// ReadBuf reads a buffer.
type ReadBuf struct {
	Channels IntFlags
	Num      int

	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
}

// Run runs the command.
func (rb ReadBuf) Run(args []string) error {
	fs := flag.NewFlagSet("readbuf", rb.flagErrorHandling)
	fs.Var(&rb.Channels, "channel", "channel(s) to read")
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
	_, err = rb.scc.ReadBuffer(p, int32(rb.Num), rb.Channels...)
	return errors.Wrap(err, "reading buffer")
}

// Usage prints a usage message.
func (rb ReadBuf) Usage() {
	fmt.Fprint(os.Stderr, `
scc [GLOBAL OPTIONS] readbuf [OPTIONS] FILE

OPTIONS
  -channel    (OPTIONAL) Channel(s) to read. Can be passed multiple times. Indices start at 0.
  -num        (REQUIRED) Buffer number.

FILE will be converted to an absolute path.
`)
}

// IntFlags provides a way to pass multiple int flags.
type IntFlags []int

// String converts intflags to a string.
func (ifs IntFlags) String() string {
	return fmt.Sprintf("%#v", ifs)
}

// Set sets intflags from the provided string.
func (ifs *IntFlags) Set(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*ifs = append(*ifs, int(i))
	return nil
}
