package lexer

import (
	"monkey/token"
)

// TODO: Lexerは現状1byte文字しか対応できないためrune型を保持するように修正

// 字句解析器
type Lexer struct {
	input    string // 解析対象の文字列
	position int    // 入力における現在の位置
	char     byte   // 現在読み込んでいる文字
}

// 解析対象の文字列を一文字分読み込む
func (l *Lexer) readChar() {
	var nextPosition = l.position + 1
	if nextPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[nextPosition]
	}
	l.position = nextPosition
}

// 現在読み込んでいる文字の次の文字を返却する
// positionの位置は変更しない
func (l *Lexer) peekChar() byte {
	var nextPosition = l.position + 1
	if nextPosition >= len(l.input) {
		return 0
	}
	return l.input[nextPosition]
}

// 現在の文字のトークンを識別して返却する
func (l *Lexer) GetToken() token.Token {
	var tokenType token.TokenType
	var literal string

	// 現在の文字がホワイトスペースだった場合は読み飛ばす
	for l.char == ' ' || l.char == '\n' || l.char == '\r' || l.char == '\t' {
		l.readChar()
	}

	switch l.char {
	case '=':
		// イコールの演算子かどうか
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			tokenType = token.EQ
			literal := string(char) + string(l.char)
			l.readChar()
			return token.Token{Type: token.EQ, Literal: literal}
		}
		// 代入の演算子だった場合こちらへ
		tokenType = token.ASSIGN
	case '+':
		tokenType = token.PLUS
	case '-':
		tokenType = token.MINUS
	case '*':
		tokenType = token.ASTERISK
	case '/':
		tokenType = token.SLASH
	case '!':
		// notイコールの演算子かどうか
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			literal := string(char) + string(l.char)
			l.readChar()
			return token.Token{Type: token.NOT_EQ, Literal: literal}
		}
		tokenType = token.BANG
	case '<':
		tokenType = token.LT
	case '>':
		tokenType = token.GT
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
	default:
		if l.isLetter() {
			literal := l.readIdentifier()
			return token.Token{Type: token.LookupIdent(literal), Literal: literal}
		} else if l.isDigit() {
			literal := l.readNumber()
			return token.Token{Type: token.INT, Literal: literal}
		} else {
			return token.Token{Type: token.ILLEGAL, Literal: string(l.char)}
		}
	}

	if tokenType == token.EOF {
		literal = ""
	} else {
		literal = string(l.char)
	}

	l.readChar()
	return token.Token{Type: tokenType, Literal: literal}
}

// 現在読み込んでいる文字がキーワードもしくは識別子の一部かを判別するため、
// アルファベットかアンダースコアがどうか判定する
func (l *Lexer) isLetter() bool {
	return ('a' <= l.char && l.char <= 'z') || ('A' <= l.char && l.char <= 'Z') || l.char == '_'
}

// 現在読み込んでいる位置からキーワードもしくは識別子の塊一つを抜き出す
func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.isLetter() {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 現在読み込んでいる位置から数値の塊一つを抜き出す
func (l *Lexer) readNumber() string {
	position := l.position
	for l.isDigit() {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 現在読み込んでいる文字が数値かどうか
func (l *Lexer) isDigit() bool {
	return '0' <= l.char && l.char <= '9'
}

// 字句解析器のインスタンスへのポインタを返却する
// 一文字読み込んだ状態でスタートする
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input, position: 0, char: input[0]}
	return l
}
