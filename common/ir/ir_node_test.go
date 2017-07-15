package ir

import (
    "testing"
    "github.com/magiconair/properties/assert"
)

func TestGetParentReturnsNodeParent(t *testing.T) {
    bn := new(BaseIRNode)
    parent := new(BaseIRNode)
    bn.SetParent(parent)
    assert.Equal(t, parent, bn.GetParent())
}


func TestReplaceParentReturnsPreviousParent(t *testing.T) {
    bn := new(BaseIRNode)
    parent := new(BaseIRNode)
    bn.SetParent(parent)
    
    newParent := new(BaseIRNode)
    assert.Equal(t, parent, bn.GetParent())
    
    assert.Equal(t, parent, bn.SetParent(newParent))
    assert.Equal(t, parent, bn.GetParent())
}

func TestAddChildSetsParentCorrectly(t *testing.T) {
    
    bn := new(BaseIRNode)
    child := new(BaseIRNode)
    bn.AddChild(child)
    
    assert.Equal(t, child.GetParent(), bn)
}


func TestAddingChildIncrementsArity(t *testing.T) {
    
    bn := new(BaseIRNode)
    child := new(BaseIRNode)
    bn.AddChild(child)
   
    assert.Equal(t, bn.Arity(), 1)
}


func TestArityOnNodeWithNodeChildrenDoesNotPanic(t *testing.T) {
    
    child := new(BaseIRNode)
    
    assert.Equal(t, child.Arity(), 0)
}
