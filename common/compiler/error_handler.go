package compiler

import (
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/pascal/tokens"
)

type ErrorHandler interface {

	Flag(Stage, core.Token, interface{})
    
    FlagError(
            Stage, 
            core.Token, 
            interface{}, 
            tokens.ErrorCode,
    )

	GetErrorCount() int
}
