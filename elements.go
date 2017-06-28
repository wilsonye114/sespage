package sespage

// import (
// 	"fmt"
// )

type Element interface {
	Decode(buf []byte, page map[string]interface{}) int
}

type Page interface {
	DecodeHead(buf []byte, map[string]interface{}) int
	Decode(buf []byte, page map[string]interface{}, context map[string]interface{}) int
}

type ByteToInt32Element struct {
	Name string
	Length int
}

func (e IntByteElement) Decode(buf []byte, page map[string]interface{}) int {
	page[e.Name] = int32(buf[0])
	return e.Length
}

func NewByteToInt32Element(name string) ByteToInt32Element {
	e := ByteToInt32Element{
		Name:name,
		Length:1}
	return e
}

type Byte2ToInt32Element struct {
	name string
	Length int
}

func (e *Int2BytesElement) Name() string {
	return e.name
}

func (e *Int2BytesElement) Decode(buf []byte) (interface{}, int) {
	e.Length = 2
	res := (int32(buf[0]) << 8) | int32(buf[1]) 
	return res, e.Length
}


type Byte4ToInt32Element struct {
	name string
	Length int
}

func (e *Int4BytesElement) Name() string {
	return e.name
}

func (e *Int4BytesElement) Decode(buf []byte) (interface{}, int) {
	e.Length = 4
	res := (int32(buf[0]) << 24) | (int32(buf[1]) << 16) | (int32(buf[2]) << 8) | int32(buf[3])
	return res, e.Length
}


