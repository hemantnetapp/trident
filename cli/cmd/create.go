package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a resource to Trident",
}
