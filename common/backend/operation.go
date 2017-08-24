package backend

import "github.com/sunshower-io/updraft/common/ir"

type OperationContext interface {
    IncrementOperations()
    
}


type Operation interface {
    Apply(
            ir.IntermediateNode, 
            OperationContext,
    ) interface{}
}


