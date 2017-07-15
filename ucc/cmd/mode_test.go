package cmd

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestRetrievingNonExistantParserModeFails(t *testing.T) {
    
    cfg := &ModeConfiguration{}
   
    var m ParserMode = "frapper"
    
    e, u := cfg.ResolveParserMode(m)
   
    assert.Nil(t, e, "Parser mode should not have been found")
    assert.NotNil(t, u, "Error should've existed")
}

func TestRetrievingNonExistantBackendModeFails(t *testing.T) {
    
    cfg := &ModeConfiguration{}
    
    var m BackendMode = "frapper"
    
    e, u := cfg.ResolveBackendMode(m)
    
    assert.Nil(t, e, "Parser mode should not have been found")
    assert.NotNil(t, u, "Error should've existed")
}
