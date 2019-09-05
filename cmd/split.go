package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// splitVar is the global struct for the split command
var splitVar struct {
	outdir string
	nest   bool
}

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
	splitCmd.Flags().StringVarP(&splitVar.outdir, "outdir", "o", "./split", "Directory to split into")
	splitCmd.Flags().BoolVarP(&splitVar.nest, "nest", "n", false, "Nested Directory per name?")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// KubernetesAPI is for unmarshaling objects
type KubernetesAPI struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name   string `yaml:"name"`
		Labels struct {
			Source string `yaml:"source"`
		} `yaml:"labels"`
	} `yaml:"metadata"`
}

// Split takes
func Split(args []string) error {

	files := readAndSplitFile(args[0])

	// TODO - error if directory exists, not empty
	os.Mkdir(splitVar.outdir, os.ModePerm)

	for _, value := range files {

		var m KubernetesAPI
		var filename string

		err := yaml.Unmarshal([]byte(value), &m)
		check(err)

		if m.Kind == "" {
			fmt.Println("yaml file with no kind")
			return nil
		}

		if splitVar.nest == true {
			dirname := splitVar.outdir + "/" + m.Metadata.Name
			os.Mkdir(dirname, os.ModePerm)

			filename = dirname + "/" + m.Kind + ".yaml"
		} else {
			filename = splitVar.outdir + "/" + m.Metadata.Name + "-" + m.Kind + ".yaml"
		}

		// TODO -v :: fmt.Println(filename)

		f, err := os.Create(filename)
		check(err)
		defer f.Close()

		_, err = f.WriteString("---\n" + value)
		check(err)
	}

	return nil
}

func readAndSplitFile(logfile string) map[int]string {

	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
	check(err)

	defer f.Close()

	rd := bufio.NewReader(f)
	c := make(map[int]string)
	count := 0

	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panicf("read file line error: %v", err)
		}

		if strings.TrimSpace(line) == "---" {
			count++
		} else {
			c[count] += line
		}
	}
	return c
}
