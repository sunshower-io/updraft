package codegen

import (
	"github.com/sunshower-io/updraft/common/backend"
	"github.com/sunshower-io/updraft/common/ir"
)

type CodeGenerator struct {
	backend.Backend

	SymbolTable ir.SymbolTable
}

func (g *CodeGenerator) Process(
	code ir.ExecutionModel,
	symbolTable ir.SymbolTable,
) error {

	return nil
}
