package database

import (
	"fmt"
	"testing"
	"bytes"
)


/*
func TestInitDatabase(t *testing.T) {


	var size uint32

	size = 10

	var d Database

	d.InitDatabase(size)


	fmt.Println(d.MaxSize)


}
*/



/*
func TestHash(t *testing.T) {

	var size uint32

	size = 10

	var d Database

	d.InitDatabase(size)

	ret := d.Hash("key1")
	fmt.Printf("key1 %d: ", ret)

	ret = d.Hash("key2")
	fmt.Printf("key2 %d: ", ret)

	ret = d.Hash("key3")
	fmt.Printf("key3 %d: ", ret)

}
*/





func TestAddEntry(t *testing.T) {

	var size uint32

	size = 10


	InitDatabase("a", size)


	d9 := bytes.NewBuffer([]byte("This is test data for key9"))

	AddEntry(StringType, "key9", d9)


	d10 := bytes.NewBuffer([]byte("This is test data for key10"))

	AddEntry(StringType, "key10", d10)

	
	/*
	d6 = bytes.NewBuffer([]byte("This is some updated data for key6"))
	d.AddEntry(StringType, "key6", d6)
	fmt.Println("///////////////")
	*/

	/*
	for i := 0; i < len(d.AddressSpace); i ++ {
		fmt.Println(d.AddressSpace[i])
	}
	*/


	// test collision logic

	_, entry9 := GetEntry("key9")
	fmt.Println(entry9)

	_, entry10 := GetEntry("key10")
	fmt.Println(entry10)

	//_, entry10 = GetEntry("key10")
	//fmt.Println(*entry9.Next)

}



/*

func TestUpdateLog(t *testing.T) {

	var size uint32

	size = 10
	name := "database1"

	var d *database.Database

	d = database.InitDatabase(name, size)


	d1 := bytes.NewBuffer([]byte("This is test data for key1"))

	database.AddEntry(database.StringType, "key1", d1)


	d2 := bytes.NewBuffer([]byte("This is test data for key2"))

	database.AddEntry(database.StringType, "key2", d2)



	UpdateLog(d)

	LoadLog(d)

}

*/

/*

func TestPath(t *testing.T) {


	var size uint32

	size = 10


	d := InitDatabase("a", size)

	path, err := GetAbsPath(d)

	if err != nil {
		fmt.Println("failed to get path")
		panic(err)
	}

	fmt.Println(path)

}

*/