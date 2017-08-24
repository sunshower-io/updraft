package common

import (
    "github.com/sunshower-io/updraft/common/ir"
)

type OperationContext interface {
    
    IncrementOperations()
    
    ResolveFor(
            Operation, 
            ir.IntermediateNodeType,
    ) Operation
    
    
    Resolve(
            Operation, 
            ir.IntermediateNode,
    ) Operation
    
}




type Operation interface {
    
    Apply(
            ir.IntermediateNode, 
            OperationContext,
    ) interface{}
}