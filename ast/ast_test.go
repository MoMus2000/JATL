package ast

import (
	"JATL/token"
	"testing"
)

func TestString(t *testing.T) {
  program := Program{
    Statement: []Statement {
      &LetStatement{
        Token: token.Token{Type: token.LET, Literal: "let"},
        Name : &Identifier{
          Token: token.Token{Type: token.IDENT, Literal: "myVar"},
          Value: "myVar",
        },
        Value: &Identifier{
          Token: token.Token{Type: token.IDENT, Literal: "myVar"},
          Value: "myVar",
        },
      },
    },
  }
  out := program.String()
  if out != "let myVar = myVar;" {
    t.Errorf("program.String() wrong, got =%q", out)
  }
}
