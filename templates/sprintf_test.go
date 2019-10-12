package templates

import (
	"html/template"
	"testing"
)

func TestSprintf(t *testing.T) {
	url := `" onclick='bad' x="`
	name := template.HTML(`<b>text</b>`)
	s := Sprintf(`<a href="%s">%s</a>`, url, name)
	res := template.HTML(`<a href="&#34; onclick=&#39;bad&#39; x=&#34;"><b>text</b></a>`)
	if s != res {
		t.Errorf("%s != %s", s, res)
	}
}
