package core


type syncSet map[TokenType]bool

func (s syncSet) Contains(t Token) bool {
    _, ok := s[t.GetType()]
    return ok
}


type TokenSet interface {
    Contains(Token) bool 
}

func NewSynchronizationSet(tokens ...TokenType) TokenSet {
   
    syncset := make(syncSet)
    for _, token := range tokens {
        syncset[token] = true 
    }
    return syncset
}