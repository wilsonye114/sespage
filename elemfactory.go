package sespage

type ElementCreator func()Element

type ElementFactory map[string]ElementCreator

func (f ElementFactory) Register(name string, creator ElementCreator) {
	f[name] = creator
}

func (f ElementFactory) Unregister(name string) {
	delete(f, name)
}

func (f ElementFactory) CreateElement(name string) Element {
	if creator, ok := f[name]; ok {
		return creator()
	}
	panic("Unknown element type: " + name)
}

func (f ElementFactory) ListAllElements() []string {
	keys := make([]string, len(f))
	i := 0
	for k, _ := range f {
		keys[i] = k
		i++
	}
	return keys
}

func NewElementFactory() ElementFactory {
	ef := make(ElementFactory)
	return ef
}

func NewBasicElementFactory() ElementFactory {
	f := NewElementFactory()
	f.Register("Uint1Element", CreateUint1Element)
	f.Register("Uint2Element", CreateUint2Element)
	f.Register("Uint3Element", CreateUint3Element)
	f.Register("Uint4Element", CreateUint4Element)
	f.Register("Uint5Element", CreateUint5Element)
	f.Register("Uint6Element", CreateUint6Element)
	f.Register("Uint7Element", CreateUint7Element)

	f.Register("BytesElement", CreateBytesElement)
	f.Register("Int8Element", CreateInt8Element)
	f.Register("Uint8Element", CreateUint8Element)
	f.Register("Int16Element", CreateInt16Element)
	f.Register("Uint16Element", CreateUint16Element)
	f.Register("Int32Element", CreateInt32Element)
	f.Register("Uint32Element", CreateUint32Element)
	f.Register("Int64Element", CreateInt64Element)
	f.Register("Uint64Element", CreateUint64Element)
	f.Register("BoolElement", CreateBoolElement)
	f.Register("StringElement", CreateStringElement)
	f.Register("HexStringElement", CreateHexStringElement)

	return f
}

func NewSesElementFactory() ElementFactory {
	f := NewElementFactory()
	// Basic Elements
	f.Register("Uint1Element", CreateUint1Element)
	f.Register("Uint2Element", CreateUint2Element)
	f.Register("Uint3Element", CreateUint3Element)
	f.Register("Uint4Element", CreateUint4Element)
	f.Register("Uint5Element", CreateUint5Element)
	f.Register("Uint6Element", CreateUint6Element)
	f.Register("Uint7Element", CreateUint7Element)

	f.Register("BytesElement", CreateBytesElement)
	f.Register("Int8Element", CreateInt8Element)
	f.Register("Uint8Element", CreateUint8Element)
	f.Register("Int16Element", CreateInt16Element)
	f.Register("Uint16Element", CreateUint16Element)
	f.Register("Int32Element", CreateInt32Element)
	f.Register("Uint32Element", CreateUint32Element)
	f.Register("Int64Element", CreateInt64Element)
	f.Register("Uint64Element", CreateUint64Element)
	f.Register("BoolElement", CreateBoolElement)
	f.Register("StringElement", CreateStringElement)
	f.Register("HexStringElement", CreateHexStringElement)
	
	// SES Elements
	f.Register("ElemTypeCodeElement", CreateElemTypeCodeElement)
	f.Register("CommonControlElement", CreateCommonControlElement)
	f.Register("ElemStatusCodeElement", CreateElemStatusCodeElement)
	f.Register("CommonStatusElement", CreateCommonStatusElement)
	f.Register("UnspecifiedControlElement", CreateUnspecifiedControlElement)
	f.Register("UnspecifiedStatusElement", CreateUnspecifiedStatusElement)

	// SES OEM Elements
	f.Register("TypeCodeElement", CreateTypeCodeElement)
	f.Register("MidplaneVpdDataElement", CreateMidplaneVpdDataElement)
	f.Register("VpdStatusDescriptorElement", CreateVpdStatusDescriptorElement)
	f.Register("CanisterVpdDataElement", CreateCanisterVpdDataElement)

	return f
}

var (
	elementFactories map[string]ElementFactory
)

func InitEFs() {
	elementFactories = make(map[string]ElementFactory)
}

func RegisterEF(name string, ef ElementFactory) {
	elementFactories[name] = ef
}

func GetEF(name string) ElementFactory {
	f, ok := elementFactories[name]
	if !ok {
		panic("Unknown factory: " + name)
	}
	return f
}

func init() {
	InitEFs()
	RegisterEF("default", NewSesElementFactory())
	RegisterEF("basic", NewBasicElementFactory())
	RegisterEF("ses", NewSesElementFactory())
}

