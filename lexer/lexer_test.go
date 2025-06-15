package lexer

import (
	"JATL/token"
	"testing"
)



func TestNextToken(t *testing.T){
  input := "=+(){},;"

  tests := []struct{
    expectedType     token.TokenType
    expectedLiteral  string
  }{
    {token.ASSIGN   , "="},
    {token.PLUS     , "+"},
    {token.LPAREN   , "("},
    {token.RPAREN   , ")"},
    {token.LBRACE   , "{"},
    {token.RBRACE   , "}"},
    {token.COMMA    , ","},
    {token.SEMICOLON, ";"},
  }

  l := New(input)

  for i, tt := range tests {
    tok := l.NextToken()

    if tok.TokenType != tt.expectedType {
      t.Fatalf("tests[%d] - tokentype wrong. expected [%s], but got [%s]",
          i, tt.expectedType, tok.TokenType,
      )
    }

    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d] - literal wrong. expected [%s], but got [%s]",
          i, tt.expectedLiteral, tok.Literal,
      )
    }

  }

}
