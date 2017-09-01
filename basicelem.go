package sespage

import (
	// "errors"
	"fmt"
	// "strings"
	"strconv"
	// "log"
)

/***************************************************************************
* Basic Element Interface
***************************************************************************/
type ElementLength uint64

func (l ElementLength) Byte() uint64 {
	return uint64(l/8)
}

func (l ElementLength) Bit() uint64 {
	return uint64(l)
}

func (l ElementLength) RemainderBit() uint64 {
	return uint64(l%8)
}

func (l ElementLength) IsAligned() bool {
	if l%8 == 0 {
		return true
	}
	return false
}

type Element interface {
	Decode(data []byte) error
	Encode() ([]byte, error)
	Length() ElementLength
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

type Uint1Field interface {
	Element
	Uint1() uint8
	SetUint1(uint8)
}

type Uint2Field interface {
	Element
	Uint2() uint8
	SetUint2(uint8)
}

type Uint3Field interface {
	Element
	Uint3() uint8
	SetUint3(uint8)
}

type Uint4Field interface {
	Element
	Uint4() uint8
	SetUint4(uint8)
}

type Uint5Field interface {
	Element
	Uint5() uint8
	SetUint5(uint8)
}

type Uint6Field interface {
	Element
	Uint6() uint8
	SetUint6(uint8)
}

type Uint7Field interface {
	Element
	Uint7() uint8
	SetUint7(uint8)
}

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

type Uint1Element uint8

func (e *Uint1Element) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 bytes data, got %d", len(data))
	}

	*e = Uint1Element(data[0]&0x01)
	return nil
}

func (e *Uint1Element) Encode() ([]byte, error) {
	data := make([]byte, 1, 1)
	data[0] = byte(*e)
	return data, nil
}

func (e *Uint1Element) Length() ElementLength {
	return ElementLength(1)
}

func (e *Uint1Element) Uint1() uint8 {
	return uint8(*e)
}

func (e *Uint1Element) SetUint1(v uint8) {
	if v > 0x01 {
		panic("value isn't 1 bit")
	}
	*e = Uint1Element(v)
}

func NewUint1Element() *Uint1Element {
	var elem Uint1Element
	return &elem
}

func CreateUint1Element() Element {
	var elem Uint1Element
	return &elem
}

type Uint2Element uint8

func (e *Uint2Element) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 bytes data, got %d", len(data))
	}

	*e = Uint2Element(data[0]&0x03)
	return nil
}

func (e *Uint2Element) Encode() ([]byte, error) {
	data := make([]byte, 1, 1)
	data[0] = byte(*e)
	return data, nil
}

func (e *Uint2Element) Length() ElementLength {
	return ElementLength(2)
}

func (e *Uint2Element) Uint2() uint8 {
	return uint8(*e)
}

func (e *Uint2Element) SetUint2(v uint8) {
	if v > 0x03 {
		panic("value isn't 2 bit")
	}
	*e = Uint2Element(v)
}

func NewUint2Element() *Uint2Element {
	var elem Uint2Element
	return &elem
}

func CreateUint2Element() Element {
	var elem Uint2Element
	return &elem
}

type Uint3Element uint8

func (e *Uint3Element) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 bytes data, got %d", len(data))
	}

	*e = Uint3Element(data[0]&0x07)
	return nil
}

func (e *Uint3Element) Encode() ([]byte, error) {
	data := make([]byte, 1, 1)
	data[0] = byte(*e)
	return data, nil
}

func (e *Uint3Element) Length() ElementLength {
	return ElementLength(3)
}

func (e *Uint3Element) Uint3() uint8 {
	return uint8(*e)
}

func (e *Uint3Element) SetUint3(v uint8) {
	if v > 0x07 {
		panic("value isn't 3 bit")
	}
	*e = Uint3Element(v)
}

func NewUint3Element() *Uint3Element {
	var elem Uint3Element
	return &elem
}

func CreateUint3Element() Element {
	var elem Uint3Element
	return &elem
}

type Uint4Element uint8

func (e *Uint4Element) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 bytes data, got %d", len(data))
	}

	*e = Uint4Element(data[0]&0x0f)
	return nil
}

func (e *Uint4Element) Encode() ([]byte, error) {
	data := make([]byte, 1, 1)
	data[0] = byte(*e)
	return data, nil
}

func (e *Uint4Element) Length() ElementLength {
	return ElementLength(4)
}

func (e *Uint4Element) Uint4() uint8 {
	return uint8(*e)
}

func (e *Uint4Element) SetUint4(v uint8) {
	if v > 0x0f {
		panic("value isn't 4 bit")
	}
	*e = Uint4Element(v)
}

func NewUint4Element() *Uint4Element {
	var elem Uint4Element
	return &elem
}

func CreateUint4Element() Element {
	var elem Uint4Element
	return &elem
}

type Uint5Element uint8

func (e *Uint5Element) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 bytes data, got %d", len(data))
	}

	*e = Uint5Element(data[0]&0x1f)
	return nil
}

func (e *Uint5Element) Encode() ([]byte, error) {
	data := make([]byte, 1, 1)
	data[0] = byte(*e)
	return data, nil
}

func (e *Uint5Element) Length() ElementLength {
	return ElementLength(5)
}

