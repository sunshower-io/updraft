package common 


type RuntimeErrorCode int

const (
    UNINITIALIZED_VALUE         RuntimeErrorCode = 1
    RANGE_ERROR                 RuntimeErrorCode = 2
    INVALID_INPUT               RuntimeErrorCode = 3
    STACK_OVERFLOW              RuntimeErrorCode = 4
    UNSUPPORTED_FEATURE         RuntimeErrorCode = 5
    DIVISION_BY_ZERO            RuntimeErrorCode = 6
)

