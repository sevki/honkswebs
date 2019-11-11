package htfilter

import (
	"io/ioutil"
	"testing"
)

func BenchmarkScrubLargeFrag(b *testing.B) {
	data, err := ioutil.ReadFile("fragment-large.html")
	if err != nil {
		b.FailNow()
		return
	}
	scrubeone(b, string(data))
}

func BenchmarkScrubSmallFrag(b *testing.B) {
	data, err := ioutil.ReadFile("fragment-small.html")
	if err != nil {
		b.FailNow()
		return
	}
	scrubeone(b, string(data))
}

func BenchmarkScrubLargeDoc(b *testing.B) {
	data, err := ioutil.ReadFile("document-large.html")
	if err != nil {
		b.FailNow()
		return
	}
	scrubeone(b, string(data))
}

func scrubeone(b *testing.B, src string) {
	var htf Filter
	for n := 0; n < b.N; n++ {
		htf.String(src)
	}
}
