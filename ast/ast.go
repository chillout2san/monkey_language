package ast

import (
	"bytes"
	"monkey/token"
)

// ノード
type Node interface {
	TokenLiteral() string // ノードに関連付けられているトークンのリテラルを返す
	String() string
}

// 識別子のノード
type Identifier struct {
	Token token.Token
	Value string
}

// 式であることを宣言するための空メソッド
func (i *Identifier) expressionNode() {}

// 宣言した識別子を返却する
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// 識別子の中身の値を返却する
func (i *Identifier) String() string {
	return i.Value
}

// 式のノード
type Expression interface {
	Node
	expressionNode()
}

// 文のノード
type Statement interface {
	Node
	statementNode()
}

// ルートノード(文の集合体)
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// 式文のノード
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// 文であることを宣言する空メソッド
func (e *ExpressionStatement) statementNode() {}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}
	return ""
}

// letのノード
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

// let文を復元して文字列で出力する
func (l *LetStatement) String() string {
	var out bytes.Buffer

	// letとホワイトスペース
	out.WriteString(l.TokenLiteral() + " ")
	// 識別子の名前
	out.WriteString(l.Name.String())
	out.WriteString(" = ")

	// let文の値があれば出力
	if l.Value != nil {
		out.WriteString(l.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// returnのノード
type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (r *ReturnStatement) statementNode() {}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

// return文を復元して文字列で出力する
func (r *ReturnStatement) String() string {
	var out bytes.Buffer

	// returnとホワイトスペース
	out.WriteString(r.TokenLiteral() + " ")

	// return文の値があれば出力
	if r.Value != nil {
		out.WriteString(r.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
