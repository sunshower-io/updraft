package ir


import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSimplePathReturnsErrorForNilNode(t *testing.T) {
    
    _, er := PathBy(Index()).To("/0").Traverse(nil)
    assert.Equal(t, er, InvalidHeightError)
}

func TestSimplePathFromRootToRootFindsRoot(t *testing.T) {
    n := new(BaseIRNode)
    node , _:= PathBy(Index()).To("/").Traverse(n)
    
    assert.Equal(t, n, node)
}


func TestSimplePathFromRootToChildAtIndex0FindsChild(t *testing.T) {
    n := new(BaseIRNode)
    c := new(BaseIRNode)
    n.AddChild(c)
    node , _:= PathBy(Index()).To("/0").Traverse(n)
    
    assert.Equal(t, c, node)
}

func TestSimplePathFromRootToGrandChildOfFullTreeWorks(t *testing.T) {
    
    p := new(BaseIRNode)
    c1 := new(BaseIRNode)
    c2 := new(BaseIRNode)
    c3 := new(BaseIRNode)
    c4 := new(BaseIRNode)
    
    p.AddChild(c1)
    p.AddChild(c2)
    p.AddChild(c3)
    p.AddChild(c4)
    
    gc1 := new(BaseIRNode)
    c3.AddChild(gc1)
    
    
    n, e := PathBy(Index()).To("/2/0").Traverse(p)
    assert.Nil(t, e)
    
    assert.Equal(t, n, gc1)
    
}

