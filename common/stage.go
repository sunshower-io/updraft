package common


type Stage string


const (
    LEXING      Stage = "stage:lexing"
    PARSING     Stage = "stage:parsing"
    OPTIMIZING  Stage = "stage:optimizing"
    EXECUTING   Stage = "stage:executing"

)
