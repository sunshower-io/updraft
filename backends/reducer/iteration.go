package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)

type IterationReducer struct {
    
}


func (r IterationReducer) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    return nil
}
