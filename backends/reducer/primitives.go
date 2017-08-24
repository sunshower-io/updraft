package reducer

import (
    "github.com/sunshower-io/updraft/common/backend"
    "github.com/sunshower-io/updraft/common/ir"
)

type PrimitiveReducer struct {
    
}

func (r PrimitiveReducer) Apply(
        node ir.IntermediateNode, 
        ctx backend.OperationContext,
) interface{} {
    return node.GetValue()
}


