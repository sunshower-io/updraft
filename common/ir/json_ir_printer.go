package ir

import (
    "io"
    "strings"
    "fmt"
    "bytes"
)

const (
    lineWidth   = 80
    indentWidth = 4
)

type JsonExecutionModelPrinter struct {
    ExecutionModelPrinter
}

func (p *JsonExecutionModelPrinter) Print(model ExecutionModel) string {
    var writer bytes.Buffer
    p.PrintTo(model, &writer)
    return writer.String()
}

func (p *JsonExecutionModelPrinter) PrintTo(
        model ExecutionModel,
        out io.Writer,
) {
    
    p.writeNode(
        model.GetRoot(),
        out,
        0,
        false,
    )
}

func (p *JsonExecutionModelPrinter) writeNode(
        n IntermediateNode,
        w io.Writer,
        depth int,
        comma bool,

) {
    
    p.write(w, depth, "{\n")
    printNode(p, w, depth, n, comma)
    if comma {
        p.write(w, depth, "},\n")
    } else {
        p.write(w, depth, "}\n")
    }
}

func printNode(
        p *JsonExecutionModelPrinter,
        w io.Writer,
        depth int,
        n IntermediateNode,
        comma bool,
) {
    p.write(w, depth+1, "\"node\":{\n")
    p.write(
        w,
        depth+4,
        "\"type\":\"%v\", \n",
        n.GetType(),
    )
    children := n.GetChildren()
    p.write(w, depth+4, "\"depth\":%d,\n", depth+1)
    token := n.GetToken()
    if token != nil {
        p.write(w, depth+4, "\"value\":\"%v\",\n", n.GetValue())
        p.write(w, depth+4, "\"text\":\"%s\",\n", token.GetText())
        p.write(w, depth+4, "\"value\":\"%s\",\n", token.GetValue())
        p.write(w, depth+4, "\"line\":%d,\n", token.GetLineNumber())
        if children == nil {
            p.write(w, depth+4, "\"position\":%d\n", token.GetPosition())
        } else {
            p.write(w, depth+4, "\"position\":%d,\n", token.GetPosition())
        }
    } else {
        
        if children == nil {
            p.write(w, depth+4, "\"value\":\"%v\"\n", n.GetValue())
        } else {
            p.write(w, depth+4, "\"value\":\"%v\",\n", n.GetValue())
        }
    }
    
    if children != nil {
        p.write(w, depth+4, "\"children\":[\n")
        l := len(children)
        for i, child := range children {
            p.writeNode(child, w, depth+5, i < l - 1)
        }
        p.write(w, depth+4, "]\n")
    }
    
    p.write(w, depth+1, "}\n")
}

func (p *JsonExecutionModelPrinter) write(
        writer io.Writer,
        depth int,
        fmtString string,
        value ...interface{},
) {
    writer.Write([]byte(strings.Repeat(" ", depth)))
    if len(value) > 0 {
        writer.Write([]byte(fmt.Sprintf(fmtString, value...)))
    } else {
        writer.Write([]byte(fmtString))
    }
}

func (p *JsonExecutionModelPrinter) indent(
        writer io.Writer,
        depth int,
) {
    writer.Write([]byte(strings.Repeat(" ", depth)))
}
