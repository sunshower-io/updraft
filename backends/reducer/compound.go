package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)

type CompoundReducer struct {
    StatementReducer
}

func (c CompoundReducer) Apply(
        node ir.IntermediateNode, 
        ctx common.OperationContext,
) interface{} {
    if children := node.GetChildren(); children != nil {
        for _, child := range children {
            c.StatementReducer.Apply(child, ctx)
        }
    }
    return nil
}