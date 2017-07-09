package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/scgolang/sc"
)

// QueryBuf queries for information about a buffer.
type QueryBuf struct {
	flagErrorHandling flag.ErrorHandling
	scc               *sc.Client
}

// Run runs the command.
func (qb QueryBuf) Run(args []string) error {
	if numargs := len(args); numargs == 0 {
		return ErrUsage
	}
	num, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return ErrUsage
	}
	buf, err := qb.scc.QueryBuffer(int32(num))
	if err != nil {
		return err
	}
	printBuf(buf)
	return nil
}

// Usage prints out a usage message.
func (qb QueryBuf) Usage() {
	fmt.Fprint(os.Stderr, `
scc [GLOBAL OPTIONS] querybuf NUM
`)
}

func printBuf(buf *sc.Buffer) {
	fmt.Printf("%-20s%-20s%-20s%-20s\n", "BUFFER", "CHANNELS", "FRAMES", "SAMPLE RATE")
	fmt.Printf("%-20d%-20d%-20d%-20f\n", buf.Num, buf.Channels, buf.Frames, buf.SampleRate)
}
