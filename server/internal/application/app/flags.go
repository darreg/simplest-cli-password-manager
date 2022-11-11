package app

import "flag"

const NotAvailable string = "N/A"

// Flags stores information about command-line flags.
type Flags struct {
	Debug bool
	A     string
	D     string
	Crt   string
	Key   string
}

func NewFlags() *Flags {
	flags := &Flags{}

	flag.BoolVar(&flags.Debug, "debug", false, "Enable debug")
	flag.StringVar(&flags.A, "a", NotAvailable, "Run address")
	flag.StringVar(&flags.D, "d", NotAvailable, "Database uri")
	flag.StringVar(&flags.Crt, "crt", NotAvailable, "Certificate file")
	flag.StringVar(&flags.Key, "key", NotAvailable, "Key file")

	flag.Parse()

	return flags
}
