package repl

import (
	"JATL/lexer"
	"JATL/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer){
  scanner := bufio.NewScanner(in)

  for {
    fmt.Printf(PROMPT)
    scanned := scanner.Scan()
    if !scanned {
      return
    }

    line := scanner.Text()

    l := lexer.New(line)

    p := parser.New(l)
    program := p.ParseProgram()

    for _, stmt := range program.Statement{
      fmt.Printf("%+v\n", stmt.String())
    }

  }

}
