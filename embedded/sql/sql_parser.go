// Code generated by goyacc -l -o sql_parser.go sql_grammar.y. DO NOT EDIT.
package sql

import __yyfmt__ "fmt"

import "fmt"

func setResult(l yyLexer, stmts []SQLStmt) {
	l.(*lexer).result = stmts
}

type yySymType struct {
	yys           int
	stmts         []SQLStmt
	stmt          SQLStmt
	datasource    DataSource
	colsSpec      []*ColSpec
	colSpec       *ColSpec
	cols          []*ColSelector
	rows          []*RowSpec
	row           *RowSpec
	values        []ValueExp
	value         ValueExp
	id            string
	number        uint64
	str           string
	boolean       bool
	blob          []byte
	sqlType       SQLValueType
	aggFn         AggregateFn
	ids           []string
	col           *ColSelector
	sel           Selector
	sels          []Selector
	distinct      bool
	ds            DataSource
	tableRef      *tableRef
	period        period
	openPeriod    *openPeriod
	periodInstant periodInstant
	joins         []*JoinSpec
	join          *JoinSpec
	joinType      JoinType
	exp           ValueExp
	binExp        ValueExp
	err           error
	ordcols       []*OrdCol
	opt_ord       bool
	logicOp       LogicOperator
	cmpOp         CmpOperator
	pparam        int
	update        *colUpdate
	updates       []*colUpdate
	onConflict    *OnConflictDo
}

const CREATE = 57346
const USE = 57347
const DATABASE = 57348
const SNAPSHOT = 57349
const SINCE = 57350
const AFTER = 57351
const BEFORE = 57352
const UNTIL = 57353
const TX = 57354
const OF = 57355
const TIMESTAMP = 57356
const TABLE = 57357
const UNIQUE = 57358
const INDEX = 57359
const ON = 57360
const ALTER = 57361
const ADD = 57362
const COLUMN = 57363
const PRIMARY = 57364
const KEY = 57365
const BEGIN = 57366
const TRANSACTION = 57367
const COMMIT = 57368
const ROLLBACK = 57369
const INSERT = 57370
const UPSERT = 57371
const INTO = 57372
const VALUES = 57373
const DELETE = 57374
const UPDATE = 57375
const SET = 57376
const CONFLICT = 57377
const DO = 57378
const NOTHING = 57379
const SELECT = 57380
const DISTINCT = 57381
const FROM = 57382
const JOIN = 57383
const HAVING = 57384
const WHERE = 57385
const GROUP = 57386
const BY = 57387
const LIMIT = 57388
const ORDER = 57389
const ASC = 57390
const DESC = 57391
const AS = 57392
const UNION = 57393
const ALL = 57394
const NOT = 57395
const LIKE = 57396
const IF = 57397
const EXISTS = 57398
const IN = 57399
const IS = 57400
const AUTO_INCREMENT = 57401
const NULL = 57402
const CAST = 57403
const DATABASES = 57404
const TABLES = 57405
const COLUMNS = 57406
const INDEXES = 57407
const NPARAM = 57408
const PPARAM = 57409
const JOINTYPE = 57410
const LOP = 57411
const CMPOP = 57412
const IDENTIFIER = 57413
const TYPE = 57414
const NUMBER = 57415
const VARCHAR = 57416
const BOOLEAN = 57417
const BLOB = 57418
const AGGREGATE_FUNC = 57419
const ERROR = 57420
const STMT_SEPARATOR = 57421

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"CREATE",
	"USE",
	"DATABASE",
	"SNAPSHOT",
	"SINCE",
	"AFTER",
	"BEFORE",
	"UNTIL",
	"TX",
	"OF",
	"TIMESTAMP",
	"TABLE",
	"UNIQUE",
	"INDEX",
	"ON",
	"ALTER",
	"ADD",
	"COLUMN",
	"PRIMARY",
	"KEY",
	"BEGIN",
	"TRANSACTION",
	"COMMIT",
	"ROLLBACK",
	"INSERT",
	"UPSERT",
	"INTO",
	"VALUES",
	"DELETE",
	"UPDATE",
	"SET",
	"CONFLICT",
	"DO",
	"NOTHING",
	"SELECT",
	"DISTINCT",
	"FROM",
	"JOIN",
	"HAVING",
	"WHERE",
	"GROUP",
	"BY",
	"LIMIT",
	"ORDER",
	"ASC",
	"DESC",
	"AS",
	"UNION",
	"ALL",
	"NOT",
	"LIKE",
	"IF",
	"EXISTS",
	"IN",
	"IS",
	"AUTO_INCREMENT",
	"NULL",
	"CAST",
	"DATABASES",
	"TABLES",
	"COLUMNS",
	"INDEXES",
	"NPARAM",
	"PPARAM",
	"JOINTYPE",
	"LOP",
	"CMPOP",
	"IDENTIFIER",
	"TYPE",
	"NUMBER",
	"VARCHAR",
	"BOOLEAN",
	"BLOB",
	"AGGREGATE_FUNC",
	"ERROR",
	"','",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'.'",
	"STMT_SEPARATOR",
	"'('",
	"')'",
	"'['",
	"']'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 72,
	54, 136,
	57, 136,
	-2, 125,
	-1, 180,
	41, 103,
	-2, 98,
	-1, 207,
	41, 103,
	-2, 100,
}

