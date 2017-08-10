package syntax

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/common/frontend"
    "github.com/sunshower-io/updraft/common/errors"
)



type ElementParser interface {
    
    NextToken()                 (core.Token, error)
    
    CurrentToken()              (core.Token, error)
    
    Parse(token core.Token)     (ir.IntermediateNode, error)
    
}



type BaseElementParser struct {
    Parser frontend.Parser
}

// Return the next token, or an error if a 
// lexer failure is encountered
func (p *BaseElementParser) NextToken() (core.Token, error)  {
    return p.Parser.GetNextToken(), nil
}


// Return the current token without advancing the token stream
// Or return an error if the lexer encounters an error
func (p *BaseElementParser) CurrentToken() (core.Token, error) {
    return p.Parser.GetNextToken(), nil
}


// Base method.  Returns errors.NotImplemented by default
func(p *BaseElementParser) Parse(token core.Token) (ir.IntermediateNode, error) {
    return nil, errors.NotImplemented
}



