package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
)

type StatementReducer struct {
    common.Operation
    Context             ReducerContext 
    ErrorHandler        common.RuntimeErrorHandler
    SymbolTables        ir.SymbolTableStack
}

func (s *StatementReducer) Apply(
        node ir.IntermediateNode,
        ctx backend.OperationContext,
) interface{} {
    if operation := s.Context.Resolve(s, node); operation != nil {
        return operation.Apply(node, ctx)
    } else {
        s.ErrorHandler.Flag(
            node, 
            backend.UNSUPPORTED_FEATURE, 
            s,
        )
        return nil
    }
}
