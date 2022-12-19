package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// トークンや文字が未知であること
	ILLEGAL = "ILLEGAL"

	// 終端記号
	EOF = "EOF"

	// 識別子
	IDENT = "IDENT"

	// 数値型
	INT = "INT"

	// 代入の演算子
	ASSIGN = "="

	// 足し算の演算子
	PLUS = "+"

	// 区切り文字
	COMMA     = ","
	SEMICOLON = ";"
	L_PAREN   = "("
	R_PAREN   = ")"
	L_BRACE   = "{"
	R_BRACE   = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
