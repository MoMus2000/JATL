package ast

import (
	"JATL/token"
	"strconv"
)

type IntegerLiteral struct {
  Token token.Token
  Value int64
}

func (is *IntegerLiteral) expressionNode() {

}

func (is *IntegerLiteral) TokenLiteral() string{
  return is.Token.Literal
}

func (is *IntegerLiteral) String() string {
  str := strconv.FormatInt(is.Value, 10)
  return str
}

