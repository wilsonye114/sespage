package sespage

import (
	"testing"
	"log"
	"encoding/json"
)

func TestVpdypeCodeElement(t *testing.T) {
	element := NewVpdTypeCodeElement()
	Uint8CodeElementCommonTest(element)
}

func TestMidplaneVpdDataElement(t *testing.T) {
	element := NewMidplaneVpdDataElement()
	element.BoardProductName.SetString("Board-001-adc     ")
	element.BoardPartNumber.SetString("abc864d3120000000000")
	element.BoardSerialNumber.SetString("80b27634dc6400000000")
	element.BoardHardwareECLevel.SetString("98656744")
	element.ProductName.SetString("Product-0210      ")
	element.PorductPartNumber.SetString("f46783aadc23bb00")
	element.ProductSerialNumber.SetString("f46fe3a6c2c3b000")
	element.ProductVersion.SetString("004A")
	js, _ := json.MarshalIndent(element, "", "    ")
	log.Printf("%s\n", js)
	data, err := element.Encode()
	if err != nil {
		panic(err)
	}
	FieldCommonTest(data, element)
	js, _ = json.MarshalIndent(element, "", "    ")
	log.Printf("%s\n", js)

}

func TestCanisterVpdDataElement(t *testing.T) {
	element := NewCanisterVpdDataElement()
	element.BoardProductName.SetString("Board-001-adc   ")
	element.BoardPartNumber.SetString("abc864d3120000000000")
	element.BoardSerialNumber.SetString("80b27634dc6400000000")
	element.BoardHardwareECLevel.SetString("9865")
	js, _ := json.MarshalIndent(element, "", "    ")
	log.Printf("%s\n", js)
	data, err := element.Encode()
	if err != nil {
		panic(err)
	}
	FieldCommonTest(data, element)
	js, _ = json.MarshalIndent(element, "", "    ")
	log.Printf("%s\n", js)

}