package parser

import (
	"JATL/ast"
	"JATL/lexer"
	"testing"
)


func TestParsingPrefixExpression(t *testing.T){
  prefixTests := []struct{
    input        string
    op           string
    integerValue string
  }{
    {"!5;", "!", "5"},
    {"-15;", "-", "15"},
  }

  for _, tt := range prefixTests {

    input   := lexer.New(tt.input)
    p       := New(input)

    program := p.ParseProgram()

    checkParserError(t, p)

    if len(program.Statement) != 1 {
      t.Errorf("expected 1 statements got %d", len(program.Statement))
    }

    stmt, ok := program.Statement[0].(*ast.ExpresssionStatement)

    if !ok {
      t.Errorf("Error with parsing to Expr")
    }

    prefix, ok := stmt.Expression.(*ast.PrefixExpression)

    if !ok {
      t.Errorf("Error with parsing to Expr")
    }

    if prefix.Operator != tt.op {
      t.Errorf("Expected OP %s but got %s", prefix.Operator, tt.op)
    }

  }

}

