package sespage

import (
	// "errors"
	"fmt"
	// "log"
)

type Element interface {
	Decode(data []byte) error
	// Encode() ([]byte, error)
	Length() int32
	// Empty() bool
	// Spec()
}

type PageHead struct {
	Elements []Element

	name string
}

func (e *PageHead) Length() {
	length := int32(0)
	for _, e := range e.Elements {
		length += e.Length()
	}
	return length
}

for (h *PageHead) Decode(data []byte) error {
	offset := int32(0)
	for _, e := range e.Elements {
		err := e.Decode(data[offset:])
		if err != nil {
			return fmt.Errorf("Decode [%s]: %s", e.name, err.Error())
		}
		offset += e.Length()
	}
	return nil
}

type PagePayload struct {

}

type Page struct {

}

type Int32Element struct {
	Value int32

	name string
	length int32
}

func (e *Int32Element) Name() string {
	return e.name
}

func (e *Int32Element) Length() int32 {
	return e.length
}

func (e *Int32Element) Decode(data []byte) error {
	var err error

	if len(data) == 0 {
		err = fmt.Errorf("Got 0 byte.")
	} else {
		switch e.length {
			case 1:
				err = e.DecodeByte1(data[0])
			case 2:
				err = e.DecodeByte2(data)
			case 4:
				err = e.DecodeByte4(data)
			default:
				err = fmt.Errorf("Can't translate %d bytes data to int32.", e.length)
		}
	}

	return err
}

func (e *Int32Element) DecodeByte1(data byte) error {
	e.Value = int32(data)
	return nil
}

func (e *Int32Element) DecodeByte2(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("Expect 2 bytes data, got %d", len(data))
	}
	e.Value = (int32(data[0]) << 8) | int32(data[1])
	return nil
}

func (e *Int32Element) DecodeByte4(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Expect 4 bytes data, got %d", len(data))
	}
	e.Value = (int32(data[0]) << 24) | (int32(data[1]) << 16) | (int32(data[2]) << 8) | int32(data[3])
	return nil
}

func NewInt32Element(name string, length int32) *Int32Element {
	e := &Int32Element{name: name, length: length}
	return e
}


// page1head := PageHead{
// 	Elements: []Element{
// 		NewPageCode(0x01)
// 		NewInt32Element("abc", 4),
// 		NewInt32Element("adf", 2),
// 		NewGenCode()
// 	}
// }

