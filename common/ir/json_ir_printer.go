package ir

import (
    "io"
    "strings"
    "fmt"
)

const (
    lineWidth     = 80
    indentWidth   = 4
)


type JsonExecutionModelPrinter struct {
    ExecutionModelPrinter
    
}


func (p *JsonExecutionModelPrinter) PrintTo(
        model ExecutionModel,
        out io.Writer,
) {

    p.writeNode(
        model.GetRoot(),
        out,
        0,
    )
}


func (p *JsonExecutionModelPrinter) writeNode(
        n IntermediateNode,
        w io.Writer,
        depth int,
) {
   
    p.write(w, depth + 1, "{\"node\"")
    
    
    p.write(
        w,
        depth + 1,
        "\"type\":\"%s\"",
        n.GetType(),
        depth + 1,
    )
    
}


func (p *JsonExecutionModelPrinter) write(
        writer io.Writer,
        depth int,
        fmtString string,
        value ...interface{},
) {
    writer.Write([]byte(strings.Repeat(" ", depth)))
    writer.Write([]byte(fmt.Sprintf(fmtString, value)))
}


func (p *JsonExecutionModelPrinter) indent(
        writer io.Writer,
        depth int,
) {
    writer.Write([]byte(strings.Repeat(" ", depth)))
}