package utils

import (
	"fmt"
	"os"
)

// NAME contains the name of the program.
const NAME = "simpleslide"

// Version contains the programs semver string.
var Version = "unknown"

func ExitOnErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}

func PrintIfErr(err error) {
	if err != nil {
		PrintErr("error: %\n", err)
	}
}

func PrintErr(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}
