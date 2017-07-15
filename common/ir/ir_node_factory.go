package ir


type IntermediateNodeFactory interface {
   
    /**
     */
    NewExecutionModel() ExecutionModel
    
    
    /**
    Create a new IR node of the given type
     */
    
    NewIntermediateNode(IntermediateNodeType) IntermediateNode
    
    
}
