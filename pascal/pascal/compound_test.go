package pascal

import (
    "testing"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/magiconair/properties/assert"
)

func TestNestedEmptyNestedCompoundStatementsWorks(t *testing.T) {
    prg := `
    
    BEGIN
        BEGIN
            helloWorld := 10 + 20;        
        
        END;
    
    END.
    
`
    model := printTree(prg)
    
    node, _ := ir.PathBy(ir.Index()).To("/0/0/0").Traverse(model.GetRoot())
    
    assert.Equal(t, node.GetType(), ir.VARIABLE)
    
}
