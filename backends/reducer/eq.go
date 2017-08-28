package reducer

import (
    "github.com/sunshower-io/updraft/backends/common"
    "github.com/sunshower-io/updraft/common/ir"
)

type EqualityReducer struct {
    common.Operation
}

func (r EqualityReducer) Apply(
    node ir.IntermediateNode, 
    ctx common.OperationContext,
) interface{} {
    
    reducer := r.Operation
    lhsNode := node.Get(0)
    rhsNode := node.Get(1)
    
    lhs := reducer.Apply(lhsNode, ctx)
    rhs := reducer.Apply(rhsNode, ctx)
    return lhs == rhs
}



type InequalityReducer struct {
    common.Operation
}

func (r InequalityReducer) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    
    reducer := r.Operation
    lhsNode := node.Get(0)
    rhsNode := node.Get(1)
    
    lhs := reducer.Apply(lhsNode, ctx)
    rhs := reducer.Apply(rhsNode, ctx)
    return lhs != rhs
}
