package database

import (
	"fmt"
	"testing"
	//"bytes"
)



func TestInitDatabase(t *testing.T) {


	var size uint32

	size = 600

	var d Database

	d.InitDatabase(size)


	fmt.Println(d.MaxSize)


}




func TestHash(t *testing.T) {

	ret := Hash("This is a bunch of nonsense")
	fmt.Printf("This is the result %d: ", ret)
}




/*
func TestAddEntry(t *testing.T) {

	var size uint32

	size = 10

	var d Database

	d.InitDatabase(size)

	d1 := bytes.NewBuffer([]byte("This is test data"))

	d.AddEntry(StringType, "key", d1)


	fmt.Println(d.AddressSpace)

}
*/