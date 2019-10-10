package synlight

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLighter(t *testing.T) {
	hl := New(Options{Format: TTY})
	data, _ := ioutil.ReadFile("synlight.go")
	hl.Highlight(data, "go", os.Stdout)
}
