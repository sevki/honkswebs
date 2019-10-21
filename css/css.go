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

// css filtering, possibly a limited set of scss, but decidedly not all of it
package css

import (
	"fmt"
	"io"
)

type rule struct {
	Rules  []*rule
	Names  []string
	Values []string
	Type   byte
}

// Read a css like file, expand macros and nestings, and write to w.
func Filter(reader io.Reader, w io.Writer) error {
	ruleErrorVerbose = true
	lexer := newRuleLexer(reader)
	ruleParse(lexer)
	if lexer.err != nil {
		return lexer.err
	}
	rules := lexer.rules

	vars := make(map[string][]string)
	var namestack [][]string

	var printRule func(rule *rule)
	printRule = func(rule *rule) {
		switch rule.Type {
		case 'r':
			if rule.Names[0][0] == '@' {
				for _, n := range rule.Names {
					fmt.Fprintf(w, "%s ", n)
				}
				fmt.Fprintf(w, "{\n")
				for _, r := range rule.Rules {
					if r.Type == 's' {
						printRule(r)
					}
				}
				for _, r := range rule.Rules {
					if r.Type == 'r' {
						printRule(r)
					}
				}
				fmt.Fprintf(w, "}\n")
			} else {
				namestack = append(namestack, rule.Names)
				for _, names := range namestack {
					for _, n := range names {
						fmt.Fprintf(w, "%s ", n)
					}
				}
				fmt.Fprintf(w, "{\n")
				for _, r := range rule.Rules {
					if r.Type == 's' {
						printRule(r)
					}
				}
				fmt.Fprintf(w, "}\n")
				for _, r := range rule.Rules {
					if r.Type == 'r' {
						printRule(r)
					}
				}
				namestack = namestack[0 : len(namestack)-1]
			}
		case 's':
			if rule.Names[0][0] == '$' {
				name := rule.Names[0]
				vars[name] = rule.Values
			} else {
				for _, n := range rule.Names {
					fmt.Fprintf(w, "%s ", n)
				}
				fmt.Fprintf(w, ":")
				for _, v := range rule.Values {
					if v[0] == '$' {
						vals := vars[v]
						for _, vv := range vals {
							fmt.Fprintf(w, " %s", vv)
						}
					} else {
						fmt.Fprintf(w, " %s", v)
					}
				}
				fmt.Fprintf(w, ";\n")
			}
		}
	}

	for _, rule := range rules {
		printRule(rule)
	}

	return nil
}
