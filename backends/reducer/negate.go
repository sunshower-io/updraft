package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)

type NegationOperation struct {
    common.Operation
}

func (n NegationOperation) Apply(
        node ir.IntermediateNode, 
        ctx common.OperationContext,
) interface{} {
    
    switch node.GetType() {
    case ir.FLOAT:
        return -node.GetValue().(float64)
    case ir.INTEGER:
        return -node.GetValue().(int64)
    }
    
    child := node.Get(0)
    value := n.Operation.Apply(child, ctx)
    switch value.(type) {
    case int64:
        return -value.(int64)
    case float64:
        return -value.(float64)
    }
    panic("Did not see that coming")
    
}
