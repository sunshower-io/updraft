package elements

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/middle/core"
)


func NewAssignmentParser (
        parent *StatementParser,
) *AssignmentStatementParser {
    return &AssignmentStatementParser{parent}
}

type AssignmentStatementParser struct {
    *StatementParser
}

func (p *AssignmentStatementParser) Parse(token core.Token) (ir.IntermediateNode, error) {
    return p.ExecutionModelFactory.NewNode(ir.NO_OP), nil
}


