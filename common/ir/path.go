package ir


type Path interface {
    Traverse(IntermediateNode) (IntermediateNode, error)
}

type Matcher interface {
    PathBuilder() PathBuilder
}

type PathBuilder interface {
    To(rep interface{}) Path
}

func PathBy(matcher Matcher) PathBuilder {
    return matcher.PathBuilder() 
}






