package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)

type StatementReducer struct {
    common.Operation
    SymbolTables        ir.SymbolTableStack
    ErrorHandler        common.RuntimeErrorHandler
}

func (s *StatementReducer) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    if operation := ctx.Resolve(s, node); operation != nil {
        return operation.Apply(node, ctx)
    } else {
        s.ErrorHandler.Flag(
            node, 
            common.UNSUPPORTED_FEATURE, 
            s,
        )
        return nil
    }
}
