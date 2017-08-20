package ir

import (
    "errors"
    "strings"
    "strconv"
    "fmt"
)

type InsufficientChildrenAtDepth struct {
    
    depth       int
    index       int
    currentNode IntermediateNode
} 

func (f InsufficientChildrenAtDepth) Error() string {
    
    return fmt.Sprintf("Error: Node %v at " +
            "depth %d has insufficient " +
            "children (%d) to satisfy index %d", 
            f.currentNode,
            f.depth, 
            f.currentNode.Arity(), 
            f.index,
    )
    
}

var InvalidHeightError = errors.New("Height of the current tree is insufficient to satisfy path")

func Index() Matcher {
    return indexMatcher{}
}


type indexMatcher struct {
    
    
}


func (m indexMatcher) PathBuilder() PathBuilder {
    return indexBuilder{}
}

type indexBuilder struct {
    
}

func (i indexBuilder) To(rep interface{}) Path {
    return indexPath{path:rep.(string)}
}



type indexPath struct {
    path string
    
}

func (i indexPath)  Traverse(
        node IntermediateNode,
) (IntermediateNode, error) {
    
    
    paths := strings.Split(i.path, "/")
    current := node
    
    for depth, turn := range paths {
        
        if turn == "" {
            continue 
        }
        
        if current == nil {
            return nil, InvalidHeightError
        }
        
        
        
        idx, er := strconv.Atoi(turn)
        
        if er != nil {
            return nil, er
        }
        
        children := current.GetChildren()
        arity := len(children)
        
        if children != nil && idx < arity {
            current = children[idx]
        } 
        
        if idx >= arity {
            return nil, InsufficientChildrenAtDepth{
                depth: depth,
                index: idx,
                currentNode: current,
            } 
        }
        
    }
    
    return current, nil
}




