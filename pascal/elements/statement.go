package elements

import (
    "github.com/sunshower-io/updraft/common/frontend"
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/pascal/tokens"
    "github.com/sunshower-io/updraft/common/syntax"
    "github.com/sunshower-io/updraft/common/compiler"
)

func NewStatementParser(
        parser frontend.Parser,
) *StatementParser {
    return &StatementParser{
        Parser: parser,
        ElementParser: &syntax.BaseElementParser{
            Parser: parser,
        },
    }
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
    
    ErrorHandler                compiler.ErrorHandler
    
    ExecutionModelFactory       ir.ExecutionModelFactory
}





func (s *StatementParser) ParseList(
        token core.Token,
        parent ir.IntermediateNode,
        terminator core.TokenType,
        errorCode tokens.ErrorCode,
) error {
    
    var (
        err error
    )
    for  {
        
        
        if terminate(token, terminator) {
            goto DONE
        }
       
        if token, err = s.CurrentToken(); err != nil {
            return err
        }
        
        child, err := s.Parse(token)
        parent.AddChild(child)
        
        if token, err = s.CurrentToken(); err != nil {
            return err
        }
        
        tokenType := token.GetType()
        switch tokenType {
        case tokens.SEMICOLON:
            token, err = s.NextToken()
        case tokens.IDENTIFIER:
            s.ErrorHandler.FlagError(
                compiler.PARSING, 
                token, 
                s,
                tokens.MISSING_SEMICOLON,
            )
        }
        
        if tokenType == terminator {
            token, err = s.NextToken()
        } else {
            s.ErrorHandler.FlagError(
                compiler.PARSING, 
                token,
                s,
                errorCode, 
            )
        }
    }
    
DONE: 
    return err
    
    
}

func (s *StatementParser) Parse(
        token core.Token,
) (ir.IntermediateNode, error) {
   
    var (
        err error
        executionModel ir.IntermediateNode
    )
    
    
    switch token.GetType() {
    case tokens.BEGIN : 
        compoundStatementParser := NewCompoundParser(s)
        executionModel, err = compoundStatementParser.Parse(token)
    case tokens.IDENTIFIER: 
        assignmentStatmentParser := NewAssignmentParser(s)
        executionModel, err = assignmentStatmentParser.Parse(token)
    default: 
        executionModel = s.ExecutionModelFactory.NewNode(ir.NO_OP)
    }
    setLineNumber(token, executionModel)
    return executionModel, err 
}



func terminate(
        token core.Token,
        terminatorType core.TokenType,
) bool {
    
    if token.GetType() == terminatorType {
        return true
    }
    
    switch token.(type) {
    
    case *core.EofToken:
        return true
    }
    
    return false
}
