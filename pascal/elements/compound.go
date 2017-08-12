package elements

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/middle/core"
    pir "github.com/sunshower-io/updraft/pascal/ir"
    "github.com/sunshower-io/updraft/pascal/tokens"
)

func NewCompoundParser(s *StatementParser) *CompoundParser {
    return &CompoundParser{s}
}

type CompoundParser struct {
    *StatementParser
}

func (c *CompoundParser) Parse(
        t core.Token,
) (ir.IntermediateNode, error) {
    
    token, er := c.NextToken()
    compoundNode := c.ExecutionModelFactory.NewNode(pir.COMPOUND)
    
    statementParser := NewStatementParser(c.StatementParser.Parser)
    statementParser.ParseList(
        token, 
        compoundNode, 
        tokens.END, 
        tokens.MISSING_END,
    )
    
    return compoundNode, er
}



