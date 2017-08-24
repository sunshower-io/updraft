package common

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/common/backend"
)

type OperationContext interface {
    
    IncrementOperations()
    
    ResolveFor(
            backend.Operation, 
            ir.IntermediateNodeType,
    ) backend.Operation
    
    
    Resolve(
            backend.Operation, 
            ir.IntermediateNode,
    ) backend.Operation
    
}




type Operation interface {
    
    Apply(
            ir.IntermediateNode, 
            OperationContext,
    ) interface{}
}