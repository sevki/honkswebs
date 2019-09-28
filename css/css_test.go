package css

import (
	"os"
	"testing"
)

func TestFilter(t *testing.T) {
	fd, _ := os.Open("input.scss")
	err := Filter(fd, os.Stdout)
	if err != nil {
		t.Error(err)
	}
}
