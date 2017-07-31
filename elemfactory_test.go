package sespage

import (
	"testing"
	"log"
)

func TestElementCreator(t *testing.T) {
	ef := NewBasicElementFactory()
	names1 := ef.ListAllElements()
	log.Printf("%v\n", names1)
	count1 := len(names1)
	ef.Unregister(names1[0])
	names2 := ef.ListAllElements()
	log.Printf("%v\n", names2)
	count2 := len(names2)
	if count1 != count2+1 {
		panic("Unregister Error")
	}

	ef = NewBasicElementFactory()
	names3 := ef.ListAllElements()
	log.Printf("%v\n", names3)
	count3 := len(names3)
	if count1 != count3 {
		panic("Unregister Error")
	}	
	elem := ef.CreateElement("Int64Element")
	log.Printf("%T\n", elem)
	i64elem := elem.(*Int64Element)
	log.Printf("%T\n", i64elem)
}

func TestDefaultElementCreator(t *testing.T) {
	ef := GetEF("Default")
	names1 := ef.ListAllElements()
	log.Printf("%v\n", names1)
	count1 := len(names1)
	ef.Unregister(names1[0])
	names2 := ef.ListAllElements()
	log.Printf("%v\n", names2)
	count2 := len(names2)
	if count1 != count2+1 {
		panic("Unregister Error")
	}
}

