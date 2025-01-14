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


	// open the file, remove the get cmds, keep track of latest set cmd for every key, compact

	// map[KeyType]ValueType --> map[string]string

	//    maps key to the latest set cmd for that key

	// iterate over the cmds in the file,
	//		- if the cmd is a get cmd --> ignore
	//		- if the cmd is a set cmd
	//			- if the cmd is already in the map --> update
	// 			- if the cmd is not in the map --> put it in the map 

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

		if cmdList[0] == "get" || cmdList[0] == "init" {
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


	for _, value := range m {
		file.WriteString(value + "\n")
	}


	// overwrite the file with the contents of the new data

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



func LoadLog(d *Database) {

	// loads a log file to the address space

	// locate the log file first

	file, err := os.Open(d.Name + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	return


}