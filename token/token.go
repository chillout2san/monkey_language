package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// トークンや文字が未知であること
	ILLEGAL TokenType = "ILLEGAL"

	// 終端記号
	EOF TokenType = "EOF"

	// 識別子
	IDENT TokenType = "IDENT"

	// 数値型
	INT TokenType = "INT"

	// 演算子
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"
	BANG     TokenType = "!"
	LT       TokenType = "<"
	GT       TokenType = ">"

	// 区切り文字
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	L_PAREN   TokenType = "("
	R_PAREN   TokenType = ")"
	L_BRACE   TokenType = "{"
	R_BRACE   TokenType = "}"

	// キーワード
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// キーワードか識別子かを判定して返却する
func LookupIdent(ident string) TokenType {
	if keyword, ok := keywords[ident]; ok {
		return keyword
	}
	return IDENT
}
