package ast

import "monkey/token"

// ノード
type Node interface {
	TokenLiteral() string // ノードに関連付けられているトークンのリテラルを返す
}

// 識別子のノード
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
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

// letのノード
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