const yyPrivate = 57344

const yyLast = 374

var yyAct = [...]int{
	71, 295, 59, 174, 133, 228, 101, 139, 231, 130,
	224, 227, 206, 6, 140, 93, 152, 44, 96, 77,
	268, 172, 249, 220, 172, 18, 278, 119, 272, 273,
	250, 32, 247, 142, 143, 144, 145, 251, 117, 118,
	74, 248, 33, 76, 49, 50, 51, 88, 84, 113,
	114, 116, 115, 86, 87, 58, 167, 141, 85, 241,
	80, 81, 82, 83, 60, 119, 198, 98, 75, 232,
	172, 112, 240, 79, 119, 122, 123, 118, 221, 239,
	125, 119, 212, 172, 233, 117, 118, 113, 114, 116,
	115, 173, 117, 118, 229, 135, 113, 114, 116, 115,
	211, 197, 132, 113, 114, 116, 115, 149, 119, 189,
	196, 136, 188, 70, 156, 157, 158, 159, 160, 161,
	146, 171, 169, 105, 193, 127, 186, 168, 185, 184,
	113, 114, 116, 115, 183, 154, 105, 126, 166, 179,
	124, 177, 106, 104, 180, 170, 92, 91, 20, 294,
	254, 289, 61, 187, 74, 182, 178, 76, 181, 94,
	119, 88, 84, 148, 192, 195, 249, 86, 87, 253,
	190, 172, 85, 100, 80, 81, 82, 83, 60, 246,
	61, 204, 75, 61, 116, 115, 60, 79, 210, 60,
	202, 56, 226, 103, 191, 137, 217, 225, 61, 131,
	215, 216, 33, 222, 214, 200, 97, 218, 153, 155,
	223, 253, 150, 230, 102, 147, 108, 62, 237, 238,
	235, 234, 48, 43, 38, 138, 69, 209, 245, 163,
	225, 267, 194, 107, 40, 244, 162, 266, 119, 164,
	256, 121, 165, 257, 260, 153, 261, 262, 263, 74,
	269, 264, 76, 63, 54, 34, 88, 84, 281, 270,
	296, 297, 86, 87, 175, 277, 288, 85, 276, 80,
	81, 82, 83, 60, 89, 259, 282, 75, 94, 284,
	39, 275, 79, 236, 287, 99, 290, 10, 11, 31,
	36, 292, 293, 18, 110, 111, 286, 298, 279, 271,
	299, 52, 12, 201, 199, 41, 30, 7, 29, 8,
	9, 13, 14, 21, 242, 15, 16, 128, 2, 90,
	22, 18, 285, 65, 203, 109, 64, 176, 42, 23,
	25, 24, 28, 68, 67, 46, 47, 26, 27, 37,
	134, 19, 252, 95, 120, 243, 265, 280, 291, 213,
	219, 258, 73, 72, 274, 208, 207, 205, 66, 45,
	53, 35, 57, 55, 78, 255, 283, 129, 151, 17,
	5, 4, 3, 1,
}

