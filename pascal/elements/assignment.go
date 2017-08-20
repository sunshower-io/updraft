package elements

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/middle/core"
    "strings"
    "github.com/sunshower-io/updraft/pascal/tokens"
    "github.com/sunshower-io/updraft/common"
)


func NewAssignmentParser (
        parent *StatementParser,
) *AssignmentStatementParser {
    return &AssignmentStatementParser{parent}
}

type AssignmentStatementParser struct {
    *StatementParser
}

func (p *AssignmentStatementParser) Parse(
        token core.Token,
) (ir.IntermediateNode, error) {
    
    symbolTables := p.Parser.GetSymbolTables()
    
    assignment := p.ExecutionModelFactory.NewNode(ir.ASSIGN, token)
    
    target := strings.ToLower(token.GetText())
    symbol, er := symbolTables.Resolve(target)
    if er != nil {
        symbol, er = symbolTables.EnterLocal(target) 
    }
    
    symbol.AddLine(&ir.Line{
        Number: token.GetLineNumber(),
    })
    
    if token, er = p.NextToken(); er != nil {
        return nil, er
    }
    
    variableNode := p.ExecutionModelFactory.NewNode(ir.VARIABLE, token)
   
    variableNode.SetAttribute(ir.ID, symbol)
    variableNode.SetValue(target)
    
    assignment.AddChild(variableNode)
    
    if token.GetType() == tokens.COLON_EQUALS {
        token, er = p.NextToken()
    } else {
        p.ErrorHandler.FlagError(
            common.PARSING, 
            token, 
            p, 
            tokens.MISSING_COLON_EQUALS,
        )
    }
    
    expressionParser := NewExpressionParser(p.StatementParser)
    
    expressionNode, er := expressionParser.Parse(token)
    assignment.AddChild(expressionNode)
    
    
    return assignment, er
}


