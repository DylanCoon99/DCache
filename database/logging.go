package database


import (
	"path/filepath"
	//"encoding/binary"
	//"encoding/json"
	"strings"
	"bufio"
	"log"
	//"fmt"
	"os"
	//"bytes"
)


var logPath string



func GetAbsPath(d* Database) (string, error) {

	logPath = "/mnt/c/Users/Dylan/My Documents/self_learning/DCache/database/" + d.Name + ".txt"

	absPath, err := filepath.Abs(logPath)
	if err != nil {
		return "", err
	}

	return absPath, nil

}



func Compact(fileName string) {

	// file is a file name

	var m map[string]string
	m = make(map[string]string)

	//file, err := os.OpenFile(fileName, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	file, err := os.Open(fileName)


	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()


	scanner := bufio.NewScanner(file)

	// iterate over the file
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		cmd := scanner.Text()

		// parse the cmd into a list
		cmdList := strings.Split(cmd, " ")

		if cmdList[0] == "get" || cmdList[0] == "init" || cmdList[0] == "echo"{
			// ignore
			continue
		}

		if cmdList[0] == "set" {

			key := cmdList[1]

			m[key] = cmd

		}


	}


	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}


	file, err = os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}


	// overwrite the file with the contents of the new data
	for _, value := range m {
		file.WriteString(value + "\n")
	}


}



func UpdateLog(d *Database, cmd string) {

	// over writes the log file with the current state of the database address space

	// locate the log file first

	file, err := os.OpenFile(d.Name + ".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	// whenever a cmd is sent, it is logged in the file 

	file.WriteString(cmd + "\n")


	return
}



func LoadLog(d *Database) string {

	// loads a log file and runs the cmds in the log file

	// locate the log file first

	Compact(d.Name + ".txt")

	file, err := os.Open(d.Name + ".txt")
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	defer file.Close()


	// iterate over the log file and run the cmds line by line
	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		cmd := scanner.Text()

		cmdType := string(cmd[0]) // set, get, etc.

		switch cmdType {
	    case "echo":
	    	return cmd[1:]
	        //return strings.Join(cmd[1:], " ")
	    case "set":
	        // handle the set command

	        ret, err := HandleSet(string(cmd[1]), string(cmd[2]))
	        if err != nil {
	            return err.Error()
	        }
	        return ret
	    case "get":
	        // handle the get command

	        key := string(cmd[1])
	        ret, err := HandleGet(key)
	        if err != nil {
	            return err.Error()
	        }
	        return ret
	    default:
	        // handle the default case
	        return "You done fucked up"
	    }
	}



	return ""


}