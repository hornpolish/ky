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

// splitVar is the global struct for the split command
var diffVar struct {
	color   bool
	nocolor bool
}

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
	diffCmd.Flags().BoolVarP(&diffVar.color, "color", "c", isTerminal(), "Colors Please")
	diffCmd.Flags().BoolVarP(&diffVar.nocolor, "nocolor", "n", false, "No Colors Please")
}

// Diff the two files
func Diff(args []string) error {
	diff, err := diff2(args[0], args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n\n", err.Error())
		return err
	}

	if diff != "" {
		fmt.Println(diff)
	}

	return nil
}

func diff2(f1 string, f2 string) (string, error) {
	yaml1, err := unmarshal(f1)
	if err != nil {
		return "", err
	}
	yaml2, err := unmarshal(f2)
	if err != nil {
		return "", err
	}

	return computeDiff(f1, f2, yaml1, yaml2), err
}

func unmarshal(filename string) (interface{}, error) {
	var contents []byte
	var err error

	if filename == "-" {
		contents, err = ioutil.ReadAll(os.Stdin)
	} else {
		contents, err = ioutil.ReadFile(filename)
	}
	if err != nil {
		return nil, err
	}

	var lines interface{}
	err = yaml.Unmarshal(contents, &lines)
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func computeDiff(f1 string, f2 string, a interface{}, b interface{}) string {
	color := diffVar.color
	if diffVar.nocolor == true {
		color = false
	}
	af := aurora.NewAurora(color)

	headr := false
	diffs := make([]string, 0)

	for _, s := range strings.Split(pretty.Compare(a, b), "\n") {

		switch {
		case strings.HasPrefix(s, "+"):
			diffs = append(diffs, af.Bold(af.Green(s)).String())
			headr = true
		case strings.HasPrefix(s, "-"):
			diffs = append(diffs, af.Bold(af.Red(s)).String())
			headr = true
		}
	}

	if headr == true {
		s1 := af.Bold(af.Green(f1 + "(+) ")).String()
		s2 := af.Bold(af.Red(f2 + "(-)")).String()

		return s1 + s2 + "\n" + strings.Join(diffs, "\n")
	}

	return ""
}

func isTerminal() bool {
	fd := os.Stdout.Fd()

	if isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd) {
		return true
	}

	return false
}
