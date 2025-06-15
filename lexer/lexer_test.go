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

    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d] - tokentype wrong. expected [%s], but got [%s]",
          i, tt.expectedType, tok.Type,
      )
    }

    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d] - literal wrong. expected [%s], but got [%s]",
          i, tt.expectedLiteral, tok.Literal,
      )
    }

  }

}
