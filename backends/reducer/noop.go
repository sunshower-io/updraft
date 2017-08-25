package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)

type NoOp struct {

}

func (o NoOp) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    return nil
}
