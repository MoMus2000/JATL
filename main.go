package main

import (
	"JATL/repl"
	"fmt"
	"os"
	"os/user"
)

func main(){
  user, err := user.Current()
  if err != nil {
    panic(err)
  }
  fmt.Printf("Hello %s, this is JATL\n", user.Name)
  repl.Start(os.Stdin, os.Stdout)
}

