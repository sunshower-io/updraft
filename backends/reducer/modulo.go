package reducer

import (
    "github.com/sunshower-io/updraft/backends/common"
    "github.com/sunshower-io/updraft/common/ir"
)

type ModuloOperation struct {
    common.Operation
}

func (o ModuloOperation) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    reducer := o.Operation
    lhsNode := node.Get(0)
    rhsNode := node.Get(1)
    
    lhs := reducer.Apply(lhsNode, ctx)
    rhs := reducer.Apply(rhsNode, ctx)
    return lhs.(int64) % rhs.(int64)
}
