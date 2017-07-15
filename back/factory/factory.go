package factory

import (
	"github.com/sunshower-io/updraft/back/codegen"
	"github.com/sunshower-io/updraft/back/interpreter"
	"github.com/sunshower-io/updraft/common/backend"
	"github.com/sunshower-io/updraft/common/observer"
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
