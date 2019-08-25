package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/kylelemons/godebug/pretty"
	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "diff compares 2 yaml files",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("requires two files to compare")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := Diff(args); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)
}

// Diff the two files
func Diff(args []string) error {
	var af aurora.Aurora

	//TODO add -n|--no-color
	noColor := false
	if noColor == false && !isTerminal() {
		noColor = true
	}
	af = aurora.NewAurora(true)

	yaml1 := unmarshal(af, args[0])
	yaml2 := unmarshal(af, args[1])

	diff := computeDiff(af, yaml1, yaml2)
	if diff != "" {
		fmt.Println(args[0], "(-)", args[1], "(+)")
		fmt.Println(diff)
		os.Exit(8)
	}

	return nil
}

func unmarshal(af aurora.Aurora, filename string) interface{} {
	var contents []byte
	var err error

	if filename == "-" {
		contents, err = ioutil.ReadAll(os.Stdin)
	} else {
		contents, err = ioutil.ReadFile(filename)
	}
	if err != nil {
		errorFail(af, err)
	}

	var lines interface{}
	err = yaml.Unmarshal(contents, &lines)
	if err != nil {
		errorFail(af, err)
	}
	return lines
}

func errorFail(af aurora.Aurora, errors ...error) {
	if len(errors) == 0 {
		return
	}
	var errorMessages []string
	for _, err := range errors {
		errorMessages = append(errorMessages, err.Error())
	}
	fmt.Fprintf(os.Stderr, "%v\n\n", af.Red(strings.Join(errorMessages, "\n")))
	os.Exit(1)
}

func computeDiff(af aurora.Aurora, a interface{}, b interface{}) string {
	diffs := make([]string, 0)
	for _, s := range strings.Split(pretty.Compare(a, b), "\n") {
		switch {
		case strings.HasPrefix(s, "+"):
			diffs = append(diffs, af.Bold(af.Green(s)).String())
		case strings.HasPrefix(s, "-"):
			diffs = append(diffs, af.Bold(af.Red(s)).String())
		}
	}
	return strings.Join(diffs, "\n")
}

func isTerminal() bool {
	fd := os.Stdout.Fd()

	if isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd) {
		return true
	}

	return false
}
