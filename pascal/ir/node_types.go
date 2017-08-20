package ir

import "github.com/sunshower-io/updraft/common/ir"

func init() {
    ir.RegisterIntermediateType(COMPOUND, "compound")
    ir.RegisterIntermediateType(LOOP, "loop")
    ir.RegisterIntermediateType(TEST, "test")
    ir.RegisterIntermediateType(ADD, "+")
    ir.RegisterIntermediateType(SUBTRACT, "-")
    ir.RegisterIntermediateType(MULTIPLY, "*")
    ir.RegisterIntermediateType(FLOAT_DIVIDE, "/")
    ir.RegisterIntermediateType(INTEGER_DIVIDE, "/")
}


const (
    
    PROGRAM             ir.IntermediateNodeType = 0
    PROCEDURE           ir.IntermediateNodeType = 1
    FUNCTION            ir.IntermediateNodeType = 2
    
    COMPOUND            ir.IntermediateNodeType = 3
    
    ASSIGNMENT          ir.IntermediateNodeType = 4
    LOOP                ir.IntermediateNodeType = 5
    TEST                ir.IntermediateNodeType = 6
    CALL                ir.IntermediateNodeType = 7
    PARAMETERS          ir.IntermediateNodeType = 8
    IF                  ir.IntermediateNodeType = 9
    SELECT              ir.IntermediateNodeType = 10
    SELECT_BRANCH       ir.IntermediateNodeType = 11
    SELECT_CONSTANTS    ir.IntermediateNodeType = 12
    NO_OP               ir.IntermediateNodeType = 13
    LESS_THAN           ir.IntermediateNodeType = 14
    GREATER_THAN        ir.IntermediateNodeType = 15
    EQUAL_TO            ir.IntermediateNodeType = 16
    NOT_EQUAL_TO        ir.IntermediateNodeType = 17
    NOT                 ir.IntermediateNodeType = 18
    GTE                 ir.IntermediateNodeType = 19
    LTE                 ir.IntermediateNodeType = 20
    ADD                 ir.IntermediateNodeType = 21
    SUBTRACT            ir.IntermediateNodeType = 22
    OR                  ir.IntermediateNodeType = 23
    NEGATE              ir.IntermediateNodeType = 24
    MULTIPLY            ir.IntermediateNodeType = 25
    INTEGER_DIVIDE      ir.IntermediateNodeType = 26
    FLOAT_DIVIDE        ir.IntermediateNodeType = 27
    MOD                 ir.IntermediateNodeType = 29
    AND                 ir.IntermediateNodeType = 30
    VARIABLE            ir.IntermediateNodeType = 31
    SUBSCRIPTS          ir.IntermediateNodeType = 32
    FIELD               ir.IntermediateNodeType = 33
    INTEGER_CONSTANT    ir.IntermediateNodeType = 34
    REAL_CONSTANT       ir.IntermediateNodeType = 35
    STRING_CONSTANT     ir.IntermediateNodeType = 36
    BOOLEAN_CONSTANT    ir.IntermediateNodeType = 37
)


