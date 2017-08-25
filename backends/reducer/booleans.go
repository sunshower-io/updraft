package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)

type NotOperation struct {
    common.Operation
}


func (o NotOperation) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    reducer := o.Operation
    exprNode := node.Get(0)
    expr := reducer.Apply(exprNode, ctx).(bool)
    return !expr
}

type BooleanOperation struct {
    common.Operation
}


func (b BooleanOperation) op(bool, bool) bool {
    return false
}

func (o BooleanOperation) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    reducer := o.Operation
    lhsNode := node.Get(0)
    rhsNode := node.Get(1)
    
    lhs := reducer.Apply(lhsNode, ctx).(bool)
    rhs := reducer.Apply(rhsNode, ctx).(bool)
    
    return o.op(lhs, rhs)
}



type AndOperation struct {
    BooleanOperation
}

func (o AndOperation) op(lhs, rhs bool) bool {
    return lhs && rhs
}


type OrOperation struct {
    BooleanOperation
}


func (o OrOperation) op(lhs, rhs bool) bool {
    return lhs || rhs
}


type XorOperation struct {
    BooleanOperation
}


func (o XorOperation) op(lhs, rhs bool) bool {
    return lhs != rhs
}



