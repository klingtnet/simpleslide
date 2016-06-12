package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-errors/errors"
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
		PrintErr(errors.Errorf("error: expected %d arguments, got %d\n", argCnt, len(args)))
		cmd.Usage()
		os.Exit(1)
	}
}

func PrintIfErr(err error) {
	if err != nil {
		PrintErr(err)
	}
}

func PrintErr(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", errors.WrapPrefix(err, "error", 0))
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
