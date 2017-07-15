package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactoryProducesCompiler(t *testing.T) {
	backend := NewBackend(COMPILER)
	assert.NotNil(t, backend)
}
