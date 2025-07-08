package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var grammarCmd = &cobra.Command{
	Use:   "grammar",
	Short: "Used to download grammar list cards for the specified JLPT level",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grammar called")
	},
}

func init() {
	rootCmd.AddCommand(grammarCmd)
}
