package sespage

type ElementCreator func()Element

type ElementFactory struct {
	creators map[string]ElementCreator
}

func (f *ElementFactory) Register(name string, creator ElementCreator) {
	f.creators[name] = creator
}

func (f *ElementFactory) Unregister(name string) {
	delete(f.creators, name)
}

func (f *ElementFactory) CreaterElement(name string) Element {
	if creator, ok := f.creators[name]; ok {
		return creator()
	}
	panic("Unknown element type: " + name)
}

func (f *ElementFactory) ListAllNames() []string {
	keys := make([]string, len(f.creators))
	i := 0
	for k, _ := range f.creators {
		keys[i] = k
		i++
	}
	return keys
}

func NewElementFactory() *ElementFactory {
	ef := ElementFactory{creators: make(map[string]ElementCreator)}
	return &ef
}

func InitElementFactory(factory *ElementFactory) {
	factory.Register("BytesElement", CreateBytesElement)
	factory.Register("Int8Element", CreateInt8Element)
	factory.Register("Uint8Element", CreateUint8Element)
	factory.Register("Int16Element", CreateInt16Element)
	factory.Register("Uint16Element", CreateUint16Element)
	factory.Register("Int32Element", CreateInt32Element)
	factory.Register("Uint32Element", CreateUint32Element)
	factory.Register("Int64Element", CreateInt64Element)
	factory.Register("Uint64Element", CreateUint64Element)
	factory.Register("BoolElement", CreateBoolElement)
	factory.Register("StringElement", CreateStringElement)
	factory.Register("HexStringElement", CreateHexStringElement)
}

var (
	defaultElementFactory *ElementFactory
)

func NewDefaultElementFactory() *ElementFactory {
	return defaultElementFactory
}

func init() {
	ef := NewElementFactory()
	InitElementFactory(ef)
	defaultElementFactory = ef
}

