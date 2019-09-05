package cmd

import (
	"testing"

	"github.com/udhos/equalfile"
)

func compareTwo(t *testing.T, file string) {
	cmp := equalfile.New(nil, equalfile.Options{}) // compare using single mode

	equal, err := cmp.CompareFile("../test-output/"+file, "../test-result/"+file)

	if err != nil {
		t.Fail()
	}
	if equal != true {
		t.Errorf("compare %s = %t, want true", file, equal)
	}

}
func TestSplit(t *testing.T) {
	arg := make([]string, 2)

	splitVar.outdir = "../test-output"
	splitVar.nest = false
	arg[0] = "../test/fifo.yaml"
	Split(arg)

	compareTwo(t, "files-Deployment.yaml")
	compareTwo(t, "files-Service.yaml")
	compareTwo(t, "files-VirtualService.yaml")

	compareTwo(t, "folders-Deployment.yaml")
	compareTwo(t, "folders-Service.yaml")
	compareTwo(t, "folders-VirtualService.yaml")
}

func TestSplitNest(t *testing.T) {
	arg := make([]string, 2)

	splitVar.outdir = "../test-output"
	splitVar.nest = true
	arg[0] = "../test/fifo.yaml"

	Split(arg)

	compareTwo(t, "files/Deployment.yaml")
	compareTwo(t, "files/Service.yaml")
	compareTwo(t, "files/VirtualService.yaml")

	compareTwo(t, "folders/Deployment.yaml")
	compareTwo(t, "folders/Service.yaml")
	compareTwo(t, "folders/VirtualService.yaml")
}
