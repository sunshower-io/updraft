package syntax

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/middle/core"
)

type ElementParser interface {
    Parse(token core.Token) ir.IntermediateNode
}
