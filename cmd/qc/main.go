package main

import (
	"github.com/kunit/qc"
	"github.com/kunit/qc/version"
	"os"
)

func main() {
	os.Exit(qc.RunCLI(qc.Env{
		Out:     os.Stdout,
		Err:     os.Stderr,
		Args:    os.Args[1:],
		Version: version.Version,
	}))
}
