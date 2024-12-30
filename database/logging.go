package database



import (
	"path/filepath"
	"encoding/json"
	"bufio"
	"log"
	//"fmt"
	"os"
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



func UpdateLog(d *Database) {

	// over writes the log file with the current state of the database address space

	// locate the log file first

	file, err := os.Create(d.Name + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// iterate over the address space
	addressSpace := d.AddressSpace

	for i := 0; i < len(addressSpace); i ++ {
		data, _ := json.Marshal(addressSpace[i])
		_, err = file.WriteString(string(data) + "\n")
		if err != nil {
			log.Fatal(err)
		}

	}

}



func LoadLog(d *Database) {

	// loads a log file to the address space

	// locate the log file first

	file, err := os.Open(d.Name + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var entry Entry


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// parse the line into an entry struct
		err := json.Unmarshal([]byte(scanner.Text()), &entry)

		// write this entry to the database address space
		
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}




}