var yyPact = [...]int{
	283, -1000, -1000, 63, -1000, -1000, -1000, 288, -1000, -1000,
	314, 331, 317, 278, 276, 249, 131, 204, 251, -1000,
	283, -1000, 153, 179, 179, 311, 152, 327, 151, 131,
	131, 131, 267, -1000, 202, 109, -1000, -1000, -1000, 146,
	200, 308, 179, -1000, -1000, 323, 101, 101, 299, 61,
	60, 235, 135, 255, -1000, 245, -1000, 94, 143, -1000,
	57, 52, 56, 177, 145, 307, -1000, 101, 101, -1000,
	196, 23, 188, -1000, 196, 196, 54, -1000, -1000, 196,
	-1000, -1000, -1000, -1000, 51, 39, -1000, -1000, -1000, -1000,
	296, 128, 128, 335, 196, 116, -1000, 155, -1000, -29,
	112, -1000, -1000, 144, 81, 141, 137, -1000, 49, 138,
	-1000, -1000, 23, 196, 196, 196, 196, 196, 196, 176,
	185, -1000, 7, 102, 255, -31, 196, 35, 137, 34,
	92, -1000, 4, 218, 310, 23, 335, 135, 196, 335,
	327, 255, 48, 43, 42, 40, 143, -1000, 25, 22,
	-1000, 91, -1000, 122, 128, 38, 102, 102, 180, 180,
	7, 50, -1000, 172, 196, 24, 14, -1000, 16, -1000,
	-1000, 273, 134, 272, -1000, 117, 306, 218, -1000, 23,
	159, 143, 13, -5, 133, 131, 131, -1000, -1000, -1000,
	174, -65, -9, 128, -1000, 7, -13, -1000, 120, 8,
	-1000, 8, -1000, -2, -1000, 235, -1000, 159, 242, -1000,
	-1000, 143, 143, -8, -1000, -15, -28, 291, -1000, 175,
	106, -1000, -55, -46, -57, 23, -50, 132, -1000, 196,
	90, -1000, -1000, 128, 231, -1000, -29, -1000, -1000, 143,
	143, 143, -2, 178, -1000, 171, -69, -1000, -1000, 196,
	-1000, -1000, -1000, 8, 264, -59, 87, -58, 239, 223,
	335, -1000, -1000, -1000, -61, -1000, -1000, -1000, -1000, 23,
	-1000, 262, -1000, -1000, 211, 196, 127, 304, -1000, 259,
	218, 221, 23, 72, -1000, 196, -1000, -1000, 127, 127,
	23, 70, 212, -1000, 127, -1000, -1000, -1000, 212, -1000,
}

var yyPgo = [...]int{
	0, 373, 318, 372, 371, 370, 13, 369, 368, 16,
	9, 8, 367, 366, 11, 5, 10, 365, 364, 19,
	363, 362, 2, 361, 360, 7, 14, 17, 359, 358,
	226, 357, 12, 356, 355, 0, 15, 354, 353, 352,
	351, 3, 350, 6, 349, 348, 347, 1, 4, 280,
	346, 345, 344, 18, 343, 342, 341,
}

var yyR1 = [...]int{
	0, 1, 2, 2, 56, 56, 3, 3, 3, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 49,
	49, 11, 11, 5, 5, 5, 5, 55, 55, 54,
	54, 53, 12, 12, 14, 14, 15, 10, 10, 13,
	13, 17, 17, 16, 16, 18, 18, 18, 18, 18,
	18, 18, 18, 18, 8, 8, 9, 42, 42, 50,
	50, 51, 51, 51, 6, 6, 7, 24, 24, 23,
	23, 20, 20, 21, 21, 19, 19, 19, 22, 22,
	25, 25, 25, 25, 25, 25, 44, 44, 26, 27,
	28, 28, 28, 29, 29, 29, 30, 30, 31, 31,
	32, 32, 33, 34, 34, 36, 36, 40, 40, 37,
	37, 41, 41, 46, 46, 48, 48, 45, 45, 47,
	47, 47, 43, 43, 43, 35, 35, 35, 35, 35,
	35, 35, 35, 38, 38, 38, 52, 52, 39, 39,
	39, 39, 39, 39, 39, 39,
}

