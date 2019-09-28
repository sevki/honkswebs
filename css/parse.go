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

package css

import (
	"bufio"
	"fmt"
	"io"
)

func newRuleLexer(r io.Reader) *ruleLex {
	return &ruleLex{r: bufio.NewReader(r)}
}

type ruleLex struct {
	r     *bufio.Reader
	rules []*Rule
	err   error
}

func (x *ruleLex) Lex(lval *ruleSymType) int {
	var b byte
	var err error

	for b, err = x.r.ReadByte(); err == nil && (b == ' ' || b == '\t' || b == '\n'); b, err = x.r.ReadByte() {
	}
	if err != nil {
		return 0
	}
	switch b {
	case ';':
		fallthrough
	case ':':
		fallthrough
	case '}':
		fallthrough
	case '{':
		return int(b)
	}
	x.r.UnreadByte()
	bb := make([]byte, 0, 16)
	inq := false
	for b, err = x.r.ReadByte(); err == nil; b, err = x.r.ReadByte() {
		if b == '(' {
			inq = true
		}
		if inq == false && (b == ' ' || b == '\t' || b == '\n') {
			break
		}
		if inq == false && (b == ';' || b == ':') {
			x.r.UnreadByte()
			break
		}
		bb = append(bb, b)
		if b == ')' {
			break
		}
	}
	if err != nil {
		return 0
	}
	t := string(bb)

	lval.s = t

	return TSTRING
}

func (x *ruleLex) Error(s string) {
	x.err = fmt.Errorf("parse error: %s", s)
}
