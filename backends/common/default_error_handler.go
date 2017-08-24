package common 

import "github.com/sunshower-io/updraft/common/ir"

type BaseRuntimeErrorHandler struct {
    
}


func (b *BaseRuntimeErrorHandler) Flag(
        node ir.IntermediateNode, 
        code RuntimeErrorCode, 
        operation Operation,
) {
    
    //lineNumber := resolveLine(node)
    
}


func resolveLine(node ir.IntermediateNode) (line string) {
    for line = node.GetLine(); line == ""; line = node.GetLine() {
        node = node.GetParent()
    }
    return 
}



func NewRuntimeErrorHandler() RuntimeErrorHandler{
    return &BaseRuntimeErrorHandler{}
}