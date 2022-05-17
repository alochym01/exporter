package main

import (
	"flag"

	"github.com/alochym01/exporter/v1/router"
)

func main() {
	// Command-line Flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	env := flag.String("env", "debug", "HTTP network address")
	timeout := flag.String("timeout", "5", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line
	// This reads in the command-line flag value and assigns it to the addr
	flag.Parse()

	r := router.Router(*env, *timeout)

	r.Run(*addr)
}
