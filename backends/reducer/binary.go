package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)




type AddOperation struct {
    StatementReducer
}

func (o AddOperation) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    reducer := o.StatementReducer
    
    lhsNode := node.Get(0)
    rhsNode := node.Get(1)
    
    lhs := reducer.Apply(lhsNode, ctx)
    rhs := reducer.Apply(rhsNode, ctx)
    
   
    if isFloat(lhs) || isFloat(rhs) {
        return lhs.(float64) + rhs.(float64)
    }
    
    return lhs.(int64) + rhs.(int64)
}


func isInt(lhs interface{}) bool {
    switch lhs.(type) {
    case int, int64:
        return true
    }
    return false
}

func isFloat(lhs interface{}) bool {
    switch lhs.(type) {
    case float32, float64:
        return true
    }
    return false
}

