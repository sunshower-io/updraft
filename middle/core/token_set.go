package core




type syncSet map[TokenType]bool

func (s syncSet) Contains(t Token) bool {
    _, ok := s[t.GetType()]
    return ok
}

func (s syncSet) CloneAndAppend(t ...TokenType) TokenSet {
    result := NewSynchronizationSet(t...).(syncSet)
    for k, v := range s {
        result[k] = v
    }
    return result
}

func (s syncSet) Begin() TokenType {
    return TokenType{} 
}


func (s syncSet) CloneAndAppendAll(
        ts TokenSet, 
        tokens...TokenType,
) TokenSet {
    
    return nil
    
    //for tt := ts.Begin(); ts.HasNext() {
    //    
    //}
    
}



type TokenSet interface {
    
    Begin() TokenType
    
    Contains(Token) bool 
    
    CloneAndAppend(...TokenType) TokenSet
    
    CloneAndAppendAll(TokenSet, ...TokenType) TokenSet
}

func NewSynchronizationSet(tokens ...TokenType) TokenSet {
   
    syncset := make(syncSet)
    for _, token := range tokens {
        syncset[token] = true 
    }
    return syncset
}