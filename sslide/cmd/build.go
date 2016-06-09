package cmd

import (
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds the slideshow",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
