//line vmparser.y:1
package vmparse

import __yyfmt__ "fmt"

//line vmparser.y:3
//line vmparser.y:5
type yySymType struct {
	yys  int
	node *Node
}

const NUM = 57346
const IDENT = 57347

var yyToknames = []string{
	"NUM",
	"IDENT",
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'^'",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 17
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 49

var yyAct = []int{

	3, 30, 31, 14, 18, 15, 29, 16, 17, 24,
	19, 20, 21, 22, 23, 25, 26, 2, 1, 28,
	9, 10, 11, 12, 13, 0, 4, 5, 0, 27,
	7, 0, 32, 11, 12, 13, 0, 6, 9, 10,
	11, 12, 13, 8, 9, 10, 11, 12, 13,
}
var yyPact = []int{

	-1000, 22, -1000, 31, -1000, -10, 22, 22, -2, 22,
	22, 22, 22, 22, 5, 22, 13, 24, 22, 24,
	24, -1000, -1000, -1000, -8, -15, 37, -1000, 37, -1000,
	-1000, 22, 37,
}
var yyPgo = []int{

	0, 18, 17, 0, 15,
}
var yyR1 = []int{

	0, 1, 1, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4,
}
var yyR2 = []int{

	0, 0, 2, 4, 1, 1, 4, 3, 4, 2,
	3, 3, 3, 3, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, 4, 5, 15, 8, 12, 7,
	8, 9, 10, 11, 13, 15, -3, -3, 6, -3,
	-3, -3, -3, -3, 4, -4, -3, 16, -3, 14,
	16, 17, -3,
}
var yyDef = []int{

	1, -2, 2, 0, 4, 5, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 9, 0, 10,
	11, 12, 13, 14, 0, 0, 15, 7, 3, 6,
	8, 0, 16,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	15, 16, 9, 7, 17, 8, 3, 10, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 12, 3,
	3, 6, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 13, 3, 14, 11,
}
var yyTok2 = []int{

	2, 3, 4, 5,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line vmparser.y:20
		{
		}
	case 2:
		//line vmparser.y:21
		{
			c := yylex.(*context)
			c.statements = append(c.statements, yyS[yypt-0].node)
		}
	case 3:
		//line vmparser.y:24
		{ /* lhs should be exp too!!! */ yyVAL.node = NewStatement(yyS[yypt-3].node, yyS[yypt-0].node)
		}
	case 4:
		//line vmparser.y:27
		{
			yyVAL = yyS[yypt-0]
		}
	case 5:
		//line vmparser.y:28
		{
			yyVAL = yyS[yypt-0]
		}
	case 6:
		//line vmparser.y:29
		{
			yyVAL.node = IndexedIdentifier(yyS[yypt-3].node, yyS[yypt-1].node)
		}
	case 7:
		//line vmparser.y:30
		{
			yyVAL = yyS[yypt-1]
		}
	case 8:
		//line vmparser.y:31
		{
			yyVAL.node = FunctionArgs(yyS[yypt-3].node.S, yyS[yypt-1].node)
		}
	case 9:
		//line vmparser.y:32
		{
			yyVAL.node = Negate(yyS[yypt-0].node)
		}
	case 10:
		//line vmparser.y:33
		{
			yyVAL.node = Function("add", yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 11:
		//line vmparser.y:34
		{
			yyVAL.node = Function("subtract", yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 12:
		//line vmparser.y:35
		{
			yyVAL.node = Function("multiply", yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 13:
		//line vmparser.y:36
		{
			yyVAL.node = Function("divide", yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 14:
		//line vmparser.y:37
		{
			yyVAL.node = Function("pow", yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 15:
		//line vmparser.y:40
		{
			yyVAL.node = NewArgList(yyS[yypt-0].node)
		}
	case 16:
		//line vmparser.y:41
		{
			yyVAL.node = yyS[yypt-2].node.AddChild(yyS[yypt-0].node)
		}
	}
	goto yystack /* stack new state and value */
}
