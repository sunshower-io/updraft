package front

import (
    "fmt"
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/common/compiler"
    "github.com/sunshower-io/updraft/common/observer"
    "github.com/sunshower-io/updraft/common/errors"
    "github.com/sunshower-io/updraft/common"
    "github.com/sunshower-io/updraft/common/tokens"
)

type PascalErrorHandler struct {
    errors.ErrorHandler
    
    MaxErrors       int
    Compiler        compiler.Compiler
    
    
    
    errorCount      int
}



func (p *PascalErrorHandler) Flag(
        stage common.Stage,
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

func (p *PascalErrorHandler) FlagError(
        stage common.Stage, 
        token core.Token, 
        value interface{},
        code tokens.ErrorCode, 
) {
    disp := p.Compiler.GetDispatcher(stage)
    disp.SendMessage(
        &SyntaxError {
            Code: code,
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
    Code    tokens.ErrorCode
}


func (s *SyntaxError) Format() string {
    return fmt.Sprintf("ERROR{type:'syntax', token: %s", s.Token)
}
