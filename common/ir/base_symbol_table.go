package ir

import (
	"errors"
	"fmt"
	"gitlab.com/sunshower.io/updraft/common/collections"
	"strings"
)

func stringComparator(a collections.Value, b collections.Value) int {
	return strings.Compare(a.(string), b.(string))
}

var NoSuchSymbolTable = errors.New("No table at that level exists")

type BaseSymbolTableStack struct {
	SymbolTableStack

	height int

	stack []SymbolTable

	factory SymbolTableFactory
}

func (s *BaseSymbolTableStack) Tables() []SymbolTable {
	return s.stack[:]
}

func (s *BaseSymbolTableStack) initialize() {
	s.stack = append(s.stack, NewSymbolTable(s.factory, s.height))
}

func (s *BaseSymbolTableStack) get(l int) (SymbolTable, error) {
	if l < len(s.stack) {
		st := s.stack[s.height]
		if st != nil {
			return st, nil
		}
	}
	return nil, NoSuchSymbolTable
}


func(s *BaseSymbolTableStack) Peek() SymbolTable {
	return nil
}

func(s *BaseSymbolTableStack) Pop() SymbolTable {
	return nil
}

func (s *BaseSymbolTableStack) Height() int {
	return s.height
}

func (s *BaseSymbolTableStack) EnterLocal(name string) (Symbol, error) {
	st, er := s.get(s.height)
	if er == nil {
		return st.Enter(name)
	}
	return nil, NoSuchSymbolTable
}

func (s *BaseSymbolTableStack) Resolve(name string) (Symbol, error) {
	return s.LookupLocal(name)
}

func (s *BaseSymbolTableStack) LookupLocal(name string) (Symbol, error) {
	st, er := s.get(s.height)
	if er == nil {
		return st.Lookup(name)
	}
	return nil, NoSuchSymbolTable
}


func (t *BaseSymbolTableStack) String() string {
	return new(CrossReferencer).Print(t)
}

type BaseSymbolTable struct {
	SymbolTable
	level int

	factory SymbolTableFactory
	values  *collections.TreeMap
}

func (t *BaseSymbolTable) Height() int {
	return t.level
}

func (t *BaseSymbolTable) Lookup(key string) (Symbol, error) {
	sym := t.values.Get(key)
	if sym != nil {
		return sym.(Symbol), nil
	}
	return nil, errors.New(fmt.Sprintf("No symbol named '%s'", key))
}

func (t *BaseSymbolTable) Enter(key string) (Symbol, error) {
	sym := t.factory.CreateSymbol(key, t)
	t.values.Put(key, sym)
	return sym, nil
}

func (t *BaseSymbolTable) Entries(bool) []Symbol {
	entries := make([]Symbol, 0)
	for iter := t.values.Iterator(); iter.HasNext(); iter = iter.Next() {
		entries = append(entries, iter.NextValue().(Symbol))
	}
	return entries
}

func (t *BaseSymbolTable) EnterLocal(key string) (Symbol, error) {
	sym := t.factory.CreateSymbol(key, t)
	t.values.Put(key, sym)
	return sym, nil
}

func (t *BaseSymbolTable) LookupLocal(key string) (Symbol, error) {
	sym := t.factory.CreateSymbol(key, t)
	t.values.Put(key, sym)
	return sym, nil
}


func NewSymbolTable(factory SymbolTableFactory, level int) SymbolTable {
	return &BaseSymbolTable{
		level:   level,
		factory: factory,
		values:  collections.NewTreeMap(stringComparator),
	}
}


