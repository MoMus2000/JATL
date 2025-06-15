package lexer

import "JATL/token"


type Lexer struct {
  input        string
  position     int
  readPosition int
  ch           byte
}


func New(input string) *Lexer{
  return &Lexer{input: input}
}

func (l *Lexer) NextToken() token.Token{
  return token.Token{}
}

// readPosition is used to peek ahead
func (l *Lexer) readChar() {
  if l.readPosition >= len(l.input) {
    l.ch = 0 // ASCII Code for NULL or EOF for us
  } else {
    l.ch = l.input[l.readPosition]
  }
  l.position = l.readPosition
  l.readPosition += 1
}

