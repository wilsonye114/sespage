package sespage

import (
	"testing"
	"log"
	"encoding/json"
)

func CommonTestElement(data []byte, elem Element) {
	CommonTestField(data, elem)
	js, _ := json.MarshalIndent(elem, "", "    ")
	log.Printf("%s\n", js)
}

func TestElemTypeCodeElement(t *testing.T) {
	element := NewElemTypeCodeElement()
	data := make([]byte, 0x100)
	for i := 0; i <= 0xff; i++ {
		data[i] = byte(i)
	}
	for i := 0; i <= 0xff; i++ {
		CommonTestField(data[i:], element)
		log.Printf("0x%x %s\n", element.Code, element.Name)
	}
	CommonTestIntField(element)
}

func TestElemTypeCodes(t *testing.T) {
	codes := NewElemTypeCodes()
	for i :=  0; i <= 0xff; i++ {
		log.Printf("%x %s\n", i, codes[uint8(i)])
	}
}

func TestCommonControlElement(t *testing.T) {
	elem := CreateCommonControlElement()
	data := []byte{0x1f}
	CommonTestElement(data, elem)
	if elem.(*CommonControlElement).RstSwap.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0x2f}
	CommonTestElement(data, elem)
	if elem.(*CommonControlElement).Disable.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0x4f}
	CommonTestElement(data, elem)
	if elem.(*CommonControlElement).Prdfail.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0x8f}
	CommonTestElement(data, elem)
	if elem.(*CommonControlElement).Select.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0xff}
	CommonTestElement(data, elem)
	if elem.(*CommonControlElement).RstSwap.Uint8() != 1 ||
		elem.(*CommonControlElement).Disable.Uint8() != 1 ||
		elem.(*CommonControlElement).Prdfail.Uint8() != 1 ||
		elem.(*CommonControlElement).Select.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
}

func TestCommonStatusElement(t *testing.T) {
	elem := CreateCommonStatusElement()
	data := []byte{0x01}
	CommonTestElement(data, elem)
	if elem.(*CommonStatusElement).ElementStatusCode.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0x12}
	CommonTestElement(data, elem)
	if elem.(*CommonStatusElement).Swap.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0x23}
	CommonTestElement(data, elem)
	if elem.(*CommonStatusElement).Disable.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0x44}
	CommonTestElement(data, elem)
	if elem.(*CommonStatusElement).Prdfail.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0xf5}
	CommonTestElement(data, elem)
	if elem.(*CommonStatusElement).ElementStatusCode.Uint8() != 5 ||
		elem.(*CommonStatusElement).Swap.Uint8() != 1 ||
		elem.(*CommonStatusElement).Disable.Uint8() != 1 ||
		elem.(*CommonStatusElement).Prdfail.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
}