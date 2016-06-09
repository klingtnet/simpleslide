package utils

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// NAME contains the name of the program.
const NAME = "simpleslide"

// Version contains the programs semver string.
var Version = "unknown"

func ExitOnErr(err error) {
	PrintIfErr(err)
	if err != nil {
		os.Exit(1)
	}
}

func ExpectExactArgs(cmd *cobra.Command, argCnt int) {
	if len(cmd.ValidArgs) != argCnt {
		PrintErr("error: expected %d arguments, got %d\n", argCnt, len(cmd.ValidArgs))
		cmd.Usage()
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
