package front

import (
    "gitlab.com/sunshower.io/updraft/middle/core"
    "gitlab.com/sunshower.io/updraft/common/compiler"
    "gitlab.com/sunshower.io/updraft/common/observer"
    "fmt"
)

type PascalErrorHandler struct {
    compiler.ErrorHandler
    
    MaxErrors       int
    Compiler        compiler.Compiler
    
    
    
    errorCount      int
}



func (p *PascalErrorHandler) Flag(
        stage compiler.Stage,
        token core.Token,
        value interface{},
) {
    
    disp := p.Compiler.GetDispatcher(stage)
    disp.SendMessage(
        &SyntaxError {
            BaseEvent: &observer.BaseEvent {
                Body:  value,
                Topic: observer.SYNTAX_ERROR,
            },
            Token :token,
        },
    )
}

type SyntaxError struct {
    *observer.BaseEvent
    
    Token   core.Token
}


func (s *SyntaxError) Format() string {
    return fmt.Sprintf("ERROR{type:'syntax', token: %s", s.Token)
}