func (e *Uint5Element) Uint5() uint8 {
	return uint8(*e)
}

func (e *Uint5Element) SetUint5(v uint8) {
	if v > 0x1f {
		panic("value isn't 5 bit")
	}
	*e = Uint5Element(v)
}

func NewUint5Element() *Uint5Element {
	var elem Uint5Element
	return &elem
}

func CreateUint5Element() Element {
	var elem Uint5Element
	return &elem
}

type Uint6Element uint8

func (e *Uint6Element) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 bytes data, got %d", len(data))
	}

	*e = Uint6Element(data[0]&0x3f)
	return nil
}

func (e *Uint6Element) Encode() ([]byte, error) {
	data := make([]byte, 1, 1)
	data[0] = byte(*e)
	return data, nil
}

func (e *Uint6Element) Length() ElementLength {
	return ElementLength(6)
}

func (e *Uint6Element) Uint6() uint8 {
	return uint8(*e)
}

func (e *Uint6Element) SetUint6(v uint8) {
	if v > 0x3f {
		panic("value isn't 6 bit")
	}
	*e = Uint6Element(v)
}

func NewUint6Element() *Uint6Element {
	var elem Uint6Element
	return &elem
}

func CreateUint6Element() Element {
	var elem Uint6Element
	return &elem
}

type Uint7Element uint8

func (e *Uint7Element) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 bytes data, got %d", len(data))
	}

	*e = Uint7Element(data[0]&0x7f)
	return nil
}

func (e *Uint7Element) Encode() ([]byte, error) {
	data := make([]byte, 1, 1)
	data[0] = byte(*e)
	return data, nil
}

func (e *Uint7Element) Length() ElementLength {
	return ElementLength(7)
}

func (e *Uint7Element) Uint7() uint8 {
	return uint8(*e)
}

func (e *Uint7Element) SetUint7(v uint8) {
	if v > 0x7f {
		panic("value isn't 7 bit")
	}
	*e = Uint7Element(v)
}

func NewUint7Element() *Uint7Element {
	var elem Uint7Element
	return &elem
}

func CreateUint7Element() Element {
	var elem Uint7Element
	return &elem
}

type BytesElement []byte

func (e *BytesElement) Decode(data []byte) error {
	*e = data
	return nil
}

func (e *BytesElement) Encode() ([]byte, error) {
	return []byte(*e), nil
}

func (e *BytesElement) Length() ElementLength {
	return ElementLength(len(*e)*8)
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

func (e *Int8Element) Length() ElementLength {
	return ElementLength(8)
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

func (e *Uint8Element) Length() ElementLength {
	return ElementLength(8)
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

func (e *Int16Element) Length() ElementLength {
	return ElementLength(16)
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

func (e *Uint16Element) Length() ElementLength {
	return ElementLength(16)
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

func (e *Int32Element) Length() ElementLength {
	return ElementLength(32)
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

func (e *Uint32Element) Length() ElementLength {
	return ElementLength(32)
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

func (e *Int64Element) Length() ElementLength {
	return ElementLength(64)
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

func (e *Uint64Element) Length() ElementLength {
	return ElementLength(64)
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

func (e *BoolElement) Length() ElementLength {
	return ElementLength(8)
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
	v := string(data)
	*e = StringElement(v)
	return nil
}

func (e *StringElement) Encode() ([]byte, error) {
	data := []byte(*e)
	return data, nil
}

func (e *StringElement) Length() ElementLength {
	return ElementLength(len(*e)*8)
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

func (e *HexStringElement) Length() ElementLength {
	return ElementLength(len(*e)/2*8)
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

type Uint8Codes map[uint8]string

type Uint4CodeElement struct {
	Code uint8
	Name string
	codes Uint8Codes
}

func (e *Uint4CodeElement) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 byte data, got %d", len(data))
	}
	e.Code = uint8(data[0]&0x0f)
	e.Name = e.codes[e.Code]
	return nil
}

func (e *Uint4CodeElement) Encode() ([]byte, error) {
	return []byte{byte(e.Code)}, nil
}

func (e *Uint4CodeElement) Length() ElementLength {
	return ElementLength(4)
}

func (e *Uint4CodeElement) Uint4() uint8 {
	return e.Code
}

func (e *Uint4CodeElement) SetUint4(v uint8) {
	if v > 0x0f {
		panic("value isn't 4-bit")
	}
	e.Code = v
	e.Name = e.codes[e.Code]
}

type Uint8CodeElement struct {
	Code uint8
	Name string
	codes Uint8Codes
}

func (e *Uint8CodeElement) Decode(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Expect 1 byte data, got %d", len(data))
	}
	e.Code = uint8(data[0])
	e.Name = e.codes[e.Code]
	return nil
}

func (e *Uint8CodeElement) Encode() ([]byte, error) {
	return []byte{byte(e.Code)}, nil
}

func (e *Uint8CodeElement) Length() ElementLength {
	return ElementLength(8)
}

func (e *Uint8CodeElement) Uint8() uint8 {
	return e.Code
}

func (e *Uint8CodeElement) SetUint8(v uint8) {
	e.Code = v
	e.Name = e.codes[e.Code]
}



