// Code generated by goyacc -o pkg/traceql/expr.y.go pkg/traceql/expr.y. DO NOT EDIT.

//line pkg/traceql/expr.y:2
package traceql

import __yyfmt__ "fmt"

//line pkg/traceql/expr.y:2

import (
	"time"
)

//line pkg/traceql/expr.y:11
type yySymType struct {
	yys               int
	root              RootExpr
	groupOperation    GroupOperation
	coalesceOperation CoalesceOperation

	spansetExpression         SpansetExpression
	spansetPipelineExpression SpansetExpression
	wrappedSpansetPipeline    Pipeline
	spansetPipeline           Pipeline
	spansetFilter             SpansetFilter
	scalarFilter              ScalarFilter
	scalarFilterOperation     Operator

	scalarPipelineExpressionFilter ScalarFilter
	scalarPipelineExpression       ScalarExpression
	scalarExpression               ScalarExpression
	wrappedScalarPipeline          Pipeline
	scalarPipeline                 Pipeline
	aggregate                      Aggregate

	fieldExpression FieldExpression
	static          Static
	intrinsicField  Attribute
	attributeField  Attribute

	binOp          Operator
	staticInt      int
	staticStr      string
	staticFloat    float64
	staticDuration time.Duration
}

const IDENTIFIER = 57346
const STRING = 57347
const INTEGER = 57348
const FLOAT = 57349
const DURATION = 57350
const DOT = 57351
const OPEN_BRACE = 57352
const CLOSE_BRACE = 57353
const OPEN_PARENS = 57354
const CLOSE_PARENS = 57355
const NIL = 57356
const TRUE = 57357
const FALSE = 57358
const STATUS_ERROR = 57359
const STATUS_OK = 57360
const STATUS_UNSET = 57361
const IDURATION = 57362
const CHILDCOUNT = 57363
const NAME = 57364
const STATUS = 57365
const PARENT = 57366
const PARENT_DOT = 57367
const RESOURCE_DOT = 57368
const SPAN_DOT = 57369
const COUNT = 57370
const AVG = 57371
const MAX = 57372
const MIN = 57373
const SUM = 57374
const BY = 57375
const COALESCE = 57376
const END_ATTRIBUTE = 57377
const PIPE = 57378
const AND = 57379
const OR = 57380
const EQ = 57381
const NEQ = 57382
const LT = 57383
const LTE = 57384
const GT = 57385
const GTE = 57386
const NRE = 57387
const RE = 57388
const DESC = 57389
const TILDE = 57390
const ADD = 57391
const SUB = 57392
const NOT = 57393
const MUL = 57394
const DIV = 57395
const MOD = 57396
const POW = 57397

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENTIFIER",
	"STRING",
	"INTEGER",
	"FLOAT",
	"DURATION",
	"DOT",
	"OPEN_BRACE",
	"CLOSE_BRACE",
	"OPEN_PARENS",
	"CLOSE_PARENS",
	"NIL",
	"TRUE",
	"FALSE",
	"STATUS_ERROR",
	"STATUS_OK",
	"STATUS_UNSET",
	"IDURATION",
	"CHILDCOUNT",
	"NAME",
	"STATUS",
	"PARENT",
	"PARENT_DOT",
	"RESOURCE_DOT",
	"SPAN_DOT",
	"COUNT",
	"AVG",
	"MAX",
	"MIN",
	"SUM",
	"BY",
	"COALESCE",
	"END_ATTRIBUTE",
	"PIPE",
	"AND",
	"OR",
	"EQ",
	"NEQ",
	"LT",
	"LTE",
	"GT",
	"GTE",
	"NRE",
	"RE",
	"DESC",
	"TILDE",
	"ADD",
	"SUB",
	"NOT",
	"MUL",
	"DIV",
	"MOD",
	"POW",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 637

