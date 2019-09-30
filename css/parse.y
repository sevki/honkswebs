%{
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
)

%}

%union {
	s string
	r *Rule
}

%type <r> grammar rule rules selectors strings

%token <s> TSTRING

%%

grammar:	/* empty */ {
       		} |
		grammar rule {
			l := rulelex.(*ruleLex)
			l.rules = append(l.rules, $2)
		} ;

rule:		selectors '{' rules '}' {
			$$ = &Rule { Type: 'r' }
			names := $1.Names
			for i := 1; i < len(names) - 1; i++ {
				if names[i] == ":" {
					names[i-1] += ":" + names[i+1]
					for j := i; j < len(names) - 2; j++ {
						names[j] = names[j+2]
					}
					names = names[0:len(names)-2]
				}
			}
			$$.Names = names
			$$.Rules = $3.Rules
	   	} |
		strings ':' strings ';' {
			$$ = &Rule { Type: 's' }
			$$.Names = $1.Names
			$$.Values = $3.Names
      		} ;

rules:		/* none */ {
	 		$$ = &Rule { Type: 'r' }
	   	} |
		rules rule {
			if $2 != nil {
				$1.Rules = append($1.Rules, $2)
				$$ = $1
			}
		} ;

selectors:	strings {
	 		$$ = $1
	 	} |
		strings ':' selectors {
       			$$ = new(Rule)
			$$.Names = append($$.Names, $1.Names...)
			$$.Names = append($$.Names, ":")
			$$.Names = append($$.Names, $3.Names...)
		} ;

strings:	TSTRING {
       			$$ = new(Rule)
			$$.Names = append($$.Names, $1)
		} |
		strings TSTRING {
			$1.Names = append($1.Names, $2)
			$$ = $1
		} ;

%%
