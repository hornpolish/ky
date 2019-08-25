package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Hugo",
  Long:  `All software has versions. KY's is slippier`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("KY v0.9 -- HEAD")
  },
}
