package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// NAME contains the name of the program.
const NAME = "simpleslide"

// Version contains the programs semver string.
var Version = "unknown"

func ExitIfErr(err error) {
	PrintIfErr(err)
	if err != nil {
		os.Exit(1)
	}
}

func ExpectExactArgs(cmd *cobra.Command, argCnt int, args []string) {
	if len(args) != argCnt {
		PrintErr("error: expected %d arguments, got %d\n", argCnt, len(args))
		cmd.Usage()
		os.Exit(1)
	}
}

func PrintIfErr(err error) {
	if err != nil {
		PrintErr("error: %s\n", err)
	}
}

func PrintErr(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

func PrintIfVerbose(verbose bool, format string, args ...interface{}) {
	if verbose {
		fmt.Printf(format, args...)
	}
}

func ReadFileAsString(filepath string) (string, error) {
	raw, err := ioutil.ReadFile(filepath)
	if raw != nil {
		return string(raw), err
	}
	return "", err
}
