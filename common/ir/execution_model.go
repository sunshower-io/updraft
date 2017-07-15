package ir


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