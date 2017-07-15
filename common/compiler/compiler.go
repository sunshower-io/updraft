package compiler

import (
    "sync"
    "gitlab.com/sunshower.io/updraft/common"
    "gitlab.com/sunshower.io/updraft/common/io"
    "gitlab.com/sunshower.io/updraft/common/ir"
    "gitlab.com/sunshower.io/updraft/common/observer"
    "gitlab.com/sunshower.io/updraft/common/backend"
    "gitlab.com/sunshower.io/updraft/common/frontend"
    "gitlab.com/sunshower.io/updraft/common/core"
)

type Stage string


const (
    LEXING      Stage = "stage:lexing"
    PARSING     Stage = "stage:parsing"
    OPTIMIZING  Stage = "stage:optimizing"
    EXECUTING   Stage = "stage:executing"

)

type ExecutionResult interface {

}


type Compiler interface {
    
    
    Compile()                   core.CompilationResult
 
    /**
    Current Source-set this compiler is working with
     */
    GetSource()                 io.Source
   
    
    /**
    Get the current back-end for this compiler
     */
    GetBackend()                backend.Backend
   
    /**
    Get the current Parser for this compiler
     */
    GetParser()                 frontend.Parser
    
   
    /**
    
    Get the symbol table for this compiler.  May not be available
    depending on where the compiler is at
     */
    GetSymbolTable()            ir.SymbolTable
    
   
    /**
    Get the execution model for this compiler.  May not be available
    depending on where the compiler is at in the stages
     */
    GetExecutionModel()         ir.ExecutionModel
    
   
    /**
    Add a listener for a given stage in the compilation pipeline
     */
    AddListener(Stage, observer.EventListener)
   
    /**
    Remove a listener for a given stage in the compilation pipeline
     */
    RemoveListener(Stage, producer observer.EventListener)
    
   
    /**
    Run this compiler with the provided configuration
     */
    Run() ExecutionResult
   
    /**
    Dispatch a message on a given stage
     */
    Dispatch(Stage, observer.Message)
   
    /**
    Return this listeners for a given stage
     */
    GetListeners(Stage) []observer.EventListener
   
   
    /**
    
     */
    GetOptions() common.Options
    
    
    GetDispatcher(Stage) observer.EventProducer
   
}


type AbstractCompiler struct {
    
    Compiler
    
    lock                sync.Mutex
    
    Source              io.Source
    
    Options             common.Options
    
    Parser              frontend.Parser
    
    Backend             backend.Backend
    
    SymbolTable         ir.SymbolTable
    
    ExecutionModel      ir.ExecutionModel
    
    listeners           map[Stage] []observer.EventListener
    
    
}


func NewCompiler(
        Source io.Source,
        intermediate ir.ExecutionModel,
)  Compiler {
    compiler := new(AbstractCompiler)
    compiler.Source = Source
    compiler.ExecutionModel = intermediate
    return compiler
}


func (c *AbstractCompiler) GetDispatcher(stage Stage) observer.EventProducer {
    return delegatingDispatcher{
        stage       : stage,
        compiler    : c,
    }
}


func (c *AbstractCompiler) GetSource() io.Source {
    return c.Source
}


func (c *AbstractCompiler) GetParser() frontend.Parser {
    return c.Parser
}


func (c *AbstractCompiler) GetBackend() backend.Backend {
    return c.Backend
}


func (c *AbstractCompiler) GetSymbolTable() ir.SymbolTable {
    return c.SymbolTable
}


func (c *AbstractCompiler) GetExecutionModel() ir.ExecutionModel {
    return c.ExecutionModel
}


func (c *AbstractCompiler) AddListener(
        stage Stage,
        listener observer.EventListener,
) {

    if c.listeners == nil {
        c.listeners = make(map[Stage] []observer.EventListener)
    }
    
   
    listeners, exists := c.listeners[stage]
    if !exists {
        listeners = make([]observer.EventListener, 0)
        c.listeners[stage] = listeners
    }
    
    c.listeners[stage] = append(listeners, listener)
}



func (c *AbstractCompiler) GetListeners(stage Stage) []observer.EventListener {
    if c.listeners != nil {
        return c.listeners[stage]
    }
    return make([]observer.EventListener, 0)
}


func(c *AbstractCompiler) Dispatch(
        stage Stage,
        message observer.Message,
) {
    
    if c.listeners == nil {
        return
    }
    
    listeners, exists := c.listeners[stage]
    
    if !exists {
        return
    }
    
    c.lock.Lock()
    defer c.lock.Unlock()
    
    for _, listener := range listeners {
        if listener.ListensFor(message) {
            listener.OnMessage(message)
        }
    }
    
}



type delegatingDispatcher struct {
    observer.EventProducer
    
    stage       Stage
  
    compiler    *AbstractCompiler
}

func (d delegatingDispatcher) SendMessage(message observer.Message) {
    d.compiler.Dispatch(d.stage, message)
}








