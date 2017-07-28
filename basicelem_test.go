package sespage

import (
	"testing"
	"log"
	"math/rand"
)

func CommonTestField(bytes []byte, element Element) {
	log.Printf("Test %T", element)
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	dbytes, err := element.Encode()
	if err != nil {
		panic(err)
	}
	length := int(element.Length())
	log.Printf("Decode %v, Encode %v, length: %d\n", bytes, dbytes, length)
	for i := 0; i < length; i++ {
		if dbytes[i] != bytes[i] {
			panic("Encode/Decode error")
		}
	}
}

func CommonTestIntField(element Element) {
	var value int64
	var evalue int64

	bytes := make([]byte, 8, 8)
	for i := 0; i < 8; i++ {
		bytes[i] = byte(rand.Int31())
	}
	CommonTestField(bytes, element)

	switch v := element.(type) {
	case Int8Field:
		tmpvalue := int8(rand.Int31())
		v.SetInt8(tmpvalue)
		evalue = int64(v.Int8())
		value = int64(tmpvalue)
	case Uint8Field:
		tmpvalue := uint8(rand.Int31())
		v.SetUint8(tmpvalue)
		evalue = int64(v.Uint8())
		value = int64(tmpvalue)
	case Int16Field:
		tmpvalue := int16(rand.Int31())
		v.SetInt16(tmpvalue)
		evalue = int64(v.Int16())
		value = int64(tmpvalue)
	case Uint16Field:
		tmpvalue := uint16(rand.Int31())
		v.SetUint16(tmpvalue)
		evalue = int64(v.Uint16())
		value = int64(tmpvalue)
	case Int32Field:
		tmpvalue := rand.Int31()
		v.SetInt32(tmpvalue)
		evalue = int64(v.Int32())
		value = int64(tmpvalue)
	case Uint32Field:
		tmpvalue := uint32(rand.Int31())
		v.SetUint32(tmpvalue)
		evalue = int64(v.Uint32())
		value = int64(tmpvalue)
	case Int64Field:
		tmpvalue := rand.Int63()
		v.SetInt64(tmpvalue)
		evalue = v.Int64()
		value = tmpvalue
	case Uint64Field:
		tmpvalue := uint64(rand.Int63())
		v.SetUint64(tmpvalue)
		evalue = int64(v.Uint64())
		value = int64(tmpvalue)
	}
	log.Printf("Set %d, Get %d\n", value, evalue)
	if value != evalue {
		panic("Get/Set error")
	}
}

func TestIntElements(t *testing.T) {
	var int8elem Int8Element
	var uint8elem Uint8Element
	var int16elem Int16Element
	var uint16elem Uint16Element
	var int32elem Int32Element
	var uint32elem Uint32Element
	var int64elem Int64Element
	var uint64elem Uint64Element

	CommonTestIntField(&int8elem)
	CommonTestIntField(&uint8elem)
	CommonTestIntField(&int16elem)
	CommonTestIntField(&uint16elem)
	CommonTestIntField(&int32elem)
	CommonTestIntField(&uint32elem)
	CommonTestIntField(&int64elem)
	CommonTestIntField(&uint64elem)
}

func TestBoolElement(t *testing.T) {
	var boolelem BoolElement

	bytes := []byte{0,1}
	CommonTestField(bytes, &boolelem)
	for _, v := range []bool{true, false} {
		boolelem.SetBool(v)
		ev := boolelem.Bool()
		log.Printf("Set %v, Get %v\n", v, ev)
		if v != ev {
			panic("Get/Set error")
		}
	}
}

func TestStringElement(t *testing.T) {
	var strelem StringElement
	bytes := []byte("hello world")
	CommonTestField(bytes, &strelem)
	v := "hi yourself"
	strelem.SetString(v)
	s := strelem.String()
	log.Printf("Set %s, Get %s\n", v, s)
	if v != s {
		panic("Get/Set error")
	}
}

func TestHexStringElement(t *testing.T) {
	var hexelem HexStringElement
	bytes1 := []byte{0xff}
	bytes2 := []byte{0xff, 0xee, 0x11, 0x04, 0x05, 0x06, 0x07}
	CommonTestField(bytes1, &hexelem)
	CommonTestField(bytes2, &hexelem)
	v := "ffee1104050607"
	ev := hexelem.String()
	log.Printf("Get %s\n", ev)
	if v != ev {
		panic("Get/Set error")
	}
	v = "ef1467e81ac7"
	hexelem.SetString(v)
	ev = hexelem.String()
	log.Printf("Set %s, Get %s\n", v, ev)
	if v != ev {
		panic("Get/Set error")
	}
}

func TestBytesElement(t *testing.T) {
	var byteselem BytesElement
	bytes1 := []byte{0x05}
	bytes2 := []byte{0xff, 0xcc, 0x14, 0x55, 0x98}
	CommonTestField(bytes1, &byteselem)
	CommonTestField(bytes2, &byteselem)
	bytes3 := []byte{0xee, 0xca, 0x14, 0x67, 0x23, 0x78}
	byteselem.SetBytes(bytes3)
	bytes4 := byteselem.Bytes()
	log.Printf("Set %v, Get %v\n", bytes3, bytes4)
	if len(bytes3) != len(bytes4) {
		panic("Get/Set Error")
	}
	for i, _ := range bytes4 {
		if bytes3[i] != bytes4[i] {
			panic("Get/Set Error")
		}
	}
}
