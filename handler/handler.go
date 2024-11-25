package handler


import (
    //"fmt"
    "strings"
    repl "github.com/openengineer/go-repl"
)


// implements Prompt(), Tab(), Eval()
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
    
    inputList := strings.Split(line, " ")
    numArgs := len(inputList)
    
    if numArgs != 2 && numArgs != 3 {
        return `Incorrect number of numArgs
Please Enter a Valid Command`
    }

    cmd := inputList[0]

    switch cmd {
    case "echo":
        return strings.Join(inputList[1:], " ")
    case "set":
        // handle the set command
        ret, err := handleSet("key", "value")
        if err != nil {
            return err.Error()
        }
        return ret
    case "get":
        // handle the get command
        ret, err := handleGet("key")
        if err != nil {
            return err.Error()
        }
        return ret
    default:
        // handle the default case
        return "You done fucked up"
    }
    
    return ""
}