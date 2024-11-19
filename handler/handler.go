package handler


import (
    repl "github.com/openengineer/go-repl"
)

/*
type Handler interface {
  Prompt() string
  Tab(buffer string) string
  Eval(line string) string
}
*/


type MyHandler struct {
    R *repl.Repl
}

func (h MyHandler) Prompt() string {
    return "> "
}


func (h MyHandler) Tab(buffer string) string {
    return ""
}


func (h MyHandler) Eval(line string) string {
    
    if line == "gg" {
        return "good game"
    }

    return ""
}