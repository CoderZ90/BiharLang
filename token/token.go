package token

type TokenType string

const (
	// Special tokens
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literals
	IDENT  = "IDENT"  // x, naam, jor
	INT    = "INT"    // 10, 20
	STRING = "STRING" // "bihar", "hello"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	// Comparison
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	RAKHA    = "RAKHA"    // let, var
	BOL      = "BOL"      // print
	JODI     = "JODI"     // if
	NATA     = "NATA"     // else
	JABTAK   = "JABTAK"   // while
	KAAM     = "KAAM"     // function
	LAUTA    = "LAUTA"    // return
	SACH     = "SACH"     // true
	JHOOT    = "JHOOT"    // false
	KUCHNAHI = "KUCHNAHI" // null
	BAS      = "BAS"      // break
	AAGE     = "AAGE"     // continue
	AUR      = "AUR"      // and
	YA       = "YA"       // OR
	ULTA     = "ULTA"     // NOT
	JUGAAD   = "JUGAAD"   // try / attempt / recoverable operation
	PAKKA    = "PAKKA"    // success
	LAA      = "LAA"      // import
	DIKHAWA  = "DIKHAWA"  // display/show
	CHALAWA  = "CHALAWA"  // run
	BAITHARE = "BAITHARE" // enter/start
	SAAF     = "SAAF"     // format/clean
	MADAD    = "MADAD"    // help
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"rakha":    RAKHA,
	"bol":      BOL,
	"jodi":     JODI,
	"nata":     NATA,
	"jabtak":   JABTAK,
	"kaam":     KAAM,
	"lauta":    LAUTA,
	"sach":     SACH,
	"jhoot":    JHOOT,
	"kuchnahi": KUCHNAHI,
	"bas":      BAS,
	"aage":     AAGE,
	"aur":      AUR,
	"ya":       YA,
	"ulta":     ULTA,
	"jugaad":   JUGAAD,
	"pakka":    PAKKA,
	"laa":      LAA,
	"dikhawa":  DIKHAWA,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
