package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version string
	date    string
	commit  string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of KY",
	Long:  `All software has versions. KY's is slippier`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("KY version %s built %s from githash %s\n", version, date, commit)
	},
}
