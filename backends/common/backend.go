package common

import (
    "github.com/pkg/errors"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/common/observer"
)

type Backend interface {
    observer.EventProducer
    
    Process(
            ir.ExecutionModel,
            ir.SymbolTableStack,
    ) error
}


type BaseBackend struct {
    observer.EventProducer
}

func (b *BaseBackend) Process(
        ir.ExecutionModel,
        ir.SymbolTableStack,
) error {
    return errors.New("I don't do anything")
    
}

func NewBaseBackend() Backend {
    return &BaseBackend{
        EventProducer: observer.NewEventProducer(),
    }
}
