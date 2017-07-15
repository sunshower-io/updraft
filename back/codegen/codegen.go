package codegen

import (
	"gitlab.com/sunshower.io/updraft/common/backend"
	"gitlab.com/sunshower.io/updraft/common/ir"
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
