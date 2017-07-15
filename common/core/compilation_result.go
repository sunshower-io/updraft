package core

type CompilationResult interface {

}


type AbstractCompilationResult struct {
    CompilationResult

}

func NewCompilationResult() CompilationResult {

    return new(AbstractCompilationResult)
}