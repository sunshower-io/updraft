package backend

import (
	"github.com/sunshower-io/updraft/common/ir"
	"github.com/sunshower-io/updraft/common/observer"
)

type Backend interface {
	observer.EventProducer

	Process(ir.ExecutionModel, ir.SymbolTable) error
}
