package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var sortFlag string
var quietFlag bool

var rootCmd = &cobra.Command{
	Use:   "kt",
	Short: "Wrapper script for kubectl top",
	Long: `Extends the functionality of kubectl top.
Additional features:
	- top all of the pods on a node
	- top a pod regardless of current namespace
	- top all pods across all namespaces
`,
}

// Execute : runs errything
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkSortFlag(cmd *cobra.Command) (err error) {
	if !cmd.Flags().Changed("sort") {
		return nil
	}

	switch sortFlag {
	case "name":
		return nil
	case "memory":
		return nil
	case "cpu":
		return nil
	case "namespace":
		return nil
	default:
		return fmt.Errorf("Error: sort specified does not match [name|memory|cpu|namespace]")
	}
}
