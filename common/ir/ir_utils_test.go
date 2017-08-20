package ir

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDepthOfNullNodeIsZero(t *testing.T) {
    assert.Equal(t, Depth(nil), 0)
}


func TestDepthOfRootNodeIsOne(t *testing.T) {
    assert.Equal(t, Depth(new(BaseIRNode)), 1)
}

func TestDepthOfNodeWithSingleChildIsTwo(t *testing.T) {
    parent := new(BaseIRNode)
    child := new(BaseIRNode)
    parent.AddChild(child)
    assert.Equal(t, Depth(parent), 2)
}

func TestDepthOfNodeWithTwoChildrenIsDepthOfMaxChild(t *testing.T) {
    
    parent := new(BaseIRNode)
    c1 := new(BaseIRNode)
    parent.AddChild(c1)
    
    c2 := new(BaseIRNode)
    parent.AddChild(c2)
    
    c2.AddChild(new(BaseIRNode))
    
    
    assert.Equal(t, Depth(parent), 3)
}


