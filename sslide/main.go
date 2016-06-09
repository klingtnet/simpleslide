package main

import (
	"github.com/klingtnet/simpleslide/lib"
	"github.com/klingtnet/simpleslide/sslide/cmd"
)

func main() {
	utils.ExitIfErr(cmd.RootCmd.Execute())
}
