package sespage

import (
	"testing"
	"log"
)

func TestElementCreator(t *testing.T) {
	ef := NewElementFactory()
	InitElementFactory(ef)
	names1 := ef.ListAllNames()
	log.Printf("%v\n", names1)
	count1 := len(names1)
	ef.Unregister(names1[0])
	names2 := ef.ListAllNames()
	log.Printf("%v\n", names2)
	count2 := len(names2)
	if count1 != count2+1 {
		panic("Unregister Error")
	}

	InitElementFactory(ef)
	names3 := ef.ListAllNames()
	log.Printf("%v\n", names3)
	count3 := len(names3)
	if count1 != count3 {
		panic("Unregister Error")
	}	
	elem := ef.CreaterElement("Int64Element")
	log.Printf("%T\n", elem)
	i64elem := elem.(*Int64Element)
	log.Printf("%T\n", i64elem)
}

func TestDefaultElementCreator(t *testing.T) {
	ef := NewDefaultElementFactory()
	InitElementFactory(ef)
	names1 := ef.ListAllNames()
	log.Printf("%v\n", names1)
	count1 := len(names1)
	ef.Unregister(names1[0])
	names2 := ef.ListAllNames()
	log.Printf("%v\n", names2)
	count2 := len(names2)
	if count1 != count2+1 {
		panic("Unregister Error")
	}
}

