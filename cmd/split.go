package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// splitVar is the global struct for the split command
var splitVar struct {
	outdir string
	nname  bool
	ntype  bool
	nlabel string
	debug  bool
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
	splitCmd.Flags().BoolVarP(&splitVar.nname, "nest-name", "n", false, "Nested Directory per name?")
	splitCmd.Flags().BoolVarP(&splitVar.ntype, "nest-type", "t", false, "Nested Directory per type?")
	splitCmd.Flags().StringVarP(&splitVar.nlabel, "nest-label", "l", "sas.com/admin", "Nested Directory per label value")
	splitCmd.Flags().BoolVarP(&splitVar.debug, "debug", "d", false, "Debug?")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Split takes
func Split(args []string) error {

	files := readAndSplitFile(args[0])

	// TODO - error if directory exists, not empty
	os.Mkdir(splitVar.outdir, os.ModePerm)

	for _, value := range files {

		var filename string
		var kind string
		var name string
		var filter string

		result := make(map[interface{}]interface{})

		err := yaml.Unmarshal([]byte(value), &result)
		if err != nil {
			panic(err)
		}

		kind, _ = result["kind"].(string)

		if _, ok := result["metadata"]; ok {
			metadata := result["metadata"].(map[interface{}]interface{})
			name, _ = metadata["name"].(string)
			if _, ok := metadata["labels"]; ok {
				labels := metadata["labels"].(map[interface{}]interface{})
				filter, _ = labels[splitVar.nlabel].(string)
			}
		}

		if err != nil || kind == "" {
			fmt.Println("yaml file with no kind")
			fmt.Println(value)
			continue
		}

		if splitVar.nname == true {
			dirname := path.Join(splitVar.outdir, name)
			os.Mkdir(dirname, os.ModePerm)
			filename = path.Join(dirname, kind+".yaml")

		} else if splitVar.ntype == true {
			dirname := path.Join(splitVar.outdir, kind)
			os.Mkdir(dirname, os.ModePerm)
			filename = path.Join(dirname, name+".yaml")

		} else if filter != "" {
			dirname := path.Join(splitVar.outdir, filter)
			os.Mkdir(dirname, os.ModePerm)
			filename = path.Join(dirname, name+"-"+kind+".yaml")

		} else {
			filename = path.Join(splitVar.outdir, name+"-"+kind+".yaml")
		}

		fmt.Println(filename)

		f, err := os.Create(filename)
		if err != nil {
			// print something!
			continue
		}

		_, err = f.WriteString("---\n" + value)
		if err != nil {
			// print something!
		}

		f.Close()
	}

	return nil
}

func readAndSplitFile(logfile string) map[int]string {

	if splitVar.debug {
		fmt.Printf("debug: split %s\n", logfile)
	}

	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil
	}

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

		if strings.HasPrefix(line, "---") {
			count++
		} else {
			c[count] += line
		}
	}

	f.Close()
	return c
}
