package cmd

import (
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "builds",
	Short: "Builds the slideshow",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
