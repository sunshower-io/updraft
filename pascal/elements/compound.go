package elements

import (
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/common/ir"
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
    return c.ExecutionModelFactory.NewNode(ir.NO_OP), nil
}
