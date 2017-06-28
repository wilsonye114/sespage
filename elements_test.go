package sespage

import (
	"testing"
	"fmt"
)

func TestInt32Element(t *testing.T) {
	fmt.Printf("==================== TestInt32Element ====================\n")
	bytes := []byte{1,2,3,4,5,6,7}
	element4 := NewInt32Element("4-byte data to int32", 4)
	err := element4.Decode(bytes)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%v to %d\n", bytes, element4.Value)
	}

	element3 := NewInt32Element("3-byte data to int32", 3)
	err = element3.Decode(bytes)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v to %d\n", bytes, element3.Value)
	}

	element2 := NewInt32Element("2-byte data to int32", 2)
	err = element2.Decode(bytes)
	if err != nil {
		panic(err)
		fmt.Println(err)
	} else {
		fmt.Printf("%v to %d\n", bytes, element2.Value)
	}

	element1 := NewInt32Element("1-byte data to int32", 1)
	err = element1.Decode(bytes)
	if err != nil {
		panic(err)
		fmt.Println(err)
	} else {
		fmt.Printf("%v to %d\n", bytes, element1.Value)
	}

	element4err := NewInt32Element("4-byte data to int32", 4)
	err = element4.Decode(bytes[:3])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v to %d\n", bytes, element4err.Value)
	}

}