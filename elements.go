package sespage

import (
	// "errors"
	"fmt"
	"strings"
	// "log"
)

type Element interface {
	Decode(data []byte) error
	// Encode() ([]byte, error)
	Length() int32
	// Empty() bool
	// Spec()
}

type Elements []Element

func (e Elements) Length() int32 {
	length := int32(0)
	for _, element := range e {
		length += element.Length()
	}
	return length
}

type Page interface {
	Decode(data []byte) error
	Length() int32
	DecodeHead(data []byte) error
	HeadLength() int32
}


type Int8Element int8

func (e *Int8Element) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 bytes data, got %d", len(data))
	}	
	*e = Int8Element(data[0])
	return nil
}

func (e *Int8Element) Length() int32 {
	return 1
}

type UInt8Element uint8

func (e *UInt8Element) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 bytes data, got %d", len(data))
	}
	*e = UInt8Element(data[0])
	return nil
}

func (e *UInt8Element) Length() int32 {
	return 1
}

type Int16Element int16

func (e *Int16Element) Decode(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("Expect 2 bytes data, got %d", len(data))
	}	
	v := (uint16(data[0]) << 8) | uint16(data[1])
	*e = Int16Element(v)

	return nil
}

func (e *Int16Element) Length() int32 {
	return 2
}

type UInt16Element uint16

func (e *UInt16Element) Decode(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("Expect 2 bytes data, got %d", len(data))
	}	
	v := (uint16(data[0]) << 8) | uint16(data[1])
	*e = UInt16Element(v)

	return nil
}

func (e *UInt16Element) Length() int32 {
	return 2
}


type Int32Element int32

func (e *Int32Element) Decode(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Expect 4 bytes data, got %d", len(data))
	}	
	v := (uint32(data[0]) << 24 | uint32(data[1]) << 16 | uint32(data[2]) << 8 | uint32(data[3]))
	*e = Int32Element(v)

	return nil
}

func (e *Int32Element) Length() int32 {
	return 4
}

type UInt32Element uint32

func (e *UInt32Element) Decode(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Expect 4 bytes data, got %d", len(data))
	}	
	v := (uint32(data[0]) << 24 | uint32(data[1]) << 16 | uint32(data[2]) << 8 | uint32(data[3]))
	*e = UInt32Element(v)

	return nil
}

func (e *UInt32Element) Length() int32 {
	return 4
}


type UInt64Element uint64

func (e *UInt64Element) Decode(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("Expect 8 bytes data, got %d", len(data))
	}	
	v := (uint64(data[0]) << 56 |
		  uint64(data[1]) << 48 |
		  uint64(data[2]) << 40 |
		  uint64(data[3]) << 32 |
		  uint64(data[4]) << 24 |
		  uint64(data[5]) << 16 |
		  uint64(data[6]) << 8 |
		  uint64(data[7]))

	*e = UInt64Element(v)

	return nil
}

func (e *UInt64Element) Length() int32 {
	return 8
}

type Int64Element int64

func (e *Int64Element) Decode(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("Expect 8 bytes data, got %d", len(data))
	}	
	v := (uint64(data[0]) << 56 |
		  uint64(data[1]) << 48 |
		  uint64(data[2]) << 40 |
		  uint64(data[3]) << 32 |
		  uint64(data[4]) << 24 |
		  uint64(data[5]) << 16 |
		  uint64(data[6]) << 8 |
		  uint64(data[7]))

	*e = Int64Element(v)

	return nil
}

func (e *Int64Element) Length() int32 {
	return 8
}

type StringElement string

func (e *StringElement) Decode(data []byte) error {
	v := strings.TrimRight(string(data), "\u0000")
	*e = StringElement(v)
	return nil
}

func (e *StringElement) Length() int32 {
	return int32(len(*e))
}

type BoolElement bool

func (e *BoolElement) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 bytes data, got %d", len(data))
	}
	v := int8(data[0])
	if v != 0 {
		*e = true
	} else {
		*e = false
	}
	return nil
}

func (e *BoolElement) Length() int32 {
	return 1
}

type BytesElement []byte

func (e *BytesElement) Decode(data []byte) error {
	*e = data
	return nil
}

func (e *BytesElement) Length() int32 {
	return int32(len(*e))
}


type UInt8SliceElement []uint8

func (e *UInt8SliceElement) Decode(data []byte) error {
	v := make(UInt8SliceElement, 0, 16)
	for _, b := range data {
		v = append(v, uint8(b))
	}
	*e = v
	return nil
}

func (e *UInt8SliceElement) Length() int32 {
	return int32(len(*e))
}

func (e *UInt8SliceElement) String() string {
	str := ""
	for _, item := range *e {
		str += fmt.Sprintf("%02x", item)
	}
	return str
}

type HexStringElement string

func (e *HexStringElement) Decode(data []byte) error {
	v := ""
	for _, b := range data {
		v += fmt.Sprintf("%02x", uint8(b))
	}
	*e = HexStringElement(v)
	return nil
}

func (e *HexStringElement) Length() int32 {
	return int32(len(*e))
}

