package elements

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/common/syntax"
)


func NewAssignmentParser (
        parent *StatementParser,
) *AssignmentStatementParser {
    return &AssignmentStatementParser{}
}

type AssignmentStatementParser struct {
    syntax.ElementParser
}

func (p *AssignmentStatementParser) Parse(token core.Token) (ir.IntermediateNode, error) {
    return nil, nil
}


