package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	if program == nil {
		t.Error("ParseProgram returned nil")
		return
	}
	if len(p.errors) > 0 {
		for _, err := range p.errors {
			t.Errorf("parser error: %s", err)
		}
		return
	}
	// テスト用のinputにlet文が3つあるため
	if len(program.Statements) != 3 {
		t.Errorf("program.Statements does not contain 3 statement. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%s", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%v", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral not %s. got=%s", name, letStmt.Name.Value)
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`
	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	if program == nil {
		t.Error("ParseProgram returned nil")
		return
	}
	if len(p.errors) > 0 {
		for _, err := range p.errors {
			t.Errorf("parser error: %s", err)
		}
		return
	}
	// テスト用のinputにlet文が3つあるため
	if len(program.Statements) != 3 {
		t.Errorf("program.Statements does not contain 3 statement. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%v", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %s", returnStmt.TokenLiteral())
		}
	}
}