var yyR2 = [...]int{
	0, 1, 2, 3, 0, 1, 1, 1, 1, 2,
	1, 1, 3, 3, 3, 11, 8, 9, 6, 0,
	3, 1, 3, 9, 8, 6, 7, 0, 4, 1,
	3, 3, 0, 1, 1, 3, 3, 1, 3, 1,
	3, 0, 1, 1, 3, 1, 1, 1, 1, 6,
	3, 1, 1, 1, 1, 3, 5, 0, 3, 0,
	1, 0, 1, 2, 1, 4, 12, 0, 1, 0,
	1, 1, 1, 2, 4, 1, 4, 4, 1, 3,
	3, 4, 4, 5, 5, 5, 0, 1, 1, 2,
	0, 2, 2, 0, 2, 2, 2, 1, 0, 1,
	1, 2, 6, 0, 1, 0, 2, 0, 3, 0,
	2, 0, 2, 0, 3, 0, 4, 2, 4, 0,
	1, 1, 0, 1, 2, 1, 1, 2, 2, 4,
	4, 6, 6, 1, 1, 3, 0, 1, 3, 3,
	3, 3, 3, 3, 3, 4,
}

var yyChk = [...]int{
	-1000, -1, -2, -3, -4, -5, -6, 24, 26, 27,
	4, 5, 19, 28, 29, 32, 33, -7, 38, -56,
	85, 25, 6, 15, 17, 16, 6, 7, 15, 30,
	30, 40, -26, 71, 51, -23, 39, -2, 71, -49,
	55, -49, 17, 71, -27, -28, 8, 9, 71, -26,
	-26, -26, 34, -24, 52, -20, 82, -21, -19, -22,
	77, 71, 71, 53, 18, -49, -29, 11, 10, -30,
	12, -35, -38, -39, 53, 81, 56, -19, -18, 86,
	73, 74, 75, 76, 61, 71, 66, 67, 60, -30,
	20, 86, 86, -36, 43, -54, -53, 71, -6, 40,
	79, -43, 71, 50, 86, 84, 86, 56, 71, 18,
	-30, -30, -35, 80, 81, 83, 82, 69, 70, 58,
	-52, 53, -35, -35, 86, -35, 86, 86, 21, -12,
	-10, 71, -10, -48, 5, -35, -36, 79, 70, -25,
	-26, 86, 62, 63, 64, 65, -19, 71, 82, -22,
	71, -8, -9, 71, 86, 71, -35, -35, -35, -35,
	-35, -35, 60, 53, 54, 57, -6, 87, -35, 87,
	-9, 87, 79, 87, -41, 46, 17, -48, -53, -35,
	-48, -27, -6, 86, 86, 86, 86, -43, 87, 87,
	79, 72, -10, 86, 60, -35, 86, 87, 50, 31,
	71, 31, 73, 18, -41, -31, -32, -33, -34, 68,
	-43, 87, 87, -44, 71, -26, -26, 22, -9, -42,
	88, 87, -10, -6, -16, -35, 72, -14, -15, 86,
	-14, -11, 71, 86, -36, -32, 41, -43, -43, 87,
	87, 87, 23, -51, 60, 53, 73, 87, 87, 79,
	87, 87, -55, 79, 18, -17, -16, -10, -40, 44,
	-25, -43, -43, -43, -11, -50, 59, 60, 89, -35,
	-15, 35, 87, 87, -37, 42, 45, -48, 87, 36,
	-46, 47, -35, -13, -22, 18, 37, -41, 45, 79,
	-35, -45, -22, -22, 79, -47, 48, 49, -22, -47,
}

