package ast

import (
	"JATL/token"
	"bytes"
)

type PrefixExpression struct {
  Right      Expression
  Operator   string
  Token      token.Token
}

 func (pe *PrefixExpression) expressionNode() {}

 func (pe *PrefixExpression) TokenLiteral() string{
   return pe.Token.Literal
 }

 func (pe *PrefixExpression) String() string{
   var out bytes.Buffer
   out.WriteString("(")
   out.WriteString(pe.Operator)
   out.WriteString(pe.Right.String())
   out.WriteString(")")
   return out.String()
 }

