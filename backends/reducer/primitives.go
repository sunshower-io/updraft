package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)

type PrimitiveReducer struct {
    
}

func (r PrimitiveReducer) Apply(
        node ir.IntermediateNode, 
        ctx common.OperationContext,
) interface{} {
    return node.GetValue()
}


