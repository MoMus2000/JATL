package parser

import (
	"JATL/ast"
	"JATL/lexer"
	"testing"
)


func TestIdentifierExpression(t *testing.T){
  input := "foobar"
  l := lexer.New(input)
  p := New(l)
  program := p.ParseProgram()
  checkParserError(t, p)
  if program == nil {
    t.Fatalf("ParseProgram() returned nil")
  }
  if len(program.Statement) != 1 {
    t.Fatalf("program.Statements does not contain 3 statements. Got %d",
      len(program.Statement),
    )
  }
  stmt, ok := program.Statement[0].(*ast.ExpresssionStatement)
  if !ok {
  }
  ident, ok := stmt.Expression.(*ast.Identifier)
  if !ok{
  }
  if ident.TokenLiteral() != "foobar"{
    t.FailNow()
  }
}

func TestIntExpression(t *testing.T){
  input := "5"
  l := lexer.New(input)
  p := New(l)
  program := p.ParseProgram()
  checkParserError(t, p)
  if program == nil {
    t.Fatalf("ParseProgram() returned nil")
  }
  if len(program.Statement) != 1 {
    t.Fatalf("program.Statements does not contain 3 statements. Got %d",
      len(program.Statement),
    )
  }
  stmt, ok := program.Statement[0].(*ast.ExpresssionStatement)
  if !ok {
  }
  ident, ok := stmt.Expression.(*ast.IntegerLiteral)
  if !ok{
  }
  if ident.TokenLiteral() != "5"{
    t.FailNow()
  }
}
