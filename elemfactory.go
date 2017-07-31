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
	
	f.Register("ElemTypeCodeElement", CreateElemTypeCodeElement)
	f.Register("CommonControlElement", CreateCommonControlElement)
	f.Register("ElemStatusCodeElement", CreateElemStatusCodeElement)
	f.Register("CommonStatusElement", CreateCommonStatusElement)


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
	RegisterEF("default", NewBasicElementFactory())
	RegisterEF("basic", NewBasicElementFactory())
	RegisterEF("ses", NewBasicElementFactory())
}

