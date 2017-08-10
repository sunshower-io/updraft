package ir



/**
    Provider of different intermediate node types
 */

type ExecutionModelFactory interface {
    NewNode(IntermediateNodeType) IntermediateNode
}





/**
    BaseExecutionModel: container for IR trees
 */

type BaseExecutionModel struct {
    ExecutionModel
    
    root        IntermediateNode
}


func (e *BaseExecutionModel) GetRoot() IntermediateNode {
    return e.root
}

func (e *BaseExecutionModel) SetRoot(root IntermediateNode) {
    e.root = root
}
