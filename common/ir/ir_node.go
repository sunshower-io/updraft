package ir

import (
    "github.com/sunshower-io/updraft/common/utils"
    "github.com/sunshower-io/updraft/middle/core"
)


const (
    
    NO_OP               IntermediateNodeType = -1 
    ASSIGN              IntermediateNodeType = 0
    VARIABLE            IntermediateNodeType = 1
    INTEGER             IntermediateNodeType = 2
    FLOAT               IntermediateNodeType = 3
    STRING_LITERAL      IntermediateNodeType = 4
    NOT                 IntermediateNodeType = 5
    EXPRESSION          IntermediateNodeType = 6
    ADD                 IntermediateNodeType = 7
    SUBTRACT            IntermediateNodeType = 8
    MULTIPLY            IntermediateNodeType = 9
    DIVIDE              IntermediateNodeType = 10
    SCOPE               IntermediateNodeType = 11
    NEGATE              IntermediateNodeType = 12
    
    INTEGER_DIVIDE      IntermediateNodeType = 13
    FLOAT_DIVIDE        IntermediateNodeType = 14
    OR                  IntermediateNodeType = 15
    
    GTE                 IntermediateNodeType = 16
    LTE                 IntermediateNodeType = 17
    LOOP                IntermediateNodeType = 18
    TEST                IntermediateNodeType = 19
    EQUAL_TO            IntermediateNodeType = 20
    NOT_EQUAL_TO        IntermediateNodeType = 21
    LESS_THAN           IntermediateNodeType = 22
    GREATER_THAN        IntermediateNodeType = 23
    MODULO              IntermediateNodeType = 24
    BOOLEAN             IntermediateNodeType = 25
    AND                 IntermediateNodeType = 26
)


func init() {
    RegisterIntermediateType(NO_OP, "no-op")
    RegisterIntermediateType(ASSIGN, "assign")
    RegisterIntermediateType(VARIABLE, "var")
    RegisterIntermediateType(INTEGER, "int64")
    RegisterIntermediateType(FLOAT, "float64")
    RegisterIntermediateType(STRING_LITERAL, "string")
    RegisterIntermediateType(NOT, "not")
    RegisterIntermediateType(EXPRESSION, "expr")
    
    
    RegisterIntermediateType(SCOPE, "scope")
    RegisterIntermediateType(LOOP, "loop")
    RegisterIntermediateType(TEST, "test")
    RegisterIntermediateType(ADD, "+")
    RegisterIntermediateType(SUBTRACT, "-")
    RegisterIntermediateType(MULTIPLY, "*")
    RegisterIntermediateType(FLOAT_DIVIDE, "/")
    RegisterIntermediateType(INTEGER_DIVIDE, "/")
    RegisterIntermediateType(DIVIDE, "/")
    RegisterIntermediateType(MODULO, "modulo")
    RegisterIntermediateType(NEGATE, "negate")
}


type BaseIRNode struct {
    IntermediateNode
    
    id          utils.Identifier
    line        string 
    lineNumber  int
    value       interface{}
    parent      IntermediateNode
    children    []IntermediateNode
    
    attributes  map[AttributeKey]interface{}
    
    Token       core.Token
    
    
    Type        IntermediateNodeType
    
    
}

func (n *BaseIRNode) Get(i int) IntermediateNode {
    return n.children[i]
}


func (n *BaseIRNode) GetToken() core.Token {
    return n.Token
}

func (n *BaseIRNode) ChildAt(idx int) IntermediateNode {
    
    if !(n.children == nil || idx >= len(n.children)) {
        return n.children[idx]
    }
    return nil
}


/**
    Mutators for ID
 */

func (n *BaseIRNode) SetId(
        identifier utils.Identifier,
) {
    n.id = identifier
}



func (n *BaseIRNode) GetId() utils.Identifier {
    return n.id
}


/**
    Mutators for line
 */
func (n *BaseIRNode) SetLineNumber(line int) {
    n.lineNumber = line
}

func (n *BaseIRNode) GetLineNumber() int {
    return n.lineNumber
}


func (n *BaseIRNode) GetLine() string {
    return n.line
}

func (n *BaseIRNode) SetLine(line string) {
    n.line = line
}


/**
    Mutators for value
 */
func (n *BaseIRNode) SetValue(value interface{}) {
    n.value = value
}


func (n *BaseIRNode) GetValue() interface{} {
    return n.value
}



/**
    Mutators for parent
 */

func (n *BaseIRNode) GetParent() IntermediateNode {
    return n.parent
}

func (n *BaseIRNode) SetParent(
        parent IntermediateNode,
) IntermediateNode {
    previous := n.parent
    n.parent = parent
    return previous
}


func (n *BaseIRNode) GetType() IntermediateNodeType {
    return n.Type
}

func (n *BaseIRNode) AddChild(child IntermediateNode) {
    if child == nil {
        panic("Don't pass me nil kiddos")
    }
    child.SetParent(n)
    n.children = append(n.children, child)
}

func(n *BaseIRNode) GetChildren() []IntermediateNode {
    return n.children[:]
}

func(n *BaseIRNode) GetAttribute(key AttributeKey) interface{} {
    if n.attributes != nil {
        return n.attributes[key]
    }
    return nil
}

func (n *BaseIRNode) SetAttribute(
        key AttributeKey,
        val interface{},
) interface{}  {
   
    if val == nil {
        panic("Use ClearAttribute() to clear attributes.  Cannot set null value")
    }
   
    return n.doSet(key, val)
}




func (n *BaseIRNode) ClearAttribute(
        key AttributeKey,
) interface{} {
    return n.doSet(key, nil)
}


func (n *BaseIRNode) Clone() IntermediateNode {
    
    cloneAttrs := make(map[AttributeKey]interface{})
    
    for key, val := range n.attributes {
        cloneAttrs[key] = val
    }
    
    clone := &BaseIRNode{
        Type: n.Type,
        attributes: cloneAttrs,
    }
    
    return clone
}

func (n *BaseIRNode) Arity() int {
    return len(n.children)
}


func (n *BaseIRNode) doSet(
        key AttributeKey,
        val interface{},
) interface{} {
    
    if n.attributes == nil {
        n.attributes = make(map[AttributeKey]interface{})
    }
    
    previous, exists := n.attributes[key]
    
    n.attributes[key] = val
    
    if exists {
        return previous
    }
    return nil
}

