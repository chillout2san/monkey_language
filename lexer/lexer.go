package lexer

import (
	"monkey/token"
)

// 字句解析器
type Lexer struct {
	input        string // 解析対象の文字
	position     int    // 入力における現在の位置
	readPosition int    // これから読み込む位置
	ch           byte   // 現在検査中の文字
}

func (l *Lexer) readChar() {
	var isEnd = l.readPosition >= len(l.input)
	if isEnd {
		l.ch = 0
	}
}

func (l *Lexer) NextToken() token.Token {
	return token.Token{}
}

// 字句解析器のインスタンスへのポインタを返却する
func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}
