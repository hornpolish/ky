package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "split splits big-ole.yaml into name*-type*.yaml or name/type.yaml",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires one file to split")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := Split(args); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)
}

// Split takes
func Split(args []string) error {
	log.Println("Hello from Split", args[0])
	return nil
}
