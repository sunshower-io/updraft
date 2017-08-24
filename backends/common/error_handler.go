package common 

import "github.com/sunshower-io/updraft/common/ir"


type RuntimeErrorHandler interface {
    
    Flag(
            ir.IntermediateNode, 
            RuntimeErrorCode,
            Operation,
    )
}


