package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)

type CompoundReducer struct {
    common.Operation

     
}

func (c CompoundReducer) Apply(
        node ir.IntermediateNode, 
        ctx common.OperationContext,
) interface{} {
    if children := node.GetChildren(); children != nil {
        for _, child := range children {
            c.Operation.Apply(child, ctx)
        }
    }
    return nil
}