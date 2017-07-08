package main

import (
	"flag"
)

type Config struct {
	ScsynthAddr string
}

func ParseConfig() Config {
	var c Config
	flag.StringVar(&c.ScsynthAddr, "scsynth", "127.0.0.1:57120", "scsynth listening address")
	flag.Parse()
	return c
}
