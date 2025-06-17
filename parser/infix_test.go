package parser

import (
	"JATL/ast"
	"JATL/lexer"
	"testing"
)

func TestInfixExpression(t *testing.T){
  infixTests := []struct{
    input       string
    leftValue   int64
    op          string
    rightValue  int64
  }{
    {"5 + 5 ;", 5,  "+", 5},
    {"5 - 5 ;", 5,  "-", 5},
    {"5 * 5 ;", 5,  "*", 5},
    {"5 / 5 ;", 5,  "/", 5},
    {"5 < 5 ;", 5,  "<", 5},
    {"5 > 5 ;", 5,  ">", 5},
    {"5 == 5;", 5, "==", 5},
    {"5 != 5;", 5, "!=", 5},
  }

  for _, tt := range infixTests{
    input   := lexer.New(tt.input)
    p       := New(input)
    program := p.ParseProgram()

    checkParserError(t, p)
    
    if len(program.Statement) != 1 {
      t.Errorf("Expected Statement to be 1 but was %d", len(program.Statement))
    }

    stmt, ok := program.Statement[0].(*ast.ExpresssionStatement)

    if !ok {
      t.Errorf("Parsing Error")
    }

    infix, ok := stmt.Expression.(*ast.InfixExpression)

    if !ok {
      t.Errorf("Parsing Error")
    }

  }
  
}

