package errors


import (
    "errors"
    "github.com/sunshower-io/updraft/common"
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/common/tokens"
)

var NotImplemented = errors.New("Method is not implemented")


type ErrorHandler interface {
    
    Flag(common.Stage, core.Token, interface{})
    
    FlagError(
            common.Stage,
            core.Token,
            interface{},
            tokens.ErrorCode,
    )
    
    GetErrorCount() int
}
