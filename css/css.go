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

// baby sass
package css

import (
	"regexp"
	"strings"
)

var varname_re = `\$[[:alpha:]][[:alnum:]_-]+`
var re_vardecls = regexp.MustCompile(`(?m)(` + varname_re +`):?(.*);`)

func Process(s string) string {
	vars := make(map[string]string)

	replfn := func(m string) string {
		m = m[:len(m)-1]
		if strings.IndexByte(m, ':') == -1 {
			v, ok := vars[m]
			if !ok {
				v = m
			}
			return v + ";"
		}
		v := strings.SplitN(m, ":", 2)
		vars[v[0]] = v[1]
		return ""
	}
	s = re_vardecls.ReplaceAllStringFunc(s, replfn)

	return s
}
