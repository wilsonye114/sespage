package sespage

import (
	"testing"
	"fmt"
)

func TestInt8Element(t *testing.T) {
	fmt.Printf("==================== TestInt8Element ====================\n")
	var element Int8Element
	bytes := []byte{0xff,2,3,4,5,6,7}
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t, length: %d\n", element, element.Length())
}

func TestUInt8Element(t *testing.T) {
	fmt.Printf("==================== TestUInt8Element ====================\n")
	var element UInt8Element
	bytes := []byte{0xff,2,3,4,5,6,7}
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t, length: %d\n", element, element.Length())
}

func TestInt16Element(t *testing.T) {
	fmt.Printf("==================== TestInt16Element ====================\n")
	var element Int16Element
	bytes := []byte{0xff,2,3,4,5,6,7}
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t, length: %d\n", element, element.Length())
}

func TestUInt16Element(t *testing.T) {
	fmt.Printf("==================== TestUInt16Element ====================\n")
	var element UInt16Element
	bytes := []byte{0xff,2,3,4,5,6,7}
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t, length: %d\n", element, element.Length())
}

func TestUInt32Element(t *testing.T) {
	fmt.Printf("==================== TestUInt32Element ====================\n")
	var element UInt32Element
	bytes := []byte{0xff,2,3,4,5,6,7}
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t, length: %d\n", element, element.Length())
}

func TestInt32Element(t *testing.T) {
	fmt.Printf("==================== TestInt32Element ====================\n")
	var element Int32Element
	bytes := []byte{0xff,2,3,4,5,6,7}
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t, length: %d\n", element, element.Length())
}

func TestInt64Element(t *testing.T) {
	fmt.Printf("==================== TestInt64Element ====================\n")
	var element Int64Element
	bytes := []byte{0xff,2,3,4,5,6,7,8}
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t, length: %d\n", element, element.Length())
}

func TestUInt64Element(t *testing.T) {
	fmt.Printf("==================== TestUInt64Element ====================\n")
	var element UInt64Element
	bytes := []byte{0xff,2,3,4,5,6,7,8}
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t, length: %d\n", element, element.Length())
}

func TestStringElement(t *testing.T) {
	fmt.Printf("==================== TestStringElement ====================\n")
	var element StringElement
	bytes := []byte{60,61,62,63,64,65,66,67,68}
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t, str: '%s', length: %d\n", element, element, element.Length())
}

func TestBoolElement(t *testing.T) {
	fmt.Printf("==================== TestBoolElement ====================\n")
	var element BoolElement
	bytes := []byte{0,1,2,3}
	err := element.Decode(bytes[1:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t, length: %d\n", element, element.Length())
}

func TestByteElement(t *testing.T) {
	fmt.Printf("==================== TestByteElement ====================\n")
	var element BytesElement
	bytes := []byte{0,1,2,3}
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t, length: %d\n", element, element.Length())
}

func TestUInt8SliceElement(t *testing.T) {
	fmt.Printf("==================== TestUInt8SliceElement ====================\n")
	var element UInt8SliceElement
	bytes := []byte{0,1,2,3}
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t, str: %s length: %d\n", element, element.String(), element.Length())
}

func TestHexStringElement(t *testing.T) {
	fmt.Printf("==================== TestHexStringElement ====================\n")
	var element HexStringElement
	bytes := []byte{0,1,2,3}
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t, str: %s length: %d\n", element, element, element.Length())
}