var yyDef = [...]int{
	0, -2, 1, 4, 6, 7, 8, 0, 10, 11,
	0, 0, 0, 0, 0, 0, 0, 64, 69, 2,
	5, 9, 0, 19, 19, 0, 0, 90, 0, 0,
	0, 0, 0, 88, 67, 0, 70, 3, 12, 0,
	0, 0, 19, 13, 14, 93, 0, 0, 0, 0,
	0, 105, 0, 0, 68, 0, 71, 72, 122, 75,
	0, 78, 0, 0, 0, 0, 89, 0, 0, 91,
	0, 97, -2, 126, 0, 0, 0, 133, 134, 0,
	45, 46, 47, 48, 0, 78, 51, 52, 53, 92,
	0, 32, 0, 115, 0, 105, 29, 0, 65, 0,
	0, 73, 123, 0, 0, 0, 0, 20, 0, 0,
	94, 95, 96, 0, 0, 0, 0, 0, 0, 0,
	0, 137, 127, 128, 0, 0, 0, 0, 0, 0,
	33, 37, 0, 111, 0, 106, 115, 0, 0, 115,
	90, 0, 0, 0, 0, 0, 122, 124, 0, 0,
	79, 0, 54, 0, 0, 0, 138, 139, 140, 141,
	142, 143, 144, 0, 0, 0, 0, 135, 0, 50,
	18, 0, 0, 0, 25, 0, 0, 111, 30, 31,
	-2, 122, 0, 0, 86, 0, 0, 74, 76, 77,
	0, 57, 0, 0, 145, 129, 0, 130, 0, 0,
	38, 0, 112, 0, 26, 105, 99, -2, 0, 104,
	80, 122, 122, 0, 87, 0, 0, 0, 55, 61,
	0, 16, 0, 0, 0, 43, 0, 27, 34, 41,
	24, 116, 21, 0, 107, 101, 0, 81, 82, 122,
	122, 122, 0, 59, 62, 0, 0, 17, 131, 0,
	132, 49, 23, 0, 0, 0, 42, 0, 109, 0,
	115, 83, 84, 85, 0, 56, 60, 63, 58, 44,
	35, 0, 36, 22, 113, 0, 0, 0, 15, 0,
	111, 0, 110, 108, 39, 0, 28, 66, 0, 0,
	102, 114, 119, 40, 0, 117, 120, 121, 119, 118,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	86, 87, 82, 80, 79, 81, 84, 83, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 88, 3, 89,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 85,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

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
		{
			yyVAL.stmts = yyDollar[1].stmts
			setResult(yylex, yyDollar[1].stmts)
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []SQLStmt{yyDollar[1].stmt}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmts = append([]SQLStmt{yyDollar[1].stmt}, yyDollar[3].stmts...)
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &BeginTransactionStmt{}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &CommitStmt{}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &RollbackStmt{}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &CreateDatabaseStmt{DB: yyDollar[3].id}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &UseDatabaseStmt{DB: yyDollar[3].id}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &UseSnapshotStmt{period: yyDollar[3].period}
		}
	case 15:
		yyDollar = yyS[yypt-11 : yypt+1]
		{
			yyVAL.stmt = &CreateTableStmt{ifNotExists: yyDollar[3].boolean, table: yyDollar[4].id, colsSpec: yyDollar[6].colsSpec, pkColNames: yyDollar[10].ids}
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &CreateIndexStmt{ifNotExists: yyDollar[3].boolean, table: yyDollar[5].id, cols: yyDollar[7].ids}
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &CreateIndexStmt{unique: true, ifNotExists: yyDollar[4].boolean, table: yyDollar[6].id, cols: yyDollar[8].ids}
		}
	case 18:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.stmt = &AddColumnStmt{table: yyDollar[3].id, colSpec: yyDollar[6].colSpec}
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = []string{yyDollar[1].id}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ids = yyDollar[2].ids
		}
	case 23:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &UpsertIntoStmt{isInsert: true, tableRef: yyDollar[3].tableRef, cols: yyDollar[5].ids, rows: yyDollar[8].rows, onConflict: yyDollar[9].onConflict}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &UpsertIntoStmt{tableRef: yyDollar[3].tableRef, cols: yyDollar[5].ids, rows: yyDollar[8].rows}
		}
	case 25:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.stmt = &DeleteFromStmt{tableRef: yyDollar[3].tableRef, where: yyDollar[4].exp, indexOn: yyDollar[5].ids, limit: int(yyDollar[6].number)}
		}
	case 26:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.stmt = &UpdateStmt{tableRef: yyDollar[2].tableRef, updates: yyDollar[4].updates, where: yyDollar[5].exp, indexOn: yyDollar[6].ids, limit: int(yyDollar[7].number)}
		}
	case 27:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.onConflict = nil
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.onConflict = &OnConflictDo{}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.updates = []*colUpdate{yyDollar[1].update}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.updates = append(yyDollar[1].updates, yyDollar[3].update)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.update = &colUpdate{col: yyDollar[1].id, op: yyDollar[2].cmpOp, val: yyDollar[3].exp}
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ids = nil
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = yyDollar[1].ids
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.rows = []*RowSpec{yyDollar[1].row}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.rows = append(yyDollar[1].rows, yyDollar[3].row)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.row = &RowSpec{Values: yyDollar[2].values}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = []string{yyDollar[1].id}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ids = append(yyDollar[1].ids, yyDollar[3].id)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.cols = []*ColSelector{yyDollar[1].col}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = append(yyDollar[1].cols, yyDollar[3].col)
		}
	case 41:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.values = nil
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.values = yyDollar[1].values
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.values = []ValueExp{yyDollar[1].exp}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].exp)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Number{val: int64(yyDollar[1].number)}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Varchar{val: yyDollar[1].str}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Bool{val: yyDollar[1].boolean}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Blob{val: yyDollar[1].blob}
		}
	case 49:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.value = &Cast{val: yyDollar[3].exp, t: yyDollar[5].sqlType}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.value = &SysFn{fn: yyDollar[1].id}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Param{id: yyDollar[1].id}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Param{id: fmt.Sprintf("param%d", yyDollar[1].pparam), pos: yyDollar[1].pparam}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &NullValue{t: AnyType}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.colsSpec = []*ColSpec{yyDollar[1].colSpec}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.colsSpec = append(yyDollar[1].colsSpec, yyDollar[3].colSpec)
		}
	case 56:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.colSpec = &ColSpec{colName: yyDollar[1].id, colType: yyDollar[2].sqlType, maxLen: int(yyDollar[3].number), notNull: yyDollar[4].boolean, autoIncrement: yyDollar[5].boolean}
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[2].number
		}
	case 59:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 61:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &UnionStmt{
				distinct: yyDollar[3].distinct,
				left:     yyDollar[1].stmt.(DataSource),
				right:    yyDollar[4].stmt.(DataSource),
			}
		}
	case 66:
		yyDollar = yyS[yypt-12 : yypt+1]
		{
			yyVAL.stmt = &SelectStmt{
				distinct:  yyDollar[2].distinct,
				selectors: yyDollar[3].sels,
				ds:        yyDollar[5].ds,
				indexOn:   yyDollar[6].ids,
				joins:     yyDollar[7].joins,
				where:     yyDollar[8].exp,
				groupBy:   yyDollar[9].cols,
				having:    yyDollar[10].exp,
				orderBy:   yyDollar[11].ordcols,
				limit:     int(yyDollar[12].number),
			}
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.distinct = true
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.distinct = false
		}
	case 69:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.distinct = false
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.distinct = true
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sels = nil
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sels = yyDollar[1].sels
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyDollar[1].sel.setAlias(yyDollar[2].id)
			yyVAL.sels = []Selector{yyDollar[1].sel}
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[3].sel.setAlias(yyDollar[4].id)
			yyVAL.sels = append(yyDollar[1].sels, yyDollar[3].sel)
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sel = yyDollar[1].col
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, col: "*"}
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, db: yyDollar[3].col.db, table: yyDollar[3].col.table, col: yyDollar[3].col.col}
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.col = &ColSelector{col: yyDollar[1].id}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.col = &ColSelector{table: yyDollar[1].id, col: yyDollar[3].id}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].tableRef.period = yyDollar[2].period
			yyDollar[1].tableRef.as = yyDollar[3].id
			yyVAL.ds = yyDollar[1].tableRef
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[2].stmt.(*SelectStmt).as = yyDollar[4].id
			yyVAL.ds = yyDollar[2].stmt.(DataSource)
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.ds = &ListDatabasesStmt{as: yyDollar[4].id}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.ds = &ListTablesStmt{db: yyDollar[3].id, as: yyDollar[5].id}
		}
	case 84:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.ds = &ListColumnsStmt{tableRef: yyDollar[3].tableRef, as: yyDollar[5].id}
		}
	case 85:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.ds = &ListIndexesStmt{tableRef: yyDollar[3].tableRef, as: yyDollar[5].id}
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.id = ""
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.id = yyDollar[1].id
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.tableRef = &tableRef{table: yyDollar[1].id}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.period = period{start: yyDollar[1].openPeriod, end: yyDollar[2].openPeriod}
		}
	case 90:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.openPeriod = nil
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.openPeriod = &openPeriod{inclusive: true, instant: yyDollar[2].periodInstant}
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.openPeriod = &openPeriod{instant: yyDollar[2].periodInstant}
		}
	case 93:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.openPeriod = nil
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.openPeriod = &openPeriod{inclusive: true, instant: yyDollar[2].periodInstant}
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.openPeriod = &openPeriod{instant: yyDollar[2].periodInstant}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.periodInstant = periodInstant{instantType: txInstant, exp: yyDollar[2].exp}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.periodInstant = periodInstant{instantType: timeInstant, exp: yyDollar[1].exp}
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.joins = nil
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = yyDollar[1].joins
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = []*JoinSpec{yyDollar[1].join}
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.joins = append([]*JoinSpec{yyDollar[1].join}, yyDollar[2].joins...)
		}
	case 102:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.join = &JoinSpec{joinType: yyDollar[1].joinType, ds: yyDollar[3].ds, indexOn: yyDollar[4].ids, cond: yyDollar[6].exp}
		}
	case 103:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.joinType = InnerJoin
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joinType = yyDollar[1].joinType
		}
	case 105:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exp = nil
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 107:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.cols = nil
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = yyDollar[3].cols
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exp = nil
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 111:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.number = yyDollar[2].number
		}
	case 113:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ordcols = nil
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ordcols = yyDollar[3].ordcols
		}
	case 115:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ids = nil
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.ids = yyDollar[4].ids
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.ordcols = []*OrdCol{{sel: yyDollar[1].col, descOrder: yyDollar[2].opt_ord}}
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.ordcols = append(yyDollar[1].ordcols, &OrdCol{sel: yyDollar[3].col, descOrder: yyDollar[4].opt_ord})
		}
	case 119:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.opt_ord = false
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = false
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = true
		}
	case 122:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.id = ""
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.id = yyDollar[1].id
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.id = yyDollar[2].id
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].exp
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].binExp
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = &NotBoolExp{exp: yyDollar[2].exp}
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = &NumExp{left: &Number{val: 0}, op: SUBSOP, right: yyDollar[2].exp}
		}
	case 129:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exp = &LikeBoolExp{val: yyDollar[1].exp, notLike: yyDollar[2].boolean, pattern: yyDollar[4].exp}
		}
	case 130:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exp = &ExistsBoolExp{q: (yyDollar[3].stmt).(*SelectStmt)}
		}
	case 131:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.exp = &InSubQueryExp{val: yyDollar[1].exp, notIn: yyDollar[2].boolean, q: yyDollar[5].stmt.(*SelectStmt)}
		}
	case 132:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.exp = &InListExp{val: yyDollar[1].exp, notIn: yyDollar[2].boolean, values: yyDollar[5].values}
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].sel
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].value
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 136:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: ADDOP, right: yyDollar[3].exp}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: SUBSOP, right: yyDollar[3].exp}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: DIVOP, right: yyDollar[3].exp}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: MULTOP, right: yyDollar[3].exp}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &BinBoolExp{left: yyDollar[1].exp, op: yyDollar[2].logicOp, right: yyDollar[3].exp}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &CmpBoolExp{left: yyDollar[1].exp, op: yyDollar[2].cmpOp, right: yyDollar[3].exp}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &CmpBoolExp{left: yyDollar[1].exp, op: EQ, right: &NullValue{t: AnyType}}
		}
	case 145:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.binExp = &CmpBoolExp{left: yyDollar[1].exp, op: NE, right: &NullValue{t: AnyType}}
		}
	}
	goto yystack /* stack new state and value */
}
