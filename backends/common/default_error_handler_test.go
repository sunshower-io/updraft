package common 

import (
    "testing"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/stretchr/testify/assert"
)

func TestResolveLineReturnsCorrectLineForNode(t *testing.T) {
    
    node := new(ir.BaseIRNode)
    node.SetLine("hello")
    
    assert.Equal(t, resolveLine(node), "hello")
}

func TestResolveLineReturnsCorrectLineForParent(t *testing.T) {
    
    child := new(ir.BaseIRNode)
    parent := new(ir.BaseIRNode)
    
    parent.SetLine("frapper")
    parent.AddChild(child)
    
    assert.Equal(t, resolveLine(child), "frapper")
}


func TestResolveLineResolvesCorrectLineForGrandParent(t *testing.T) {
    
    child := new(ir.BaseIRNode)
    parent := new(ir.BaseIRNode)
    gparent := new(ir.BaseIRNode)
    
    gparent.SetLine("frapper")
    gparent.AddChild(parent)
    
    parent.AddChild(child)
    
    assert.Equal(t, resolveLine(child), "frapper")
}