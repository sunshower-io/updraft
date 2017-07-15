package compiler

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "gitlab.com/sunshower.io/updraft/common/observer"
)


type MockListener struct {
    observer.EventListener
    
    topic           observer.EventType
    
    message         observer.Message
}

func (t *MockListener) ListensFor(m observer.Message) bool {
    return t.topic == m.TopicId()
}

func (t *MockListener) OnMessage(m observer.Message) {
    t.message = m
}



func TestAddingListenerToCompilerDoesNotThrowExceptions(t *testing.T) {
    compiler := new(AbstractCompiler)
    compiler.AddListener(EXECUTING, new(MockListener))
}

func TestAddingListenerResultsInListenerAppearingInListenerListForStage(t *testing.T) {
    compiler := new(AbstractCompiler)
    compiler.AddListener(EXECUTING, new(MockListener))
    assert.Equal(t, len(compiler.GetListeners(EXECUTING)), 1)
}


func TestAddingListenerResultsInListenerNotAppearingInIncorrectRegion(t *testing.T) {
    compiler := new(AbstractCompiler)
    compiler.AddListener(EXECUTING, new(MockListener))
    assert.Equal(t, len(compiler.GetListeners(EXECUTING)), 1)
    assert.Equal(t, len(compiler.GetListeners(OPTIMIZING)), 0)
}


func TestDispatchingMessageResultsInCorrectMessageDelivered(t *testing.T) {
    
    compiler := new(AbstractCompiler)
    listener := new(MockListener)
    listener.topic = "frapper"
    message := newMessage("frapper")
    compiler.AddListener(EXECUTING, listener)
    compiler.Dispatch(EXECUTING, message)
    assert.Equal(t, listener.message, message)
}



func newMessage(t observer.EventType) observer.Message {
    return &observer.BaseEvent{
        Topic: t,
    }
}
