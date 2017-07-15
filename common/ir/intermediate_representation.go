package ir

import (
    "io"
    "gitlab.com/sunshower.io/updraft/common/utils"
)

/**
Enumeration of possible IR attributes
 */
type AttributeKey int


/**
Enumeration of IR node types
 */
type IntermediateNodeType int



/**
    Base type for intermediate representation (IR) trees
 */

type IntermediateNode interface {
   
    /**
    Update the line-number associated with this node
     */
    SetLine(int)
   
    /**
    Retrieve the line-number associated with this node
     */
    GetLine() int
    
   
    /**
    Get the identifier for this node
     */
    GetId()   utils.Identifier
   
    /**
    Set the identifier for this node
     */
    SetId(utils.Identifier)
    
   
    /**
    Set the value of this node.  Many node-types may not have values,
    depending on the structure of the language
     */
    SetValue(interface{})
    
    /**
    Get the value of this node (many node-types may not have values)
     */
    
    GetValue() interface{}
    
    /**
    Return the parent of this IR node, or null if none exists
     */
    GetParent() IntermediateNode
    
    /**
    Set the parent.  Accepts nil to clear a parent
     */
    
    SetParent(
            IntermediateNode,
    ) IntermediateNode
    
   
    /**
    Return the type of this IR node. Return -1 (error condition) if none exists
     */
    
    GetType()   IntermediateNodeType
    
   
    /**
    Add a child to this intermediate node
     */
    AddChild(IntermediateNode)
   
    
    /**
    Return an immutable list of children in this node
     */
    GetChildren() []IntermediateNode
    
   
    /**
    Return the attribute associated with the given key (if it exists) or
    nil
     */
    GetAttribute(AttributeKey) interface{}
   
    /**
    Set an attribute with the given key.  If nil is passed as an attribute value, this method
    will panic
     */
    SetAttribute(
            AttributeKey,
            interface{},
    ) interface{}
   
    /**
    Clear an attribute with the given key
     */
    ClearAttribute(AttributeKey) interface{}
    /**
    Clone this entire node.  This method will only copy attributes.  No children will be copied via this method.
    Use a tree reduction to copy an entire tree
     */
    Clone() IntermediateNode
    
    
    /**
    Return the degree (number of children) of this node
     */
    
    Arity() int
}


type ExecutionModelPrinter interface {
    
    
    Print(ExecutionModel) string
    
    
    PrintTo(
            ExecutionModel,
            io.Writer,
    )
}

/**
    Container for IR tree constructed by parsing and rewriting
 */

type ExecutionModel interface {
  
    /**
    Get the current IR tree
     */
    GetRoot() IntermediateNode
    
    /**
    Set the current tree
     */
    SetRoot(IntermediateNode)
    
    
    Print(ExecutionModelPrinter)
}





type AbstractExecutionModel struct {
    ExecutionModel
}

func NewExecutionModel() ExecutionModel {
    return new(AbstractExecutionModel)
}



