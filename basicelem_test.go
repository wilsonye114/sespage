package sespage

import (
	"testing"
	"log"
	"math/rand"
	// "time"
)

func TestElementLength(t *testing.T) {
	l := ElementLength(39)
	log.Printf("bits: %d, byte: %d, RemainderBit: %d, IsAligned: %v\n", l.Bit(), l.Byte(), l.RemainderBit(), l.IsAligned())
}

func FieldCommonTest(bytes []byte, element Element) {
	log.Printf("Test %T", element)
	err := element.Decode(bytes)
	if err != nil {
		panic(err)
	}
	dbytes, err := element.Encode()
	if err != nil {
		panic(err)
	}
	length := element.Length().Byte()
	if element.Length().IsAligned() == false {
		length++
	}
	log.Printf("Decode %v, Encode %v, Check Byte: %d, length: %dB.%db\n", bytes, dbytes, length, element.Length().Byte(), element.Length().RemainderBit())
	for i := uint64(0); i < length; i++ {
		if dbytes[i] != bytes[i] {
			log.Printf("byte[%d] : %d, %d\n", i, dbytes[i], bytes[i])
			panic("Encode/Decode error")
		}
	}
}

func IntFieldCommonTest(element Element) {
	var value int64
	var evalue int64

	bytes := make([]byte, 8, 8)
	for i := 0; i < 8; i++ {
		bytes[i] = byte(rand.Int31())
	}
	FieldCommonTest(bytes, element)

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

func BitFieldCommonTest(element Element) {
	var value uint8
	var evalue uint8

	data := make([]byte, 0x80, 0x80)
	for i := 0; i <= 0x7f; i++ {
		data[i] = byte(i)
	}

	switch v := element.(type) {
	case Uint1Field:
		for i := 0; i <= 0x1; i++ {
			FieldCommonTest(data[i:], element)
			tmpvalue := uint8(data[i])
			v.SetUint1(tmpvalue)
			evalue = v.Uint1()
			value = tmpvalue
			log.Printf("Set %d, Get %d\n", value, evalue)
			if value != evalue {
				panic("Get/Set error")
			}
		}
	case Uint2Field:
		for i := 0; i <= 0x3; i++ {
			FieldCommonTest(data[i:], element)
			tmpvalue := uint8(data[i])
			v.SetUint2(tmpvalue)
			evalue = v.Uint2()
			value = tmpvalue
			log.Printf("Set %d, Get %d\n", value, evalue)
			if value != evalue {
				panic("Get/Set error")
			}
		}
	case Uint3Field:
		for i := 0; i <= 0x7; i++ {
			FieldCommonTest(data[i:], element)
			tmpvalue := uint8(data[i])
			v.SetUint3(tmpvalue)
			evalue = v.Uint3()
			value = tmpvalue
			log.Printf("Set %d, Get %d\n", value, evalue)
			if value != evalue {
				panic("Get/Set error")
			}
		}
	case Uint4Field:
		for i := 0; i <= 0xf; i++ {
			FieldCommonTest(data[i:], element)
			tmpvalue := uint8(data[i])
			v.SetUint4(tmpvalue)
			evalue = v.Uint4()
			value = tmpvalue
			log.Printf("Set %d, Get %d\n", value, evalue)
			if value != evalue {
				panic("Get/Set error")
			}
		}
	case Uint5Field:
		for i := 0; i <= 0x1f; i++ {
			FieldCommonTest(data[i:], element)
			tmpvalue := uint8(data[i])
			v.SetUint5(tmpvalue)
			evalue = v.Uint5()
			value = tmpvalue
			log.Printf("Set %d, Get %d\n", value, evalue)
			if value != evalue {
				panic("Get/Set error")
			}
		}
	case Uint6Field:
		for i := 0; i <= 0x3f; i++ {
			FieldCommonTest(data[i:], element)
			tmpvalue := uint8(data[i])
			v.SetUint6(tmpvalue)
			evalue = v.Uint6()
			value = tmpvalue
			log.Printf("Set %d, Get %d\n", value, evalue)
			if value != evalue {
				panic("Get/Set error")
			}
		}
	case Uint7Field:
		for i := 0; i <= 0x7f; i++ {
			FieldCommonTest(data[i:], element)
			tmpvalue := uint8(data[i])
			v.SetUint7(tmpvalue)
			evalue = v.Uint7()
			value = tmpvalue
			log.Printf("Set %d, Get %d\n", value, evalue)
			if value != evalue {
				panic("Get/Set error")
			}
		}
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

	IntFieldCommonTest(&int8elem)
	IntFieldCommonTest(&uint8elem)
	IntFieldCommonTest(&int16elem)
	IntFieldCommonTest(&uint16elem)
	IntFieldCommonTest(&int32elem)
	IntFieldCommonTest(&uint32elem)
	IntFieldCommonTest(&int64elem)
	IntFieldCommonTest(&uint64elem)
}


func TestBitElements(t *testing.T) {
	var uint1elem Uint1Element
	var uint2elem Uint2Element
	var uint3elem Uint3Element
	var uint4elem Uint4Element
	var uint5elem Uint5Element
	var uint6elem Uint6Element
	var uint7elem Uint7Element

	BitFieldCommonTest(&uint1elem)
	BitFieldCommonTest(&uint2elem)
	BitFieldCommonTest(&uint3elem)
	BitFieldCommonTest(&uint4elem)
	BitFieldCommonTest(&uint5elem)
	BitFieldCommonTest(&uint6elem)
	BitFieldCommonTest(&uint7elem)

}

func TestBoolElement(t *testing.T) {
	var boolelem BoolElement

	bytes := []byte{0,1}
	FieldCommonTest(bytes, &boolelem)
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
	FieldCommonTest(bytes, &strelem)
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
	FieldCommonTest(bytes1, &hexelem)
	FieldCommonTest(bytes2, &hexelem)
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
	FieldCommonTest(bytes1, &byteselem)
	FieldCommonTest(bytes2, &byteselem)
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

func Uint8CodeElementCommonTest(element *Uint8CodeElement) {
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

func TestUint4CodeElement(t *testing.T) {

	codes := make(Uint8Codes)
	codes[0x00] = "Unsupported"
	codes[0x01] = "OK"
	codes[0x02] = "Critical"
	codes[0x03] = "Noncritical"
	codes[0x04] = "Unrecoverable"
	codes[0x05] = "Not Installed"
	codes[0x06] = "Unknown"
	codes[0x07] = "Not Available"
	codes[0x08] = "No Access Allowed"
	for i := uint8(0x09); (i >= 0x09) && (i <= 0x0f); i++ {
		codes[i] = "Reserved"
	}
	data := []byte{1,2,3,4}
	elem := Uint4CodeElement{codes: codes}
	FieldCommonTest(data[3:], &elem)
	log.Printf("%d: %s\n", elem.Code, elem.Name)
	BitFieldCommonTest(&elem)
}

func TestUint8CodeElement(t *testing.T) {

	codes := make(Uint8Codes)
	codes[0x00] = "Unsupported"
	codes[0x01] = "OK"
	codes[0x02] = "Critical"
	codes[0x03] = "Noncritical"
	codes[0x04] = "Unrecoverable"
	codes[0x05] = "Not Installed"
	codes[0x06] = "Unknown"
	codes[0x07] = "Not Available"
	codes[0x08] = "No Access Allowed"
	for i := uint8(0x09); (i >= 0x09) && (i <= 0x0f); i++ {
		codes[i] = "Reserved"
	}
	data := []byte{1,2,3,4}
	elem := Uint8CodeElement{codes: codes}
	FieldCommonTest(data[3:], &elem)
	log.Printf("%d: %s\n", elem.Code, elem.Name)
	IntFieldCommonTest(&elem)
}