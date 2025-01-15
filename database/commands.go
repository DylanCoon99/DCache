package database

import (
	"bytes"
	"fmt"
)


func handleInit(name string) string {
	// returns a message
	defaultSize := 10
	d := InitDatabase(name, uint32(defaultSize))
	//fmt.Println("in the init function")
	//fmt.Printf("%p \n", d)


	LoadLog(d)


	return "Successfully initialized database"
}



func handleSet(key, value string) (string, error) {

	// sets the key to the value in memory and returns a message (string)
	fmt.Println("in the set function")
	//fmt.Println(h)

	
	buf := new(bytes.Buffer)
	buf.WriteString(value)
	err := AddEntry(StringType, key, buf)

	if err != nil {
		return "Set Unsuccessful", err
	}


	//fmt.Println(database.D)
	//fmt.Printf("%p \n", database.D)
	

	return "Set Successful", nil
}


func handleGet(key string) (string, error) {

	// gets the value in memory for the given key and returns a message (string)
	fmt.Println(key)

	present, entry := GetEntry(key)

	//fmt.Println(database.D)

	if !present {
		return "Data " + key + " does not exist", nil 
	}
	

	

	return entry.Data.String() , nil

}