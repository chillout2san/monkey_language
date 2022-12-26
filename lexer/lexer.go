package lexer

import (
	"monkey/token"
)

// TODO: Lexerは現状1byte文字しか対応できないためrune型を保持するように修正

// 字句解析器
type Lexer struct {
	input        string // 解析対象の文字列
	position     int    // 入力における現在の位置
	nextPosition int    // これから読み込む位置
	char         byte   // 現在読み込み終わった文字
}

// 解析対象の文字列を一文字分読み込む
func (l *Lexer) ReadChar() {
	var isEnd = l.nextPosition >= len(l.input)
	if isEnd {
		l.char = 0
	} else {
		l.char = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

// 現在の文字のトークンを識別して返却する
func (l *Lexer) GetToken() token.Token {
	var tokenType string
	var literal string

	switch l.char {
	case '=':
		tokenType = token.ASSIGN
	case '+':
		tokenType = token.PLUS
	case ',':
		tokenType = token.COMMA
	case ';':
		tokenType = token.SEMICOLON
	case '(':
		tokenType = token.L_PAREN
	case ')':
		tokenType = token.R_PAREN
	case '{':
		tokenType = token.L_BRACE
	case '}':
		tokenType = token.R_BRACE
	case 0:
		tokenType = token.EOF
	}

	if tokenType == token.EOF {
		literal = ""
	} else {
		literal = string(l.char)
	}

	return token.Token{Type: token.TokenType(tokenType), Literal: literal}
}

// 字句解析器のインスタンスへのポインタを返却する
// 一文字読み込んだ状態でスタートする
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar()
	return l
}
