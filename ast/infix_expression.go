package ast

import (
	"JATL/token"
	"bytes"
)

type InfixExpression struct{
  Token token.Token
  Right Expression
  Op    string
  Left  Expression
}

func (ie *InfixExpression) TokenLiteral() string{
  return ie.Token.Literal
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) String() string{
  var out bytes.Buffer

  out.WriteString("(")
  out.WriteString(ie.Left.String())
  out.WriteString(ie.Op)
  out.WriteString(ie.Right.String())
  out.WriteString(")")

  return out.String()
}

