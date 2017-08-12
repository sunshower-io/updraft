package compiler

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/sunshower-io/updraft/common/observer"
    "github.com/sunshower-io/updraft/common"
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
    compiler.AddListener(common.EXECUTING, new(MockListener))
}

func TestAddingListenerResultsInListenerAppearingInListenerListForStage(t *testing.T) {
    compiler := new(AbstractCompiler)
    compiler.AddListener(common.EXECUTING, new(MockListener))
    assert.Equal(t, len(compiler.GetListeners(common.EXECUTING)), 1)
}


func TestAddingListenerResultsInListenerNotAppearingInIncorrectRegion(t *testing.T) {
    compiler := new(AbstractCompiler)
    compiler.AddListener(common.EXECUTING, new(MockListener))
    assert.Equal(t, len(compiler.GetListeners(common.EXECUTING)), 1)
    assert.Equal(t, len(compiler.GetListeners(common.OPTIMIZING)), 0)
}


func TestDispatchingMessageResultsInCorrectMessageDelivered(t *testing.T) {
    
    compiler := new(AbstractCompiler)
    listener := new(MockListener)
    listener.topic = "frapper"
    message := newMessage("frapper")
    compiler.AddListener(common.EXECUTING, listener)
    compiler.Dispatch(common.EXECUTING, message)
    assert.Equal(t, listener.message, message)
}



func newMessage(t observer.EventType) observer.Message {
    return &observer.BaseEvent{
        Topic: t,
    }
}
