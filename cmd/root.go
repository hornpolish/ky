package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ky",
	Short: "ky - kubernetes yaml swiss army knife",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	//cwd, _ := os.Getwd()
	//log.Println("Hello from init cwd is", cwd)

	return
}

func initConfig() {
	//log.Println("Hello from initConfig")
}
