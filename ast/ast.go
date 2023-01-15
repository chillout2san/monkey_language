package ast

// ノード
type Node interface {
	TokenLiteral() string
}

// 文のノード
type Statement interface {
	Node
	statementNode()
}

// 式のノード
type Expression interface {
	Node
	expressionNode()
}

// ルートノード
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
