package backend

import (
	"gitlab.com/sunshower.io/updraft/common/ir"
	"gitlab.com/sunshower.io/updraft/common/observer"
)

type Backend interface {
	observer.EventProducer

	Process(ir.ExecutionModel, ir.SymbolTable) error
}
