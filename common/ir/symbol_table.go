package ir

type Line struct {
	Number int
}

type SymbolTable interface {
	Height() int

	Enter(string) (Symbol, error)

	Lookup(string) (Symbol, error)

	Entries(bool) []Symbol
}

type Key struct {
	Id string
}

func KeyedBy(k string) Key {
	return Key{
		Id: k,
	}
}

type Symbol interface {
	GetName() string

	GetSymbolTable() SymbolTable

	AddLine(*Line)

	GetLines() []*Line

	SetAttribute(Key, interface{})

	GetAttribute(Key) interface{}
}

type SymbolTableStack interface {
	
	Height() int

	Peek() SymbolTable

	Pop() SymbolTable

	EnterLocal(string) (Symbol, error)

	LookupLocal(string) (Symbol, error)

	Resolve(string) (Symbol, error)

	Tables() []SymbolTable
	
	String() string
}

type SymbolTableFactory interface {
	CreateStack() SymbolTableStack

	CreateSymbolTable(int) SymbolTable

	CreateSymbol(key string, table SymbolTable) Symbol
}
