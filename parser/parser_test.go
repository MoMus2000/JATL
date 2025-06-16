package parser

import (
	"JATL/ast"
	"JATL/lexer"
	"testing"
)

func checkParserError(t *testing.T, p *Parser){
  errors := p.errors

  if len(errors) == 0 {
    return
  }

  if len(errors) > 0{
    t.Errorf("Parser has %d errors", len(errors))
    for _, msg := range errors {
      t.Errorf("parser error: %q", msg)
    }
  }
  t.FailNow()
}


func TestLetStatement(t *testing.T){

  input := `
    let x      = 5;
    let y      = 10;
    let foobar = 69420;
  `

  l := lexer.New(input)
  p := New(l)
  program := p.ParseProgram()
  checkParserError(t, p)


  if program == nil {
    t.Fatalf("ParseProgram() returned nil")
  }

  if len(program.Statement) != 3 {
    t.Fatalf("program.Statements does not contain 3 statements. Got %d",
      len(program.Statement),
    )
  }

  tests := []struct{
    expectedIdentifier string }{
    {"x"},
    {"y"},
    {"foobar"},
  }

  for i, tt := range tests{
    stmt := program.Statement[i]
    if !testLetStatement(t, stmt, tt.expectedIdentifier){
      return
    }
  }
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
  if s.TokenLiteral() != "let" {
    t.Errorf("s.TokenLiteral() not `let` got %s", s.TokenLiteral())
    return false
  }

  letStmt, ok := s.(*ast.LetStatement)

  if !ok {
    t.Errorf("s not *ast.LetStatement, got %T", letStmt)
    return false
  }

  if letStmt.Name.Value != name {
    t.Errorf("letStmt.Name.Value does not equal name, got %s", letStmt.Name.Value) 
    return false
  }

  if letStmt.Name.TokenLiteral() != name {
    t.Errorf("letStmt.Name does not equal name, got %s", letStmt.Name) 
    return false
  }

  return true
}

