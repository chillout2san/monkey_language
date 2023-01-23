package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

// 構造解析器
type Parser struct {
	lexer     *lexer.Lexer // 字句解析器
	curToken  token.Token  // 現在のトークン
	peekToken token.Token  // 次のトークン
	errors    []string
}

// 構造解析時のエラーを出力する
func (p *Parser) Errors() []string {
	return p.errors
}

// 構造解析時のエラーを作成する
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// 次のトークンをセットする
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.GetToken()
}

// 文を次々解析することでプログラムを解析する
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Statements: []ast.Statement{}}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// 文を解析する
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// let文を解析する
func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// 識別子が来るか
	if !p.checkPeekAndNext(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// 代入演算子が来るか
	if !p.checkPeekAndNext(token.ASSIGN) {
		return nil
	}

	// TODO: セミコロンに遭遇するまで式を読み飛ばしている
	for p.peekToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

// 次のトークンが引数に渡したトークンタイプと一致しているかどうか検査する
// 一致した場合はトークンを次に読み進める
func (p *Parser) checkPeekAndNext(tt token.TokenType) bool {
	if p.peekToken.Type == tt {
		p.nextToken()
		return true
	}
	p.peekError(tt)
	return false
}

// 構造解析器を新規に生成する
func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}
