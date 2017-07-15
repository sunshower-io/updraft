package ir

type BaseSymbol struct {
	Symbol
	name  string
	lines []*Line
	table SymbolTable

	attributes map[string]interface{}
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
	s.attributes[k.Id] = value
}

func (s *BaseSymbol) GetAttribute(key Key) interface{} {
	return s.attributes[key.Id]
}
