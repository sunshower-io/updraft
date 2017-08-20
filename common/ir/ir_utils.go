package ir




func DepthOf(model ExecutionModel) int {
    return Depth(model.GetRoot())
}

func Depth(node IntermediateNode) int {
    return depthOf(0, node)
}

func depthOf(cd int, node IntermediateNode) int {
    if node == nil {
        return cd 
    }
   
    children := node.GetChildren()
    
    if children == nil || len(children) == 0 {
        return cd + 1
    }
    
   
    max := 0 
    for _, child := range children {
        depth := depthOf(cd + 1, child)
        if depth > max {
            max = depth
        }
    }
    return max
}


//func DecendentAtPosition(
//        node IntermediateNode,
//        path Path,
//)  IntermediateNode {
//    
//    
//    
//}


