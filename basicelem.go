package sespage

import (
	// "errors"
	"fmt"
	"strings"
	"strconv"
	// "log"
)

/***************************************************************************
* Basic Element Interface
***************************************************************************/
type Element interface {
	Decode(data []byte) error
	Encode() ([]byte, error)
	Length() int32
}

// List

// type Page interface {
// 	DecodeHead(data []byte) error
// 	Decode(data []byte, context ProtocolContext) error
// 	EncodeHead(context ProtocolContext) ([]byte, error)
// 	Encode() ([]byte, error)
// 	HeadLength() int32
// 	Length() int32
// }

type BytesField interface {
	Element
	Bytes() []byte
	SetBytes([]byte)
}

type Int8Field interface {
	Element
	Int8() int8
	SetInt8(int8)
}

type Uint8Field interface {
	Element
	Uint8() uint8
	SetUint8(uint8)
}

type Int16Field interface {
	Element
	Int16() int16
	SetInt16(int16)
}

type Uint16Field interface {
	Element
	Uint16() uint16
	SetUint16(uint16)
}

type Int32Field interface {
	Element
	Int32() int32
	SetInt32(int32)
}

type Uint32Field interface {
	Element
	Uint32() uint32
	SetUint32(uint32)
}

type Int64Field interface {
	Element
	Int64() int64
	SetInt64(int64)
}

type Uint64Field interface {
	Element
	Uint64() uint64
	SetUint64(uint64)
}

type BoolField interface {
	Element
	Bool() bool
	SetBool(bool)
}

type StringField interface {
	Element
	String() string
	SetString(string)
}

type HexStringField interface {
	Element
	String() string
	SetString(string)
}

/***************************************************************************
* Basic Elements
***************************************************************************/
type BytesElement []byte

func (e *BytesElement) Decode(data []byte) error {
	*e = data
	return nil
}

func (e *BytesElement) Encode() ([]byte, error) {
	return []byte(*e), nil
}

func (e *BytesElement) Length() int32 {
	return int32(len(*e))
}

func (e *BytesElement) Bytes() []byte {
	return []byte(*e)
}

func (e *BytesElement) SetBytes(v []byte) {
	*e = BytesElement(v)
}

func NewBytesElement() *BytesElement {
	var elem BytesElement
	return &elem
}

func CreateBytesElement() Element {
	var elem BytesElement
	return &elem
}

type Int8Element int8

func (e *Int8Element) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 bytes data, got %d", len(data))
	}	
	*e = Int8Element(data[0])
	return nil
}

func (e *Int8Element) Encode() ([]byte, error) {
	data := make([]byte, 1, 1)
	data[0] = byte(*e)
	return data, nil
}

func (e *Int8Element) Length() int32 {
	return 1
}

func (e *Int8Element) Int8() int8 {
	return int8(*e)
}

func (e *Int8Element) SetInt8(v int8) {
	*e = Int8Element(v)
}

func NewInt8Element() *Int8Element {
	var elem Int8Element
	return &elem
}

func CreateInt8Element() Element {
	var elem Int8Element
	return &elem
}

type Uint8Element uint8

func (e *Uint8Element) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 bytes data, got %d", len(data))
	}
	*e = Uint8Element(data[0])
	return nil
}

func (e *Uint8Element) Encode() ([]byte, error) {
	data := make([]byte, 1, 1)
	data[0] = byte(*e)
	return data, nil
}

func (e *Uint8Element) Length() int32 {
	return 1
}

func (e *Uint8Element) Uint8() uint8 {
	return uint8(*e)
}

func (e *Uint8Element) SetUint8(v uint8) {
	*e = Uint8Element(v)
}

func NewUint8Element() *Uint8Element {
	var elem Uint8Element
	return &elem
}

