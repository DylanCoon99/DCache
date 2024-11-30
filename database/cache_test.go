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

	var d Database

	d.InitDatabase(size)


	d9 := bytes.NewBuffer([]byte("This is test data for key9"))

	d.AddEntry(StringType, "key9", d9)


	d10 := bytes.NewBuffer([]byte("This is test data for key10"))

	d.AddEntry(StringType, "key10", d10)

	
	
	/*
	d6 = bytes.NewBuffer([]byte("This is some updated data for key6"))
	d.AddEntry(StringType, "key6", d6)
	fmt.Println("///////////////")
	*/

	
	for i := 0; i < len(d.AddressSpace); i ++ {
		fmt.Println(d.AddressSpace[i])
	}
	
	

}


