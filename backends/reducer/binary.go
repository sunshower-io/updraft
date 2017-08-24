package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/common/backend"
)




type AddOperation struct {
    StatementReducer
}

func (o AddOperation) Apply(
        node ir.IntermediateNode,
        ctx backend.OperationContext,
) interface{} {
    reducer := o.StatementReducer
    
    lhsNode := node.Get(0)
    rhsNode := node.Get(1)
    
    lhs := reducer.Apply(lhsNode, ctx)
    rhs := reducer.Apply(rhsNode, ctx)
    
   
    if isFloat(lhs) || isFloat(rhs) {
        return float64(lhs) + float64(rhs)
    }
    
    return int64(lhs) + int64(rhs)
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

