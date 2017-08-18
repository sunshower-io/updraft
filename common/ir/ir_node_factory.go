package ir

import "github.com/sunshower-io/updraft/middle/core"

type IntermediateNodeFactory interface {
   
    /**
     */
    NewExecutionModel() ExecutionModel
    
    
    /**
    Create a new IR node of the given type
     */
    
    NewIntermediateNode(
            IntermediateNodeType, 
            token core.Token,
    ) IntermediateNode
    
    
}
