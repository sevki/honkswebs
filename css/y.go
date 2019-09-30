// Code generated by goyacc -p rule parse.y. DO NOT EDIT.

//line parse.y:2
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

import __yyfmt__ "fmt"

//line parse.y:17

import ()

//line parse.y:24
type ruleSymType struct {
	yys int
	s   string
	r   *Rule
}

const TSTRING = 57346

var ruleToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TSTRING",
	"'{'",
	"'}'",
	"':'",
	"';'",
}
var ruleStatenames = [...]string{}

const ruleEofCode = 1
const ruleErrCode = 2
const ruleInitialStackSize = 16

//line parse.y:92

//line yacctab:1
var ruleExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const rulePrivate = 57344

const ruleLast = 21

var ruleAct = [...]int{

	4, 3, 8, 2, 8, 15, 14, 15, 10, 11,
	5, 8, 12, 13, 7, 6, 16, 11, 5, 9,
	1,
}
var rulePact = [...]int{

	-1000, 14, -1000, 10, 7, -1000, -1000, 14, -1000, 6,
	-2, -1000, -1000, -1000, -1000, 14, 0,
}
var rulePgo = [...]int{

	0, 20, 3, 19, 1, 0,
}
var ruleR1 = [...]int{

	0, 1, 1, 2, 2, 3, 3, 4, 4, 5,
	5,
}
var ruleR2 = [...]int{

	0, 0, 2, 4, 4, 0, 2, 1, 3, 1,
	2,
}
var ruleChk = [...]int{

	-1000, -1, -2, -4, -5, 4, 5, 7, 4, -3,
	-5, -4, 6, -2, 8, 7, -5,
}
var ruleDef = [...]int{

	1, -2, 2, 0, 7, 9, 5, 0, 10, 0,
	7, 8, 3, 6, 4, 0, 7,
}
var ruleTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 7, 8,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 5, 3, 6,
}
var ruleTok2 = [...]int{

	2, 3, 4,
}
var ruleTok3 = [...]int{
	0,
}

var ruleErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	ruleDebug        = 0
	ruleErrorVerbose = false
)

type ruleLexer interface {
	Lex(lval *ruleSymType) int
	Error(s string)
}

type ruleParser interface {
	Parse(ruleLexer) int
	Lookahead() int
}

type ruleParserImpl struct {
	lval  ruleSymType
	stack [ruleInitialStackSize]ruleSymType
	char  int
}

func (p *ruleParserImpl) Lookahead() int {
	return p.char
}

func ruleNewParser() ruleParser {
	return &ruleParserImpl{}
}

const ruleFlag = -1000

