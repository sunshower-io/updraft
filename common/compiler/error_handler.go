package compiler

import "github.com/sunshower-io/updraft/middle/core"

type ErrorHandler interface {

	Flag(Stage, core.Token, interface{})

	GetErrorCount() int
}
