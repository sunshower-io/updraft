package symbols

import "github.com/sunshower-io/updraft/common/ir"

var (
	CONSTANT_VALUE       = ir.KeyedBy("CONSTANT_VALUE")
	ROUTINE_CODE         = ir.KeyedBy("ROUTINE_CODE")
	ROUTINE_SYMBOL_TABLE = ir.KeyedBy("ROUTING_CODE")

	ROUTINE_PARAMETERS = ir.KeyedBy("ROUTINE_PARAMETERS")
	ROUTINE_ROUTINES   = ir.KeyedBy("ROUTINE_ROUTINES")
	DATA_VALUE         = ir.KeyedBy("DATA_VALUE")
)
