//
// Copyright (c) 2019 Ted Unangst <tedu@tedunangst.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// hypertext filter
package htfilter

import (
	"fmt"
	"html/template"
	"io"
	"net/url"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

type Filter struct {
	Imager func (node *html.Node) string
}

func New() *Filter {
	f := new(Filter)
	f.Imager = replaceimg
	return f
}

var permittedtags = []string{
	"div", "h1", "h2", "h3", "h4", "h5", "h6", "hr",
	"table", "thead", "tbody", "th", "tr", "td", "colgroup", "col",
	"p", "br", "pre", "code", "blockquote", "q",
	"samp", "mark", "ins", "dfn", "cite", "abbr", "address",
	"strong", "em", "b", "i", "s", "u", "sub", "sup", "del", "tt", "small",
	"ol", "ul", "li", "dl", "dt", "dd",
}
var permittedattr = []string{"colspan", "rowspan"}
var bannedtags = []string{"script", "style"}

func init() {
	sort.Strings(permittedtags)
	sort.Strings(permittedattr)
	sort.Strings(bannedtags)
}

func contains(array []string, tag string) bool {
	idx := sort.SearchStrings(array, tag)
	return idx < len(array) && array[idx] == tag
}

func GetAttr(node *html.Node, attr string) string {
	for _, a := range node.Attr {
		if a.Key == attr {
			return a.Val
		}
	}
	return ""
}

func HasClass(node *html.Node, class string) bool {
	return strings.Contains(" "+GetAttr(node, "class")+" ", " "+class+" ")
}

func writetag(w io.Writer, node *html.Node) {
	io.WriteString(w, "<")
	io.WriteString(w, node.Data)
	for _, attr := range node.Attr {
		if contains(permittedattr, attr.Key) {
			fmt.Fprintf(w, ` %s="%s"`, attr.Key, html.EscapeString(attr.Val))
		}
	}
	io.WriteString(w, ">")
}

func (filt *Filter) render(w io.Writer, node *html.Node) {
	if node.Type == html.ElementNode {
		tag := node.Data
		switch {
		case tag == "a":
			href := GetAttr(node, "href")
			hrefurl, err := url.Parse(href)
			if err != nil {
				href = "#BROKEN-" + href
			} else {
				href = hrefurl.String()
			}
			fmt.Fprintf(w, `<a href="%s" rel=noreferrer>`, html.EscapeString(href))
		case tag == "img":
			div := filt.Imager(node)
			if div != "skip" {
				io.WriteString(w, html.EscapeString(div))
			}
		case tag == "span":
		case tag == "iframe":
			src := html.EscapeString(GetAttr(node, "src"))
			fmt.Fprintf(w, `&lt;iframe src="<a href="%s">%s</a>"&gt;`, src, src)
		case contains(permittedtags, tag):
			writetag(w, node)
		case contains(bannedtags, tag):
			return
		}
	} else if node.Type == html.TextNode {
		io.WriteString(w, html.EscapeString(node.Data))
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		filt.render(w, c)
	}

	if node.Type == html.ElementNode {
		tag := node.Data
		if tag == "a" || (contains(permittedtags, tag) && tag != "br") {
			fmt.Fprintf(w, "</%s>", tag)
		}
		if tag == "p" || tag == "div" {
			io.WriteString(w, "\n")
		}
	}
}

func replaceimg(node *html.Node) string {
	src := GetAttr(node, "src")
	alt := GetAttr(node, "alt")
	//title := GetAttr(node, "title")
	if HasClass(node, "Emoji") && alt != "" {
		return alt
	}
	return fmt.Sprintf(`<img src="%s">`, src)
}

func (filt *Filter) cleannode(node *html.Node) template.HTML {
	var buf strings.Builder
	filt.render(&buf, node)
	return template.HTML(buf.String())
}

func (filt *Filter) String(shtml string) (template.HTML, error) {
	reader := strings.NewReader(shtml)
	body, err := html.Parse(reader)
	if err != nil {
		return "", err
	}
	return filt.cleannode(body), nil
}

func (filt *Filter) TextOnly(node *html.Node) string {
	var buf strings.Builder
	filt.gathertext(&buf, node, false)
	return buf.String()
}

func (filt *Filter) gathertext(w io.Writer, node *html.Node, withlinks bool) {
	switch node.Type {
	case html.ElementNode:
		tag := node.Data
		switch {
		case tag == "a":
			fmt.Fprintf(w, " ")
			if withlinks {
				href := GetAttr(node, "href")
				fmt.Fprintf(w, `<a href="%s">`, href)
			}
		case tag == "img":
			div := filt.Imager(node)
			io.WriteString(w, div)
		case tag == "span":
			if HasClass(node, "tco-ellipsis") {
				return
			}
		case contains(bannedtags, tag):
			return
		}
	case html.TextNode:
		io.WriteString(w, node.Data)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		filt.gathertext(w, c, withlinks)
	}
	if node.Type == html.ElementNode {
		tag := node.Data
		if withlinks && tag == "a" {
			fmt.Fprintf(w, "</%s>", tag)
		}
		if tag == "p" || tag == "div" {
			io.WriteString(w, "\n")
		}
	}
}
