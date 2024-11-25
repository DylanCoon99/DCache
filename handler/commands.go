package handler




func handleSet(key, value string) (string, error) {

	// sets the key to the value in memory and returns a message (string)


	return "Set Successful", nil
}


func handleGet(key string) (string, error) {

	// gets the value in memory for the given key and returns a message (string)


	return "Get Successful", nil

}