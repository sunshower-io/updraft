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




type AndOperation struct {
    common.Operation 
}

func (o AndOperation) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    lhs, rhs := apply(o.Operation, node, ctx)
    return lhs && rhs
}


type OrOperation struct {
    common.Operation
}

func (o OrOperation) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    lhs, rhs := apply(o.Operation, node, ctx)
    return lhs || rhs
}


/**
    Xor Operation:
    Computes the logical exclusive-or of two boolean expressions
 */
type XorOperation struct {
    common.Operation
}


func (o XorOperation) Apply(
        node ir.IntermediateNode, 
        ctx common.OperationContext,
) interface{} {
    lhs, rhs := apply(o.Operation, node, ctx)
    return lhs != rhs
}



func apply(
        reducer common.Operation, 
        node ir.IntermediateNode,
        ctx common.OperationContext,
) (bool, bool) {
    
    lhsNode := node.Get(0)
    rhsNode := node.Get(1)
    
    lhs := reducer.Apply(lhsNode, ctx)
    rhs := reducer.Apply(rhsNode, ctx)
    return lhs.(bool), rhs.(bool)
}