func ruleTokname(c int) string {
	if c >= 1 && c-1 < len(ruleToknames) {
		if ruleToknames[c-1] != "" {
			return ruleToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func ruleStatname(s int) string {
	if s >= 0 && s < len(ruleStatenames) {
		if ruleStatenames[s] != "" {
			return ruleStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func ruleErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !ruleErrorVerbose {
		return "syntax error"
	}

	for _, e := range ruleErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + ruleTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := rulePact[state]
	for tok := TOKSTART; tok-1 < len(ruleToknames); tok++ {
		if n := base + tok; n >= 0 && n < ruleLast && ruleChk[ruleAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if ruleDef[state] == -2 {
		i := 0
		for ruleExca[i] != -1 || ruleExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; ruleExca[i] >= 0; i += 2 {
			tok := ruleExca[i]
			if tok < TOKSTART || ruleExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if ruleExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += ruleTokname(tok)
	}
	return res
}

func rulelex1(lex ruleLexer, lval *ruleSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = ruleTok1[0]
		goto out
	}
	if char < len(ruleTok1) {
		token = ruleTok1[char]
		goto out
	}
	if char >= rulePrivate {
		if char < rulePrivate+len(ruleTok2) {
			token = ruleTok2[char-rulePrivate]
			goto out
		}
	}
	for i := 0; i < len(ruleTok3); i += 2 {
		token = ruleTok3[i+0]
		if token == char {
			token = ruleTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = ruleTok2[1] /* unknown char */
	}
	if ruleDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", ruleTokname(token), uint(char))
	}
	return char, token
}

func ruleParse(rulelex ruleLexer) int {
	return ruleNewParser().Parse(rulelex)
}

func (rulercvr *ruleParserImpl) Parse(rulelex ruleLexer) int {
	var rulen int
	var ruleVAL ruleSymType
	var ruleDollar []ruleSymType
	_ = ruleDollar // silence set and not used
	ruleS := rulercvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	rulestate := 0
	rulercvr.char = -1
	ruletoken := -1 // rulercvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		rulestate = -1
		rulercvr.char = -1
		ruletoken = -1
	}()
	rulep := -1
	goto rulestack

ret0:
	return 0

ret1:
	return 1

rulestack:
	/* put a state and value onto the stack */
	if ruleDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", ruleTokname(ruletoken), ruleStatname(rulestate))
	}

	rulep++
	if rulep >= len(ruleS) {
		nyys := make([]ruleSymType, len(ruleS)*2)
		copy(nyys, ruleS)
		ruleS = nyys
	}
	ruleS[rulep] = ruleVAL
	ruleS[rulep].yys = rulestate

rulenewstate:
	rulen = rulePact[rulestate]
	if rulen <= ruleFlag {
		goto ruledefault /* simple state */
	}
	if rulercvr.char < 0 {
		rulercvr.char, ruletoken = rulelex1(rulelex, &rulercvr.lval)
	}
	rulen += ruletoken
	if rulen < 0 || rulen >= ruleLast {
		goto ruledefault
	}
	rulen = ruleAct[rulen]
	if ruleChk[rulen] == ruletoken { /* valid shift */
		rulercvr.char = -1
		ruletoken = -1
		ruleVAL = rulercvr.lval
		rulestate = rulen
		if Errflag > 0 {
			Errflag--
		}
		goto rulestack
	}

ruledefault:
	/* default state action */
	rulen = ruleDef[rulestate]
	if rulen == -2 {
		if rulercvr.char < 0 {
			rulercvr.char, ruletoken = rulelex1(rulelex, &rulercvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if ruleExca[xi+0] == -1 && ruleExca[xi+1] == rulestate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			rulen = ruleExca[xi+0]
			if rulen < 0 || rulen == ruletoken {
				break
			}
		}
		rulen = ruleExca[xi+1]
		if rulen < 0 {
			goto ret0
		}
	}
	if rulen == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			rulelex.Error(ruleErrorMessage(rulestate, ruletoken))
			Nerrs++
			if ruleDebug >= 1 {
				__yyfmt__.Printf("%s", ruleStatname(rulestate))
				__yyfmt__.Printf(" saw %s\n", ruleTokname(ruletoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for rulep >= 0 {
				rulen = rulePact[ruleS[rulep].yys] + ruleErrCode
				if rulen >= 0 && rulen < ruleLast {
					rulestate = ruleAct[rulen] /* simulate a shift of "error" */
					if ruleChk[rulestate] == ruleErrCode {
						goto rulestack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if ruleDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", ruleS[rulep].yys)
				}
				rulep--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if ruleDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", ruleTokname(ruletoken))
			}
			if ruletoken == ruleEofCode {
				goto ret1
			}
			rulercvr.char = -1
			ruletoken = -1
			goto rulenewstate /* try again in the same state */
		}
	}

	/* reduction by production rulen */
	if ruleDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", rulen, ruleStatname(rulestate))
	}

	rulent := rulen
	rulept := rulep
	_ = rulept // guard against "declared and not used"

	rulep -= ruleR2[rulen]
	// rulep is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if rulep+1 >= len(ruleS) {
		nyys := make([]ruleSymType, len(ruleS)*2)
		copy(nyys, ruleS)
		ruleS = nyys
	}
	ruleVAL = ruleS[rulep+1]

	/* consult goto table to find next state */
	rulen = ruleR1[rulen]
	ruleg := rulePgo[rulen]
	rulej := ruleg + ruleS[rulep].yys + 1

	if rulej >= ruleLast {
		rulestate = ruleAct[ruleg]
	} else {
		rulestate = ruleAct[rulej]
		if ruleChk[rulestate] != -rulen {
			rulestate = ruleAct[ruleg]
		}
	}
	// dummy call; replaced with literal code
	switch rulent {

	case 1:
		ruleDollar = ruleS[rulept-0 : rulept+1]
//line parse.y:35
		{
		}
	case 2:
		ruleDollar = ruleS[rulept-2 : rulept+1]
//line parse.y:37
		{
			l := rulelex.(*ruleLex)
			l.rules = append(l.rules, ruleDollar[2].r)
		}
	case 3:
		ruleDollar = ruleS[rulept-4 : rulept+1]
//line parse.y:42
		{
			ruleVAL.r = &Rule{Type: 'r'}
			names := ruleDollar[1].r.Names
			for i := 1; i < len(names)-1; i++ {
				if names[i] == ":" {
					names[i-1] += ":" + names[i+1]
					for j := i; j < len(names)-2; j++ {
						names[j] = names[j+2]
					}
					names = names[0 : len(names)-2]
				}
			}
			ruleVAL.r.Names = names
			ruleVAL.r.Rules = ruleDollar[3].r.Rules
		}
	case 4:
		ruleDollar = ruleS[rulept-4 : rulept+1]
//line parse.y:57
		{
			ruleVAL.r = &Rule{Type: 's'}
			ruleVAL.r.Names = ruleDollar[1].r.Names
			ruleVAL.r.Values = ruleDollar[3].r.Names
		}
	case 5:
		ruleDollar = ruleS[rulept-0 : rulept+1]
//line parse.y:63
		{
			ruleVAL.r = &Rule{Type: 'r'}
		}
	case 6:
		ruleDollar = ruleS[rulept-2 : rulept+1]
//line parse.y:66
		{
			if ruleDollar[2].r != nil {
				ruleDollar[1].r.Rules = append(ruleDollar[1].r.Rules, ruleDollar[2].r)
				ruleVAL.r = ruleDollar[1].r
			}
		}
	case 7:
		ruleDollar = ruleS[rulept-1 : rulept+1]
//line parse.y:73
		{
			ruleVAL.r = ruleDollar[1].r
		}
	case 8:
		ruleDollar = ruleS[rulept-3 : rulept+1]
//line parse.y:76
		{
			ruleVAL.r = new(Rule)
			ruleVAL.r.Names = append(ruleVAL.r.Names, ruleDollar[1].r.Names...)
			ruleVAL.r.Names = append(ruleVAL.r.Names, ":")
			ruleVAL.r.Names = append(ruleVAL.r.Names, ruleDollar[3].r.Names...)
		}
	case 9:
		ruleDollar = ruleS[rulept-1 : rulept+1]
//line parse.y:83
		{
			ruleVAL.r = new(Rule)
			ruleVAL.r.Names = append(ruleVAL.r.Names, ruleDollar[1].s)
		}
	case 10:
		ruleDollar = ruleS[rulept-2 : rulept+1]
//line parse.y:87
		{
			ruleDollar[1].r.Names = append(ruleDollar[1].r.Names, ruleDollar[2].s)
			ruleVAL.r = ruleDollar[1].r
		}
	}
	goto rulestack /* stack new state and value */
}
