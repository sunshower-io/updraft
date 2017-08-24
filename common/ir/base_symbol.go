package ir

type BaseSymbol struct {
	Symbol
	name  string
	lines []*Line
	table SymbolTable

	attributes map[Key]interface{}
}

func (s *BaseSymbol) GetName() string {
	return s.name
}

func (s *BaseSymbol) GetSymbolTable() SymbolTable {
	return s.table
}

func (s *BaseSymbol) AddLine(line *Line) {
	s.lines = append(s.lines, line)
}

func (s *BaseSymbol) GetLines() []*Line {
	return s.lines
}

func (s *BaseSymbol) SetAttribute(k Key, value interface{}) {
    if s.attributes == nil {
        s.attributes = make(map[Key]interface{})
    }
	s.attributes[k] = value
}

func (s *BaseSymbol) GetAttribute(key Key) interface{} {
	return s.attributes[key]
}
