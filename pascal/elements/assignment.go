package elements

import "github.com/sunshower-io/updraft/common/syntax"

type AssignmentStatementParser struct {
    syntax.ElementParser
}


func NewAssignmentParser (
        parent *StatementParser,
) *AssignmentStatementParser {
    return &AssignmentStatementParser{}
}