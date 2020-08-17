package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the root command
var RootCmd = &cobra.Command{
	Use:   "notes",
	Short: "ClI based Sticky notes application",
}
