package elements

import (
    "github.com/sunshower-io/updraft/common/frontend"
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/pascal/tokens"
    "github.com/sunshower-io/updraft/common/syntax"
)

func NewStatementParser(
        parser frontend.Parser,
) *StatementParser {
    return &StatementParser{}
}


func setLineNumber(
        token core.Token,
        node ir.IntermediateNode,
) {
    if token != nil {
        node.SetLine(token.GetLineNumber())
    }
}

type StatementParser struct {
    syntax.ElementParser
    
    Parser                      frontend.Parser
    
    ExecutionModelFactory       ir.ExecutionModelFactory
}


func (s *StatementParser) Parse(
        token core.Token,
) ir.IntermediateNode {
   
    var executionModel ir.IntermediateNode 
    
    switch token.GetType() {

    case tokens.BEGIN : 
        
        
        compoundStatementParser := NewCompoundParser(s)
        executionModel = compoundStatementParser.Parse(token)
    case tokens.IDENTIFIER: 
        
        
        assignmentStatmentParser := NewAssignmentParser(s)
        executionModel = assignmentStatmentParser.Parse(token)
    default: 
        
        
        executionModel = s.ExecutionModelFactory.NewNode(ir.NO_OP)
    }
    
    setLineNumber(token, executionModel)
    
    return executionModel
}