func CreateUint8Element() Element {
	var elem Uint8Element
	return &elem
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

func (e *Int16Element) Encode() ([]byte, error) {
	data := make([]byte, 2, 2)
	data[0] = byte((*e >> 8) & 0xff)
	data[1] = byte(*e & 0xff)
	return data, nil
}

func (e *Int16Element) Length() int32 {
	return 2
}

func (e *Int16Element) Int16() int16 {
	return int16(*e)
}

func (e *Int16Element) SetInt16(v int16) {
	*e = Int16Element(v)
}

func NewInt16Element() *Int16Element {
	var elem Int16Element
	return &elem
}

func CreateInt16Element() Element {
	var elem Int16Element
	return &elem
}

type Uint16Element uint16

func (e *Uint16Element) Decode(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("Expect 2 bytes data, got %d", len(data))
	}	
	v := (uint16(data[0]) << 8) | uint16(data[1])
	*e = Uint16Element(v)

	return nil
}

func (e *Uint16Element) Encode() ([]byte, error) {
	data := make([]byte, 2, 2)
	data[0] = byte((*e >> 8) & 0xff)
	data[1] = byte(*e & 0xff)
	return data, nil
}

func (e *Uint16Element) Length() int32 {
	return 2
}

func (e *Uint16Element) Uint16() uint16 {
	return uint16(*e)
}

func (e *Uint16Element) SetUint16(v uint16) {
	*e = Uint16Element(v)
}

func NewUint16Element() *Uint16Element {
	var elem Uint16Element
	return &elem
}

func CreateUint16Element() Element {
	var elem Uint16Element
	return &elem
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

func (e *Int32Element) Encode() ([]byte, error) {
	data := make([]byte, 4, 4)
	data[0] = byte((*e >> 24) & 0xff)
	data[1] = byte((*e >> 16) & 0xff)
	data[2] = byte((*e >> 8) & 0xff)
	data[3] = byte(*e & 0xff)
	return data, nil
}

func (e *Int32Element) Length() int32 {
	return 4
}

func (e *Int32Element) Int32() int32 {
	return int32(*e)
}

func (e *Int32Element) SetInt32(v int32) {
	*e = Int32Element(v)
}

func NewInt32Element() *Int32Element {
	var elem Int32Element
	return &elem
}

func CreateInt32Element() Element {
	var elem Int32Element
	return &elem
}

type Uint32Element uint32

func (e *Uint32Element) Decode(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Expect 4 bytes data, got %d", len(data))
	}	
	v := (uint32(data[0]) << 24 | uint32(data[1]) << 16 | uint32(data[2]) << 8 | uint32(data[3]))
	*e = Uint32Element(v)

	return nil
}

func (e *Uint32Element) Encode() ([]byte, error) {
	data := make([]byte, 4, 4)
	data[0] = byte((*e >> 24) & 0xff)
	data[1] = byte((*e >> 16) & 0xff)
	data[2] = byte((*e >> 8) & 0xff)
	data[3] = byte(*e & 0xff)
	return data, nil
}

func (e *Uint32Element) Length() int32 {
	return 4
}

func (e *Uint32Element) Uint32() uint32 {
	return uint32(*e)
}

func (e *Uint32Element) SetUint32(v uint32) {
	*e = Uint32Element(v)
}

func NewUint32Element() *Uint32Element {
	var elem Uint32Element
	return &elem
}

