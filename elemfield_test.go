package sespage

import (
	"testing"
	"log"
	"encoding/json"
)

func ElementCommonTest(data []byte, elem Element) {
	FieldCommonTest(data, elem)
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
		FieldCommonTest(data[i:], element)
		log.Printf("0x%x %s\n", element.Code, element.Name)
	}
	IntFieldCommonTest(element)
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
	ElementCommonTest(data, elem)
	if elem.(*CommonControlElement).RstSwap.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0x2f}
	ElementCommonTest(data, elem)
	if elem.(*CommonControlElement).Disable.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0x4f}
	ElementCommonTest(data, elem)
	if elem.(*CommonControlElement).Prdfail.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0x8f}
	ElementCommonTest(data, elem)
	if elem.(*CommonControlElement).Select.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0xff}
	ElementCommonTest(data, elem)
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
	ElementCommonTest(data, elem)
	if elem.(*CommonStatusElement).ElementStatusCode.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0x12}
	ElementCommonTest(data, elem)
	if elem.(*CommonStatusElement).Swap.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0x23}
	ElementCommonTest(data, elem)
	if elem.(*CommonStatusElement).Disable.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0x44}
	ElementCommonTest(data, elem)
	if elem.(*CommonStatusElement).Prdfail.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	data = []byte{0xf5}
	ElementCommonTest(data, elem)
	if elem.(*CommonStatusElement).ElementStatusCode.Uint8() != 5 ||
		elem.(*CommonStatusElement).Swap.Uint8() != 1 ||
		elem.(*CommonStatusElement).Disable.Uint8() != 1 ||
		elem.(*CommonStatusElement).Prdfail.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
}

func TestThresholdControlElement(t *testing.T) {
	elem := CreateThresholdControlElement()
	data := []byte{1,0xae,3,0xff}
	ElementCommonTest(data, elem)
	if elem.(*ThresholdControlElement).RequestedHighCriticalThreshold.Uint8() != 1 {
		panic("Encode/Decode Error")
	}
	if elem.(*ThresholdControlElement).RequestedHighWarningThreshold.Uint8() != 0xae {
		panic("Encode/Decode Error")
	}
	if elem.(*ThresholdControlElement).RequestedLowWarningThreshold.Uint8() != 3 {
		panic("Encode/Decode Error")
	}
	if elem.(*ThresholdControlElement).RequestedLowCriticalThreshold.Uint8() != 0xff {
		panic("Encode/Decode Error")
	}
}


func TestThresholdStatusElement(t *testing.T) {
	elem := CreateThresholdStatusElement()
	data := []byte{10,0xbc,31,0xee}
	ElementCommonTest(data, elem)
	if elem.(*ThresholdStatusElement).HighCriticalThreshold.Uint8() != 10 {
		panic("Encode/Decode Error")
	}
	if elem.(*ThresholdStatusElement).HighWarningThreshold.Uint8() != 0xbc {
		panic("Encode/Decode Error")
	}
	if elem.(*ThresholdStatusElement).LowWarningThreshold.Uint8() != 31 {
		panic("Encode/Decode Error")
	}
	if elem.(*ThresholdStatusElement).LowCriticalThreshold.Uint8() != 0xee {
		panic("Encode/Decode Error")
	}
}

func TestUnspecifiedControlElement(t *testing.T) {
	elem := CreateUnspecifiedControlElement()
	data := []byte{0xff, 0x20, 0x51, 0xff}
	ElementCommonTest(data, elem)
}

func TestUnspecifiedStatusElement(t *testing.T) {
	elem := CreateUnspecifiedStatusElement()
	data := []byte{0xff, 0x20, 0x51, 0xff}
	ElementCommonTest(data, elem)
}