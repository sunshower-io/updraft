package ir

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultFactoryCreatesSymbolTableStack(t *testing.T) {
	stack := DefaultSymbolTableFactory.CreateStack()
	assert.NotNil(t, stack)
}

func TestDefaultStackReturnsErrorForNoSymbol(t *testing.T) {
	stack := DefaultSymbolTableFactory.CreateStack()
	_, er := stack.Resolve("cool")
	assert.Error(t, er)

}

func TestEnsureResolvingNonExistantValueFailsWithoutPanic(t *testing.T) {

	stack := DefaultSymbolTableFactory.CreateStack()
	_, er := stack.Resolve("cool")
	assert.Error(t, er)
}

func TestDefaultStackReturnsLocalUponEntering(t *testing.T) {
	stack := DefaultSymbolTableFactory.CreateStack()
	_, er := stack.EnterLocal("cool")
	assert.NoError(t, er)
}
