package factory

import (
	"gitlab.com/sunshower.io/updraft/back/codegen"
	"gitlab.com/sunshower.io/updraft/back/interpreter"
	"gitlab.com/sunshower.io/updraft/common/backend"
	"gitlab.com/sunshower.io/updraft/common/observer"
)

type BackendType string

const (
	COMPILER    = "backend::compiler"
	INTERPRETER = "backend::interpreter"
)

func NewBackend(t BackendType) backend.Backend {
	switch t {
	case COMPILER:
		return newCompiler()
	}
	return newInterpreter()
}

func newCompiler() backend.Backend {
	return &codegen.CodeGenerator{}
}

func newInterpreter() backend.Backend {
	return &interpreter.Interpreter{
		EventProducer: observer.NewEventProducer(),
	}
}
