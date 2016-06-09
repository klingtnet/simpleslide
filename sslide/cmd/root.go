package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/klingtnet/simpleslide/lib"
)

// RootCmd is the base command
//
// TODO
// - add ids to each slide (allow hyperlinking)
// - add page number
// - show progressbar
// - add print media query to enable PDF export
var RootCmd = &cobra.Command{
	Use:   utils.NAME,
	Short: "A super simple slideshow generator",
	Long:  utils.NAME + ` is a super simple slideshow generator that takes markdown as input and outputs beautiful, self-contained HTML5  slides.`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.NAME, utils.Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(buildCmd)
	RootCmd.AddCommand(serveCmd)
}
