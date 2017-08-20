package ir

import "testing"

func TestPathApiWorksForIndices(t *testing.T) {
    PathBy(Index()).To("/1/2/3/4/5").Traverse(new(BaseIRNode))
}
