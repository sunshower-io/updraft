package ir

var DefaultSymbolTableFactory SymbolTableFactory = new(BaseSymbolTableFactory)

type BaseSymbolTableFactory struct {
	SymbolTableFactory
}

func (b *BaseSymbolTableFactory) CreateStack() SymbolTableStack {
	stack := &BaseSymbolTableStack{
		height:  0,
		factory: b,
	}
	stack.initialize()
	return stack
}

func (b *BaseSymbolTableFactory) CreateSymbolTable(h int) SymbolTable {
	return NewSymbolTable(b, h)
}

func (b *BaseSymbolTableFactory) CreateSymbol(key string, table SymbolTable) Symbol {
	return &BaseSymbol{
		name: key,
	}
}

