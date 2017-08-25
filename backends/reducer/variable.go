package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)

type VariableOperation struct {
    common.Operation 
}

func (o VariableOperation) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    sts := ctx.GetSymbolTables()
    st := sts.Peek()
    v, _ := st.Lookup(node.GetValue().(string))
    return v.GetAttribute(ir.DATA_VALUE)
}
