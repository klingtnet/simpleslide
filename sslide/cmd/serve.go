package cmd

import (
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the slideshow",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
