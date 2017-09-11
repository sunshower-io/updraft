package elements

import (
    "github.com/sunshower-io/updraft/common/frontend"
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/pascal/tokens"
    ctokens "github.com/sunshower-io/updraft/common/tokens"
    "github.com/sunshower-io/updraft/common/syntax"
    "github.com/sunshower-io/updraft/common/errors"
    "github.com/sunshower-io/updraft/common"
)



var statementSynchronizationCache = make(map[core.TokenType] core.TokenSet)

var statementSynchronizationSet = core.NewSynchronizationSet(
    tokens.BEGIN,
    tokens.CASE,
    tokens.FOR,
    tokens.IF,
    tokens.REPEAT,
    tokens.WHILE,
    tokens.IDENTIFIER,
    tokens.SEMICOLON,
) 


var statementTerminatorSynchronizationSet = core.NewSynchronizationSet(
    tokens.SEMICOLON,
    tokens.END,
    tokens.ELSE,
    tokens.UNTIL,
    tokens.DOT,
)


func NewStatementParser(
        parser frontend.Parser,
) *StatementParser {
    return &StatementParser{
        Parser: parser,
        ElementParser: &syntax.BaseElementParser{
            Parser: parser,
        },
        ExecutionModelFactory: parser.GetExecutionModelFactory(),
        ErrorHandler: parser.GetErrorHandler(),
        
    }
}


func setLineNumber(
        token core.Token,
        node ir.IntermediateNode,
) {
    if token != nil {
        node.SetLineNumber(token.GetLineNumber())
    }
}

type StatementParser struct {
    syntax.ElementParser
    
    Parser                      frontend.Parser
    
    ErrorHandler                errors.ErrorHandler
    
    ExecutionModelFactory       ir.ExecutionModelFactory
}





func (s *StatementParser) ParseList(
        token core.Token,
        parent ir.IntermediateNode,
        terminator core.TokenType,
        errorCode ctokens.ErrorCode,
) error {
    
    var (
        err error
        terminated bool
        tokenType core.TokenType 
    )

    
    terminators := createOrRetrieveCached(terminator)

    for  {
    
    
        if tokenType, terminated = terminate(token, terminator); terminated {
            break
        }
       
        if token, err = s.CurrentToken(); err != nil {
            return err
        }
        
        child, err := s.Parse(token)
        parent.AddChild(child)
        
        if token, err = s.CurrentToken(); err != nil {
            return err
        }
        
        tokenType = token.GetType()
        switch tokenType {
        case tokens.SEMICOLON:
            token, err = s.NextToken()
        case tokens.IDENTIFIER:
            s.ErrorHandler.FlagError(
                common.PARSING, 
                token, 
                s,
                tokens.MISSING_SEMICOLON,
            )

        default:
            if terminators.Contains(token) {
                errorCode = tokens.MISSING_SEMICOLON 
            } else {
                errorCode = tokens.UNEXPECTED_TOKEN
            }
            
            token, err = s.Parser.Synchronize(terminators)
            goto DONE
        }
        
    }

DONE:
    if tokenType == terminator {
        token, err = s.NextToken()
    } else {
        s.ErrorHandler.FlagError(
            common.PARSING,
            token,
            s,
            errorCode,
        )
    }
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
    case tokens.REPEAT:
        repeatParser := NewRepeatParser(s)
        executionModel, err = repeatParser.Parse(token)

    //case tokens.WHILE: 
    //    whileParser := NewWhileParser(s)
    //    executionModel, er = whileParser.Parse(token)
    //case tokens.FOR:
    //    forParser := NewForParser(s)
    //    executionModel, err = forParser.Parse(token)
    //case tokens.IF: 
    //    ifParser := NewIfParser(s)
    //    executionModel, err = ifParser.Parse(token)
    //case tokens.CASE:
    //    caseParser := NewCaseParser(s) 
    //    executionModel, err = caseParser.Parse(token)
    default: 
        executionModel = s.ExecutionModelFactory.NewNode(ir.NO_OP, token)
    }
    setLineNumber(token, executionModel)
    return executionModel, err 
}



func terminate(
        token core.Token,
        terminatorType core.TokenType,
) (core.TokenType, bool) {
    
    if token.GetType() == terminatorType {
        return terminatorType, true
    }
    
    switch token.(type) {
    
    case *core.EofToken:
        return token.GetType(), true
    }
    return terminatorType, false
}




func createOrRetrieveCached(terminator core.TokenType) core.TokenSet {
    if tokenSet, ok := statementSynchronizationCache[terminator]; ok {
        return tokenSet
    }
    terminatorSet := statementSynchronizationSet.CloneAndAppend(terminator)
    statementSynchronizationCache[terminator] = terminatorSet 
    return terminatorSet
}
