package handler


import (
    //"fmt"
    "strings"
    repl "github.com/openengineer/go-repl"
    "github.com/DylanCoon99/DCache/database"
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
    case "init":
        // init <name>
        msg := database.handleInit(inputList[1])
        return msg
    case "echo":
        //database.UpdateLog(database.D, line)
        return strings.Join(inputList[1:], " ")
    case "set":
        // handle the set command
        //database.UpdateLog(database.D, line)
        ret, err := database.handleSet(inputList[1], inputList[2])
        if err != nil {
            return err.Error()
        }
        return ret
    case "get":
        // handle the get command
        //database.UpdateLog(database.D, line)
        key := inputList[1]
        ret, err := database.handleGet(key)
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