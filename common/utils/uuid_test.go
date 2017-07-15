package utils

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestUUIDCanBeUsedToStoreValueInMap(t *testing.T) {
    
    value := make(map[Identifier]string)
    
    id := RandomId()
    
    value[id] = "frap"
    
    assert.Equal(t, value[id], "frap")
    
}