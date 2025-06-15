package token

type TokenType string

const (
  ILLEGAL    = "ILLEGAL"
  EOF        = "EOF"

  IDENT      = "IDENT"
  INT        = "INT"

  ASSIGN     = "="
  PLUS       = "+"
  MINUS      = "-"
  BANG       = "!"
  BANG_EQUAL = "!="
  EQUAL      = "=="
  ASTERISK   = "*"
  LESS       = "<"
  GREATER    = ">"
  SLASH      = "/"

  COMMA      = ","
  SEMICOLON  = ";"

  LPAREN     = "("
  RPAREN     = ")"

  LBRACE     = "{"
  RBRACE     = "}"

  FUNCTION   = "FUNCTION"
  LET        = "LET"
  TRUE       = "TRUE"
  FALSE      = "FALSE"
  IF         = "IF"
  ELSE       = "ELSE"
  RETURN     = "RETURN"
)

var keywords = map[string]TokenType {
  "fn"     : FUNCTION,
  "let"    : LET,
  "false"  : FALSE,
  "true"   : TRUE,
  "if"     : IF,
  "else"   : ELSE,
  "return" : RETURN,
}

func LookupIdent(identString string) TokenType{
  keyword, found := keywords[identString]

  if !found {
    return IDENT
  }

  return keyword
}

type Token struct {
  Type       TokenType
  Literal    string
}

