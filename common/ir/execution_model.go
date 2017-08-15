package ir


func NewExecutionModelFactory() ExecutionModelFactory {
    return &baseExecutionModelFactory{}
}

/**
    Provider of different intermediate node types
 */

type ExecutionModelFactory interface {
    NewExecutionModel() ExecutionModel
    
    NewNode(IntermediateNodeType) IntermediateNode
    
}



type baseExecutionModelFactory struct {
    ExecutionModelFactory
}

func (e *baseExecutionModelFactory) NewExecutionModel() ExecutionModel {
    return &BaseExecutionModel{}
}

func (e *baseExecutionModelFactory) NewNode(
        t IntermediateNodeType,
) IntermediateNode {
    return &BaseIRNode{Type: t}
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