var yyAct = [...]int{

	75, 17, 6, 7, 5, 169, 2, 149, 12, 17,
	69, 56, 111, 46, 45, 112, 33, 49, 116, 136,
	137, 204, 138, 139, 140, 149, 64, 65, 203, 66,
	67, 68, 69, 195, 17, 33, 92, 94, 93, 138,
	139, 140, 149, 194, 104, 106, 107, 108, 109, 193,
	192, 118, 202, 168, 64, 65, 71, 66, 67, 68,
	69, 15, 161, 105, 17, 17, 17, 17, 17, 17,
	17, 115, 126, 128, 129, 130, 131, 132, 133, 141,
	142, 143, 144, 145, 146, 148, 147, 47, 10, 136,
	137, 119, 138, 139, 140, 149, 99, 17, 197, 111,
	17, 166, 91, 51, 52, 167, 53, 54, 55, 56,
	166, 90, 89, 17, 40, 92, 94, 93, 41, 43,
	17, 171, 112, 88, 87, 173, 201, 134, 17, 152,
	153, 154, 66, 67, 68, 69, 167, 70, 117, 120,
	121, 122, 123, 124, 125, 162, 163, 164, 165, 114,
	150, 151, 141, 142, 143, 144, 145, 146, 148, 147,
	158, 63, 136, 137, 196, 138, 139, 140, 149, 157,
	156, 17, 50, 17, 155, 46, 77, 46, 173, 49,
	76, 49, 159, 160, 16, 51, 52, 48, 53, 54,
	55, 56, 14, 175, 176, 177, 178, 179, 180, 181,
	182, 183, 184, 185, 186, 187, 188, 189, 190, 23,
	24, 25, 29, 83, 4, 11, 72, 9, 28, 26,
	27, 31, 30, 32, 78, 79, 80, 81, 82, 86,
	84, 85, 200, 57, 58, 59, 60, 61, 62, 53,
	54, 55, 56, 64, 65, 95, 66, 67, 68, 69,
	1, 199, 0, 0, 73, 74, 150, 151, 141, 142,
	143, 144, 145, 146, 148, 147, 0, 0, 136, 137,
	198, 138, 139, 140, 149, 150, 151, 141, 142, 143,
	144, 145, 146, 148, 147, 0, 0, 136, 137, 191,
	138, 139, 140, 149, 150, 151, 141, 142, 143, 144,
	145, 146, 148, 147, 0, 0, 136, 137, 174, 138,
	139, 140, 149, 150, 151, 141, 142, 143, 144, 145,
	146, 148, 147, 0, 0, 136, 137, 135, 138, 139,
	140, 149, 150, 151, 141, 142, 143, 144, 145, 146,
	148, 147, 0, 0, 136, 137, 116, 138, 139, 140,
	149, 0, 0, 150, 151, 141, 142, 143, 144, 145,
	146, 148, 147, 0, 0, 136, 137, 0, 138, 139,
	140, 149, 57, 58, 59, 60, 61, 62, 0, 0,
	0, 0, 64, 65, 113, 66, 67, 68, 69, 57,
	58, 59, 60, 61, 62, 0, 110, 0, 0, 51,
	52, 0, 53, 54, 55, 56, 39, 42, 39, 42,
	0, 0, 40, 0, 40, 0, 41, 43, 41, 43,
	34, 37, 0, 35, 0, 0, 35, 36, 38, 0,
	36, 38, 23, 24, 25, 29, 0, 15, 0, 96,
	0, 28, 26, 27, 31, 30, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 18, 21, 19, 20, 22,
	13, 97, 34, 37, 0, 0, 0, 0, 35, 0,
	0, 0, 36, 38, 23, 24, 25, 29, 0, 15,
	0, 172, 0, 28, 26, 27, 31, 30, 32, 0,
	0, 0, 0, 0, 0, 0, 0, 18, 21, 19,
	20, 22, 13, 23, 24, 25, 29, 0, 15, 0,
	170, 0, 28, 26, 27, 31, 30, 32, 0, 0,
	0, 0, 0, 0, 0, 0, 18, 21, 19, 20,
	22, 13, 23, 24, 25, 29, 0, 15, 0, 8,
	0, 28, 26, 27, 31, 30, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 18, 21, 19, 20, 22,
	13, 23, 24, 25, 29, 0, 15, 0, 96, 0,
	28, 26, 27, 31, 30, 32, 0, 0, 0, 0,
	0, 0, 44, 3, 18, 21, 19, 20, 22, 23,
	24, 25, 29, 0, 0, 0, 127, 0, 28, 26,
	27, 31, 30, 32, 0, 0, 0, 0, 0, 0,
	0, 0, 18, 21, 19, 20, 22, 98, 100, 101,
	102, 103, 23, 24, 25, 29, 0, 0, 0, 119,
	0, 28, 26, 27, 31, 30, 32,
}
var yyPact = [...]int{

	527, -1000, -20, 425, -1000, 369, -1000, -1000, 527, -1000,
	350, -1000, 194, 125, -1000, 204, -1000, -1000, 112, 111,
	100, 99, 90, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 427, 84, 84, 84, 84, 84, 51,
	51, 51, 51, 51, 383, 86, 371, 136, 58, 333,
	617, 79, 79, 79, 79, 79, 79, -1000, -1000, -1000,
	-1000, -1000, -1000, 584, 584, 584, 584, 584, 584, 584,
	204, 316, 204, 204, 204, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 170, 166, 165, 156, 49, 204, 204,
	204, 204, -1000, 369, -1000, -1000, 556, 41, 380, 498,
	-1000, -1000, 380, -1000, 71, 51, -1000, -1000, 71, -1000,
	-1000, -1000, 427, -1000, -1000, -1000, -1000, 54, -1000, 469,
	187, 187, -44, -44, -44, -44, -23, 584, 80, 80,
	-45, -45, -45, -45, 295, -1000, 204, 204, 204, 204,
	204, 204, 204, 204, 204, 204, 204, 204, 204, 204,
	204, 204, 276, -13, -13, 15, 14, 8, -2, 160,
	94, -1000, 257, 238, 219, 113, 371, 5, 39, -1,
	498, 194, 469, -21, -1000, -13, -13, -48, -48, -48,
	-30, -30, -30, -30, -30, -30, -30, -30, -48, 40,
	40, -1000, -1000, -1000, -1000, -1000, -7, -14, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 250, 3, 245, 4, 582, 217, 5, 215, 2,
	161, 214, 87, 8, 192, 187, 184, 56, 0, 180,
	176,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 5, 5, 5, 5, 5, 5,
	5, 6, 7, 7, 7, 7, 7, 7, 7, 2,
	3, 4, 4, 4, 4, 4, 4, 4, 8, 9,
	10, 10, 10, 10, 10, 10, 11, 11, 12, 12,
	12, 12, 12, 12, 12, 12, 14, 15, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 16, 16, 16,
	16, 16, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 18, 18, 18, 18, 18, 18,
	18, 18, 18, 18, 19, 19, 19, 19, 19, 20,
	20, 20, 20, 20, 20,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 3, 3, 3, 3, 3, 3,
	1, 3, 1, 1, 1, 3, 3, 3, 3, 4,
	3, 3, 3, 3, 3, 3, 3, 1, 3, 3,
	1, 1, 1, 1, 1, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 1, 1, 3, 4, 4,
	4, 4, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	3, 3, 3, 4, 4,
}
var yyChk = [...]int{

	-1000, -1, -7, -5, -11, -4, -9, -2, 12, -6,
	-12, -8, -13, 33, -14, 10, -16, -18, 28, 30,
	31, 29, 32, 5, 6, 7, 15, 16, 14, 8,
	18, 17, 19, 36, 37, 43, 47, 38, 48, 37,
	43, 47, 38, 48, -5, -7, -4, -12, -15, -13,
	-10, 49, 50, 52, 53, 54, 55, 39, 40, 41,
	42, 43, 44, -10, 49, 50, 52, 53, 54, 55,
	12, -17, 12, 50, 51, -18, -19, -20, 20, 21,
	22, 23, 24, 9, 26, 27, 25, 12, 12, 12,
	12, 12, -9, -4, -2, -3, 12, 34, -5, 12,
	-5, -5, -5, -5, -4, 12, -4, -4, -4, -4,
	13, 13, 36, 13, 13, 13, 13, -12, -18, 12,
	-12, -12, -12, -12, -12, -12, -13, 12, -13, -13,
	-13, -13, -13, -13, -17, 11, 49, 50, 52, 53,
	54, 39, 40, 41, 42, 43, 44, 46, 45, 55,
	37, 38, -17, -17, -17, 4, 4, 4, 4, 26,
	27, 13, -17, -17, -17, -17, -4, -13, 12, -7,
	12, -13, 12, -7, 13, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, 13, 35, 35, 35, 35, 4, 4, 13, 13,
	13, 13, 13, 35, 35,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 12, 13, 14, 0, 10,
	0, 27, 0, 0, 45, 0, 55, 56, 0, 0,
	0, 0, 0, 84, 85, 86, 87, 88, 89, 90,
	91, 92, 93, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 12, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 30, 31, 32,
	33, 34, 35, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 81, 82, 83, 94, 95,
	96, 97, 98, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 15, 16, 17, 18, 0, 0, 5, 0,
	6, 7, 8, 9, 22, 0, 23, 24, 25, 26,
	4, 11, 0, 21, 38, 46, 48, 36, 37, 0,
	39, 40, 41, 42, 43, 44, 29, 0, 49, 50,
	51, 52, 53, 54, 0, 28, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 79, 80, 0, 0, 0, 0, 0,
	0, 57, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 47, 0, 0, 19, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 73, 74, 75, 76, 77,
	78, 62, 99, 100, 101, 102, 0, 0, 58, 59,
	60, 61, 20, 103, 104,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
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

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
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
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
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
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
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
			if yyn < 0 || yyn == yytoken {
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
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
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
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
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
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:93
		{
			yylex.(*lexer).expr = newRootExpr(yyDollar[1].spansetPipeline)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:94
		{
			yylex.(*lexer).expr = newRootExpr(yyDollar[1].spansetPipelineExpression)
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:95
		{
			yylex.(*lexer).expr = newRootExpr(yyDollar[1].scalarPipelineExpressionFilter)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:102
		{
			yyVAL.spansetPipelineExpression = yyDollar[2].spansetPipelineExpression
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:103
		{
			yyVAL.spansetPipelineExpression = newSpansetOperation(OpSpansetAnd, yyDollar[1].spansetPipelineExpression, yyDollar[3].spansetPipelineExpression)
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:104
		{
			yyVAL.spansetPipelineExpression = newSpansetOperation(OpSpansetChild, yyDollar[1].spansetPipelineExpression, yyDollar[3].spansetPipelineExpression)
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:105
		{
			yyVAL.spansetPipelineExpression = newSpansetOperation(OpSpansetDescendant, yyDollar[1].spansetPipelineExpression, yyDollar[3].spansetPipelineExpression)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:106
		{
			yyVAL.spansetPipelineExpression = newSpansetOperation(OpSpansetUnion, yyDollar[1].spansetPipelineExpression, yyDollar[3].spansetPipelineExpression)
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:107
		{
			yyVAL.spansetPipelineExpression = newSpansetOperation(OpSpansetSibling, yyDollar[1].spansetPipelineExpression, yyDollar[3].spansetPipelineExpression)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:108
		{
			yyVAL.spansetPipelineExpression = yyDollar[1].wrappedSpansetPipeline
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:112
		{
			yyVAL.wrappedSpansetPipeline = yyDollar[2].spansetPipeline
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:115
		{
			yyVAL.spansetPipeline = newPipeline(yyDollar[1].spansetExpression)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:116
		{
			yyVAL.spansetPipeline = newPipeline(yyDollar[1].scalarFilter)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:117
		{
			yyVAL.spansetPipeline = newPipeline(yyDollar[1].groupOperation)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:118
		{
			yyVAL.spansetPipeline = yyDollar[1].spansetPipeline.addItem(yyDollar[3].scalarFilter)
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:119
		{
			yyVAL.spansetPipeline = yyDollar[1].spansetPipeline.addItem(yyDollar[3].spansetExpression)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:120
		{
			yyVAL.spansetPipeline = yyDollar[1].spansetPipeline.addItem(yyDollar[3].groupOperation)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:121
		{
			yyVAL.spansetPipeline = yyDollar[1].spansetPipeline.addItem(yyDollar[3].coalesceOperation)
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
//line pkg/traceql/expr.y:125
		{
			yyVAL.groupOperation = newGroupOperation(yyDollar[3].fieldExpression)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:129
		{
			yyVAL.coalesceOperation = newCoalesceOperation()
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:133
		{
			yyVAL.spansetExpression = yyDollar[2].spansetExpression
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:134
		{
			yyVAL.spansetExpression = newSpansetOperation(OpSpansetAnd, yyDollar[1].spansetExpression, yyDollar[3].spansetExpression)
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:135
		{
			yyVAL.spansetExpression = newSpansetOperation(OpSpansetChild, yyDollar[1].spansetExpression, yyDollar[3].spansetExpression)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:136
		{
			yyVAL.spansetExpression = newSpansetOperation(OpSpansetDescendant, yyDollar[1].spansetExpression, yyDollar[3].spansetExpression)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:137
		{
			yyVAL.spansetExpression = newSpansetOperation(OpSpansetUnion, yyDollar[1].spansetExpression, yyDollar[3].spansetExpression)
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:138
		{
			yyVAL.spansetExpression = newSpansetOperation(OpSpansetSibling, yyDollar[1].spansetExpression, yyDollar[3].spansetExpression)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:139
		{
			yyVAL.spansetExpression = yyDollar[1].spansetFilter
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:143
		{
			yyVAL.spansetFilter = newSpansetFilter(yyDollar[2].fieldExpression)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:147
		{
			yyVAL.scalarFilter = newScalarFilter(yyDollar[2].scalarFilterOperation, yyDollar[1].scalarExpression, yyDollar[3].scalarExpression)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:151
		{
			yyVAL.scalarFilterOperation = OpEqual
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:152
		{
			yyVAL.scalarFilterOperation = OpNotEqual
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:153
		{
			yyVAL.scalarFilterOperation = OpLess
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:154
		{
			yyVAL.scalarFilterOperation = OpLessEqual
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:155
		{
			yyVAL.scalarFilterOperation = OpGreater
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:156
		{
			yyVAL.scalarFilterOperation = OpGreaterEqual
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:163
		{
			yyVAL.scalarPipelineExpressionFilter = newScalarFilter(yyDollar[2].scalarFilterOperation, yyDollar[1].scalarPipelineExpression, yyDollar[3].scalarPipelineExpression)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:164
		{
			yyVAL.scalarPipelineExpressionFilter = newScalarFilter(yyDollar[2].scalarFilterOperation, yyDollar[1].scalarPipelineExpression, yyDollar[3].static)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:168
		{
			yyVAL.scalarPipelineExpression = yyDollar[2].scalarPipelineExpression
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:169
		{
			yyVAL.scalarPipelineExpression = newScalarOperation(OpAdd, yyDollar[1].scalarPipelineExpression, yyDollar[3].scalarPipelineExpression)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:170
		{
			yyVAL.scalarPipelineExpression = newScalarOperation(OpSub, yyDollar[1].scalarPipelineExpression, yyDollar[3].scalarPipelineExpression)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:171
		{
			yyVAL.scalarPipelineExpression = newScalarOperation(OpMult, yyDollar[1].scalarPipelineExpression, yyDollar[3].scalarPipelineExpression)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:172
		{
			yyVAL.scalarPipelineExpression = newScalarOperation(OpDiv, yyDollar[1].scalarPipelineExpression, yyDollar[3].scalarPipelineExpression)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:173
		{
			yyVAL.scalarPipelineExpression = newScalarOperation(OpMod, yyDollar[1].scalarPipelineExpression, yyDollar[3].scalarPipelineExpression)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:174
		{
			yyVAL.scalarPipelineExpression = newScalarOperation(OpPower, yyDollar[1].scalarPipelineExpression, yyDollar[3].scalarPipelineExpression)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:175
		{
			yyVAL.scalarPipelineExpression = yyDollar[1].wrappedScalarPipeline
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:179
		{
			yyVAL.wrappedScalarPipeline = yyDollar[2].scalarPipeline
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:183
		{
			yyVAL.scalarPipeline = yyDollar[1].spansetPipeline.addItem(yyDollar[3].scalarExpression)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:187
		{
			yyVAL.scalarExpression = yyDollar[2].scalarExpression
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:188
		{
			yyVAL.scalarExpression = newScalarOperation(OpAdd, yyDollar[1].scalarExpression, yyDollar[3].scalarExpression)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:189
		{
			yyVAL.scalarExpression = newScalarOperation(OpSub, yyDollar[1].scalarExpression, yyDollar[3].scalarExpression)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:190
		{
			yyVAL.scalarExpression = newScalarOperation(OpMult, yyDollar[1].scalarExpression, yyDollar[3].scalarExpression)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:191
		{
			yyVAL.scalarExpression = newScalarOperation(OpDiv, yyDollar[1].scalarExpression, yyDollar[3].scalarExpression)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:192
		{
			yyVAL.scalarExpression = newScalarOperation(OpMod, yyDollar[1].scalarExpression, yyDollar[3].scalarExpression)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:193
		{
			yyVAL.scalarExpression = newScalarOperation(OpPower, yyDollar[1].scalarExpression, yyDollar[3].scalarExpression)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:194
		{
			yyVAL.scalarExpression = yyDollar[1].aggregate
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:195
		{
			yyVAL.scalarExpression = yyDollar[1].static
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:199
		{
			yyVAL.aggregate = newAggregate(aggregateCount, nil)
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
//line pkg/traceql/expr.y:200
		{
			yyVAL.aggregate = newAggregate(aggregateMax, yyDollar[3].fieldExpression)
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
//line pkg/traceql/expr.y:201
		{
			yyVAL.aggregate = newAggregate(aggregateMin, yyDollar[3].fieldExpression)
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
//line pkg/traceql/expr.y:202
		{
			yyVAL.aggregate = newAggregate(aggregateAvg, yyDollar[3].fieldExpression)
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
//line pkg/traceql/expr.y:203
		{
			yyVAL.aggregate = newAggregate(aggregateSum, yyDollar[3].fieldExpression)
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:210
		{
			yyVAL.fieldExpression = yyDollar[2].fieldExpression
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:211
		{
			yyVAL.fieldExpression = newBinaryOperation(OpAdd, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:212
		{
			yyVAL.fieldExpression = newBinaryOperation(OpSub, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:213
		{
			yyVAL.fieldExpression = newBinaryOperation(OpMult, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:214
		{
			yyVAL.fieldExpression = newBinaryOperation(OpDiv, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:215
		{
			yyVAL.fieldExpression = newBinaryOperation(OpMod, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:216
		{
			yyVAL.fieldExpression = newBinaryOperation(OpEqual, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:217
		{
			yyVAL.fieldExpression = newBinaryOperation(OpNotEqual, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:218
		{
			yyVAL.fieldExpression = newBinaryOperation(OpLess, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:219
		{
			yyVAL.fieldExpression = newBinaryOperation(OpLessEqual, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:220
		{
			yyVAL.fieldExpression = newBinaryOperation(OpGreater, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:221
		{
			yyVAL.fieldExpression = newBinaryOperation(OpGreaterEqual, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:222
		{
			yyVAL.fieldExpression = newBinaryOperation(OpRegex, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:223
		{
			yyVAL.fieldExpression = newBinaryOperation(OpNotRegex, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:224
		{
			yyVAL.fieldExpression = newBinaryOperation(OpPower, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:225
		{
			yyVAL.fieldExpression = newBinaryOperation(OpAnd, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:226
		{
			yyVAL.fieldExpression = newBinaryOperation(OpOr, yyDollar[1].fieldExpression, yyDollar[3].fieldExpression)
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
//line pkg/traceql/expr.y:227
		{
			yyVAL.fieldExpression = newUnaryOperation(OpSub, yyDollar[2].fieldExpression)
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
//line pkg/traceql/expr.y:228
		{
			yyVAL.fieldExpression = newUnaryOperation(OpNot, yyDollar[2].fieldExpression)
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:229
		{
			yyVAL.fieldExpression = yyDollar[1].static
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:230
		{
			yyVAL.fieldExpression = yyDollar[1].intrinsicField
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:231
		{
			yyVAL.fieldExpression = yyDollar[1].attributeField
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:238
		{
			yyVAL.static = NewStaticString(yyDollar[1].staticStr)
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:239
		{
			yyVAL.static = NewStaticInt(yyDollar[1].staticInt)
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:240
		{
			yyVAL.static = NewStaticFloat(yyDollar[1].staticFloat)
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:241
		{
			yyVAL.static = NewStaticBool(true)
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:242
		{
			yyVAL.static = NewStaticBool(false)
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:243
		{
			yyVAL.static = NewStaticNil()
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:244
		{
			yyVAL.static = NewStaticDuration(yyDollar[1].staticDuration)
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:245
		{
			yyVAL.static = NewStaticStatus(StatusOk)
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:246
		{
			yyVAL.static = NewStaticStatus(StatusError)
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:247
		{
			yyVAL.static = NewStaticStatus(StatusUnset)
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:251
		{
			yyVAL.intrinsicField = NewIntrinsic(IntrinsicDuration)
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:252
		{
			yyVAL.intrinsicField = NewIntrinsic(IntrinsicChildCount)
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:253
		{
			yyVAL.intrinsicField = NewIntrinsic(IntrinsicName)
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:254
		{
			yyVAL.intrinsicField = NewIntrinsic(IntrinsicStatus)
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
//line pkg/traceql/expr.y:255
		{
			yyVAL.intrinsicField = NewIntrinsic(IntrinsicParent)
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:259
		{
			yyVAL.attributeField = NewAttribute(yyDollar[2].staticStr)
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:260
		{
			yyVAL.attributeField = NewScopedAttribute(AttributeScopeResource, false, yyDollar[2].staticStr)
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:261
		{
			yyVAL.attributeField = NewScopedAttribute(AttributeScopeSpan, false, yyDollar[2].staticStr)
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
//line pkg/traceql/expr.y:262
		{
			yyVAL.attributeField = NewScopedAttribute(AttributeScopeNone, true, yyDollar[2].staticStr)
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
//line pkg/traceql/expr.y:263
		{
			yyVAL.attributeField = NewScopedAttribute(AttributeScopeResource, true, yyDollar[3].staticStr)
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
//line pkg/traceql/expr.y:264
		{
			yyVAL.attributeField = NewScopedAttribute(AttributeScopeSpan, true, yyDollar[3].staticStr)
		}
	}
	goto yystack /* stack new state and value */
}
