package elements

import (
    "github.com/sunshower-io/updraft/common/syntax"
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/common/ir"
)

func NewCompoundParser(s *StatementParser) *CompoundParser {
    return &CompoundParser{}
}

type CompoundParser struct {
    syntax.ElementParser
}

func (c *CompoundParser) Parse(
        t core.Token,
) ir.IntermediateNode {
    
    return nil
    
}
