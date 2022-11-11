package app

import "flag"

const NotAvailable string = "N/A"

// Flags stores information about command-line flags.
type Flags struct {
	Debug bool
	A     string
}

func NewFlags() *Flags {
	flags := &Flags{}

	flag.BoolVar(&flags.Debug, "debug", false, "Enable debug")
	flag.StringVar(&flags.A, "a", NotAvailable, "Server address")

	flag.Parse()

	return flags
}
