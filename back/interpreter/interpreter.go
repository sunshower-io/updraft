package interpreter

import (
	"github.com/sunshower-io/updraft/common/backend"
	"github.com/sunshower-io/updraft/common/ir"
	"github.com/sunshower-io/updraft/common/observer"
	"time"
)

type InterpreterSummary struct {
	RuntimeErrorCount int
	ExecutionCount    int
	ElapsedTime       time.Duration
}

type Interpreter struct {
	backend.Backend
	observer.EventProducer
}

func (g *Interpreter) RemoveMessageListener(m observer.EventListener) {
	g.EventProducer.RemoveMessageListener(m)
}

func (g *Interpreter) AddEventListener(m observer.EventListener) {
	g.EventProducer.AddEventListener(m)
}

func (g *Interpreter) SendMessage(m observer.Message) {
	g.EventProducer.SendMessage(m)
}

func (g *Interpreter) Process(
	code ir.ExecutionModel,
	symbolTable ir.SymbolTable,
) error {

	startTime := time.Now()

	endTime := time.Since(startTime)

	g.SendMessage(observer.CreateEvent(
		observer.INTERPRETER_SUMMARY,
		&InterpreterSummary{
			RuntimeErrorCount: 0,
			ExecutionCount:    0,
			ElapsedTime:       endTime,
		}))

	return nil
}
