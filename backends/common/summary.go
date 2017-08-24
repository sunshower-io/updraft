package common

import "time"

type Summary struct {
    OperationCount uint
    ErrorCount     uint 
    ElapsedTime    time.Duration
}