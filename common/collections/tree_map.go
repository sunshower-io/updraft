package collections

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

var IllegalState = errors.New("Illegal state")

type Value interface{}

type Comparator interface {
	Compare(Value, Value) int
}

type FunctionComparator struct {
	Comparator

	Ord func(Value, Value) int
}

func (c *FunctionComparator) Compare(lhs, rhs Value) int {
	return c.Ord(lhs, rhs)
}

type TreeMap struct {
	size int

	root *node

	minimum *node
	maximum *node

	comparator Comparator
}

func NewTreeMap(f func(Value, Value) int) *TreeMap {
	return &TreeMap{
		size:    0,
		root:    nil,
		minimum: nil,
		maximum: nil,
		comparator: &FunctionComparator{
			Ord: f,
		},
	}
}

func (t *TreeMap) FirstValue() Value {
	if t.size >= 0 {
		return t.minimum.value
	}
	return nil
}

func (t *TreeMap) FirstKey() Value {
	if t.size >= 0 {
		return t.minimum.key
	}
	return nil
}

func (t *TreeMap) Iterator() Iterator {
	return Iterator{tree: t, node: t.minimum}
}

func (t *TreeMap) IsEmpty() bool {
	return t.size == 0
}

func (t *TreeMap) Size() int {
	return t.size
}

func (t *TreeMap) String() string {

	b := new(bytes.Buffer)

	write(t.root, 0, b)
	return b.String()
}

func write(n *node, depth int, b *bytes.Buffer) {
	if n != nil {
		b.WriteString(strings.Repeat(" ", depth))
		b.WriteString(fmt.Sprintf("Node{key: %s, value: %s", n.key, n.value))
		b.WriteString("\n")
		write(n.left, depth+1, b)
		write(n.right, depth+1, b)
	}
}

func (t *TreeMap) compare(lhs, rhs Value) int {
	return t.comparator.Compare(lhs, rhs)
}

func (t *TreeMap) setMinimum(n *node) {
	if t.minimum == nil {
		t.minimum = n
		t.maximum = n
	} else if t.compare(n.key, t.minimum.key) <= 0 {
		t.minimum = n
	}
}

func (t *TreeMap) setMaximum(n *node) {

	if t.maximum == nil {
		t.maximum = n
		t.minimum = n
	} else if t.compare(n.key, t.maximum.key) > 0 {
		t.maximum = n
	}
}
