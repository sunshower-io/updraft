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
)


func init() {
    RegisterIntermediateType(NO_OP, "NoOp")
    RegisterIntermediateType(ASSIGN, "Assign")
    RegisterIntermediateType(VARIABLE, "var")
    RegisterIntermediateType(INTEGER, "int64")
    RegisterIntermediateType(FLOAT, "float64")
    RegisterIntermediateType(STRING_LITERAL, "string")
    RegisterIntermediateType(NOT, "not")
}


type BaseIRNode struct {
    IntermediateNode
    
    id          utils.Identifier
    line        int
    value       interface{}
    parent      IntermediateNode
    children    []IntermediateNode
    
    attributes  map[AttributeKey]interface{}
    
    Token       core.Token
    
    
    Type        IntermediateNodeType
    
    
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
func (n *BaseIRNode) SetLine(line int) {
    n.line = line
}

func (n *BaseIRNode) GetLine() int {
    return n.line
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

