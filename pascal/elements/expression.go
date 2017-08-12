package elements

import (
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/pascal/tokens"
    pir "github.com/sunshower-io/updraft/pascal/ir"
    "strings"
    "github.com/sunshower-io/updraft/common/compiler"
)


type OperationMap map[core.TokenType]ir.IntermediateNodeType

var relationTypes = relationalTypes()

var additiveTypes = getAdditiveTypes()

var multiplicativeTypes = getMultiplicativeTypes()


func NewExpressionParser(p *StatementParser) *ExpressionParser {
    return &ExpressionParser{p}
}


type ExpressionParser struct {
    *StatementParser
}


func (p *ExpressionParser) parseFactor(
        token core.Token,
) ir.IntermediateNode {
    
    tokenType := token.GetType()
    
    
    var rootNode ir.IntermediateNode
   
    symbolTables := p.Parser.GetSymbolTables()
    
    switch tokenType {
    
    case tokens.IDENTIFIER:
        name := strings.ToLower(token.GetText())
        entry, _ := symbolTables.Resolve(name)
        
        if entry == nil {
            p.ErrorHandler.FlagError(
                compiler.PARSING, 
                token, 
                p, 
                tokens.IDENTIFIER_UNDEFINED,
            )
            entry, _ = symbolTables.EnterLocal(name)
        }
        
        rootNode = p.ExecutionModelFactory.NewNode(ir.VARIABLE)
        rootNode.SetAttribute(ir.ID, name)
        entry.AddLine(&ir.Line{
            Number:token.GetLineNumber(),
        })
        token, _ = p.NextToken()
        
    case tokens.INTEGER:
        rootNode = p.ExecutionModelFactory.NewNode(ir.INTEGER)
        rootNode.SetValue(token.GetValue())
        token, _ = p.NextToken()
    case tokens.REAL:
        rootNode = p.ExecutionModelFactory.NewNode(ir.FLOAT)
        rootNode.SetValue(token.GetValue())
        token, _ = p.NextToken()
        
    case tokens.STRING:
        rootNode = p.ExecutionModelFactory.NewNode(ir.STRING_LITERAL)
        rootNode.SetValue(token.GetValue())
        token, _ = p.NextToken()
        
    case tokens.NOT:
        token, _ = p.NextToken()
        rootNode = p.ExecutionModelFactory.NewNode(ir.NOT)
        
        rootNode.AddChild(p.parseFactor(token))
        
    case tokens.LBRACKET:
        
        token, _ = p.NextToken()
        
        rootNode, _ = p.parseExpression(token)
        
        token, _ = p.CurrentToken()
        
        if token.GetType() == tokens.RPAREN {
            token, _ = p.NextToken()
        } else {
            p.ErrorHandler.FlagError(
                compiler.PARSING, 
                token,
                p, 
                tokens.UNEXPECTED_TOKEN,
            )
        }

    default:
        p.ErrorHandler.FlagError(
            compiler.PARSING,
            token,
            p,
            tokens.UNEXPECTED_TOKEN,
        )
        
    }
    return rootNode
}


func (p *ExpressionParser) parseTerm(
        token core.Token,
) ir.IntermediateNode {
    
    
    rootNode := p.parseFactor(token)
    
    token, _ = p.CurrentToken()
    tokenType := token.GetType()
    
    for nodeType, ok := multiplicativeTypes[tokenType]; ok; {
        operandNode := p.ExecutionModelFactory.NewNode(nodeType)
        operandNode.AddChild(rootNode)
        token, _ = p.NextToken()
        
        child := p.parseFactor(token)
        operandNode.AddChild(child)
        rootNode = operandNode
        token, _ = p.CurrentToken()
        tokenType = token.GetType()
    }
    
    return rootNode
}


func (p *ExpressionParser) parseSimpleExpression(
        token core.Token,
) (ir.IntermediateNode, error) {
    
    
    var (
        er error 
        signType core.TokenType
    )
    
    tokenType := token.GetType()
    
    if tokenType == tokens.PLUS || tokenType == tokens.MINUS {
        signType =  tokenType
        token, er = p.NextToken()
    }
    
    
    rootNode := p.parseTerm(token)
    
    
    if signType == tokens.MINUS {
        negate := p.ExecutionModelFactory.NewNode(pir.NEGATE)
        negate.AddChild(rootNode)
        rootNode = negate
    }
    
    
    token, _ = p.CurrentToken()
    
    tokenType = token.GetType()
    
   
    
    for nodeType, ok := additiveTypes[tokenType]; ok; {
        operandNode := p.ExecutionModelFactory.NewNode(nodeType)
        operandNode.AddChild(rootNode)
        token, er = p.NextToken()
        
        term := p.parseTerm(token)
        
        operandNode.AddChild(term)
        rootNode = operandNode
        token, _ = p.CurrentToken()
        tokenType = token.GetType()
    }
    
    
    return rootNode, er
}

func (p *ExpressionParser) parseExpression(
        token core.Token,
) (ir.IntermediateNode, error) {
    
    
    root, er  := p.parseSimpleExpression(token)
    
    token, er = p.CurrentToken()
    
    tokenType := token.GetType()
    
    
    if irNodeType, exists := relationTypes[tokenType]; exists {
        operandNode := p.ExecutionModelFactory.NewNode(irNodeType)
        operandNode.AddChild(root)
        token, er = p.NextToken()
        child, _ := p.parseSimpleExpression(token)
        operandNode.AddChild(child)
        root = operandNode
    }
    
    return root, er
}



func (p *ExpressionParser) Parse(
        token core.Token,
) (ir.IntermediateNode, error) {
    
    return nil, nil
}


func getMultiplicativeTypes() OperationMap {
    result := make(OperationMap)
    result[tokens.STAR] = pir.MULTIPLY
    result[tokens.SLASH] = pir.FLOAT_DIVIDE
    return result
}

func getAdditiveTypes() OperationMap {
    result := make(OperationMap)
    result[tokens.PLUS] = pir.ADD
    result[tokens.MINUS] = pir.SUBTRACT
    result[tokens.OR] = pir.OR
    return result
}



func relationalTypes() OperationMap {
    result := make(OperationMap)
    result[tokens.EQUALS] = pir.EQUAL_TO
    result[tokens.NOT_EQUALS] = pir.NOT_EQUAL_TO
    result[tokens.LT] = pir.LESS_THAN
    result[tokens.GT] = pir.GREATER_THAN
    result[tokens.LTE] = pir.LTE
    result[tokens.GTE] = pir.GTE
    return result
}
