package elements

import (
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/pascal/tokens"
)


func NewRepeatParser(parent *StatementParser) *RepeatParser {
    return &RepeatParser{
        parent,
    }
    
}

type RepeatParser struct {
    *StatementParser
}

func (r *RepeatParser) Parse(
        token core.Token,
) (ir.IntermediateNode, error) {
    execFactory := r.ExecutionModelFactory
    
    token, er := r.NextToken()
    
    iterator := execFactory.NewNode(ir.ITERATE, token)
    condition := execFactory.NewNode(ir.TEST, token)
    
    r.StatementParser.ParseList(
        token, 
        iterator, 
        tokens.UNTIL, 
        tokens.MISSING_UNTIL,
    )
    
    token, er = r.CurrentToken()
    
    expressionParser := NewExpressionParser(r.StatementParser)
    if expressionNode, er := expressionParser.Parse(token); er == nil {
        condition.AddChild(expressionNode)
        iterator.AddChild(condition)
    
    }
    
    
    return iterator, er
}