func CreateUint32Element() Element {
	var elem Uint32Element
	return &elem
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

func (e *Int64Element) Encode() ([]byte, error) {
	data := make([]byte, 8, 8)
	data[0] = byte((*e >> 56) & 0xff)
	data[1] = byte((*e >> 48) & 0xff)
	data[2] = byte((*e >> 40) & 0xff)
	data[3] = byte((*e >> 32) & 0xff)
	data[4] = byte((*e >> 24) & 0xff)
	data[5] = byte((*e >> 16) & 0xff)
	data[6] = byte((*e >> 8) & 0xff)
	data[7] = byte(*e & 0xff)
	return data, nil
}

func (e *Int64Element) Length() int32 {
	return 8
}

func (e *Int64Element) Int64() int64 {
	return int64(*e)
}

func (e *Int64Element) SetInt64(v int64) {
	*e = Int64Element(v)
}

func NewInt64Element() *Int64Element {
	var elem Int64Element
	return &elem
}

func CreateInt64Element() Element {
	var elem Int64Element
	return &elem
}

type Uint64Element uint64

func (e *Uint64Element) Decode(data []byte) error {
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

	*e = Uint64Element(v)

	return nil
}

func (e *Uint64Element) Encode() ([]byte, error) {
	data := make([]byte, 8, 8)
	data[0] = byte((*e >> 56) & 0xff)
	data[1] = byte((*e >> 48) & 0xff)
	data[2] = byte((*e >> 40) & 0xff)
	data[3] = byte((*e >> 32) & 0xff)
	data[4] = byte((*e >> 24) & 0xff)
	data[5] = byte((*e >> 16) & 0xff)
	data[6] = byte((*e >> 8) & 0xff)
	data[7] = byte(*e & 0xff)
	return data, nil
}

func (e *Uint64Element) Length() int32 {
	return 8
}

func (e *Uint64Element) Uint64() uint64 {
	return uint64(*e)
}

func (e *Uint64Element) SetUint64(v uint64) {
	*e = Uint64Element(v)
}

func NewUint64Element() *Uint64Element {
	var elem Uint64Element
	return &elem
}

func CreateUint64Element() Element {
	var elem Uint64Element
	return &elem
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

func (e *BoolElement) Encode() ([]byte, error) {
	data := make([]byte, 1, 1)
	if *e {
		data[0] = 1
	} else {
		data[0] = 0
	}
	return data, nil
}

func (e *BoolElement) Length() int32 {
	return 1
}

func (e *BoolElement) Bool() bool {
	return bool(*e)
}

func (e *BoolElement) SetBool(v bool) {
	*e = BoolElement(v)
}

func NewBoolElement() *BoolElement {
	var elem BoolElement
	return &elem
}

func CreateBoolElement() Element {
	var elem BoolElement
	return &elem
}

type StringElement string

func (e *StringElement) Decode(data []byte) error {
	v := strings.TrimRight(string(data), "\u0000")
	*e = StringElement(v)
	return nil
}

func (e *StringElement) Encode() ([]byte, error) {
	data := []byte(*e)
	return data, nil
}

func (e *StringElement) Length() int32 {
	return int32(len(*e))
}

func (e *StringElement) String() string {
	return string(*e)
}

func (e *StringElement) SetString(v string) {
	*e = StringElement(v)
}

func NewStringElement() *StringElement {
	var elem StringElement
	return &elem
}

func CreateStringElement() Element {
	var elem StringElement
	return &elem
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

func (e *HexStringElement) Encode() ([]byte, error) {
	length := len(*e)/2
	data := make([]byte, 0, length)
	s := string(*e)
	for i := 0; i < length; i++ {
		offset := i*2
		v, err := strconv.ParseUint(s[offset:offset+2], 16, 8)
		if err != nil {
			return data, err
		}
		data = append(data, byte(v))
	}
	return data, nil
}

func (e *HexStringElement) Length() int32 {
	return int32(len(*e)/2)
}

func (e *HexStringElement) String() string {
	return string(*e)
}

func (e *HexStringElement) SetString(v string) {
	*e = HexStringElement(v)
}

func NewHexStringElement() *HexStringElement {
	var elem HexStringElement
	return &elem
}

func CreateHexStringElement() Element {
	var elem HexStringElement
	return &elem
}

type Codes map[uint8]string

type CodeElement struct {
	Code uint8
	Name string
	codes Codes
}

func (e *CodeElement) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 byte data, got %d", len(data))
	}
	e.Code = uint8(data[0])
	e.Name = e.codes[e.Code]
	return nil
}

func (e *CodeElement) Encode() ([]byte, error) {
	return []byte{byte(e.Code)}, nil
}

func (e *CodeElement) Length() int32 {
	return 1
}

func (e *CodeElement) Uint8() uint8 {
	return e.Code
}

func (e *CodeElement) SetUint8(v uint8) {
	e.Code = v
	e.Name = e.codes[e.Code]
}



