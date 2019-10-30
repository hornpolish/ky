package cmd

import (
	"io/ioutil"
	"strings"
	"testing"
)

var diffTests = []struct {
	f1    string
	f2    string
	color bool
}{
	{"a", "a2", true},
	{"a", "a3", false},
	{"b", "b2", false},
	{"c", "c2", false},
	{"c", "c3", true},
	{"c2", "c3", false},
}

func TestDiff(t *testing.T) {
	for _, tt := range diffTests {
		diffVar.color = tt.color
		diffVar.nocolor = false

		got, err := diff2("../test/"+tt.f1+".yaml", "../test/"+tt.f2+".yaml")
		if err != nil {
			t.Errorf("diff %s-%s diff2 failed.", tt.f1, tt.f2)
			continue
		}

		res := "../test/expected/diff-" + tt.f1 + "-" + tt.f2
		want, err := ioutil.ReadFile(res)
		if err != nil {
			t.Errorf("diff %s-%s ReadFile(%s) failed.", tt.f1, tt.f2, res)
			continue
		}

		wot := strings.TrimSuffix(string(want), "\n")
		cmp := strings.Compare(got, wot)
		if cmp != 0 {
			t.Errorf("diff %s-%s compare failed. got...\n%s\n...wanted...\n%s\n... boo!", tt.f1, tt.f2, got, wot)
		}
	}
}
