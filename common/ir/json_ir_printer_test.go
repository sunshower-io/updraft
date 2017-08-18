package ir

import (
    "testing"
    "bytes"
    "github.com/sunshower-io/updraft/middle/core"
)

var factory = NewExecutionModelFactory()

func TestJsonPrinterWorksForSingleNode(t *testing.T) {
    
    node := factory.NewNode(
        INTEGER,
        core.DetachedToken("frap", "adap", 1, 1),
    )
    node.AddChild(
        factory.NewNode(
            INTEGER,
            core.DetachedToken("frap", "adap", 1, 1),
        ),
    )
    node.AddChild(
        factory.NewNode(
            INTEGER,
            core.DetachedToken("frap", "adap", 1, 1),
        ),
    )
    model := factory.NewExecutionModel()
    model.SetRoot(node)
    
    writer := &bytes.Buffer{}
    
    new(JsonExecutionModelPrinter).PrintTo(model, writer)
    
    println(writer.String())
    
}
