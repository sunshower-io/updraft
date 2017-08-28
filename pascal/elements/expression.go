package elements

import (
    "strings"
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/pascal/tokens"
    "github.com/sunshower-io/updraft/common"
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
                common.PARSING, 
                token, 
                p, 
                tokens.IDENTIFIER_UNDEFINED,
            )
            entry, _ = symbolTables.EnterLocal(name)
        }
        
        rootNode = p.ExecutionModelFactory.NewNode(ir.VARIABLE, token)
        rootNode.SetValue(name)
        rootNode.SetAttribute(ir.ID, name)
        entry.AddLine(&ir.Line{
            Number:token.GetLineNumber(),
        })
        token, _ = p.NextToken()
    case core.BOOLEAN_TOKEN:
        rootNode = p.ExecutionModelFactory.NewNode(ir.BOOLEAN, token)
        rootNode.SetValue(token.GetValue())
        token, _ = p.NextToken()
    case tokens.INTEGER:
        rootNode = p.ExecutionModelFactory.NewNode(ir.INTEGER, token)
        rootNode.SetValue(token.GetValue())
        token, _ = p.NextToken()
    case tokens.REAL:
        rootNode = p.ExecutionModelFactory.NewNode(ir.FLOAT, token)
        rootNode.SetValue(token.GetValue())
        token, _ = p.NextToken()
        
    case tokens.STRING:
        rootNode = p.ExecutionModelFactory.NewNode(ir.STRING_LITERAL, token)
        rootNode.SetValue(token.GetValue())
        token, _ = p.NextToken()
        
    case tokens.NOT:
        token, _ = p.NextToken()
        rootNode = p.ExecutionModelFactory.NewNode(ir.NOT, token)
        
        rootNode.AddChild(p.parseFactor(token))
        
    case tokens.LPAREN:
        token, _ = p.NextToken()
        rootNode, _ = p.parseExpression(token)
        token, _ = p.CurrentToken()
        
        if token.GetType() == tokens.RPAREN {
            token, _ = p.NextToken()
        } else {
            p.ErrorHandler.FlagError(
                common.PARSING, 
                token,
                p, 
                tokens.UNEXPECTED_TOKEN,
            )
        }

    default:
        p.ErrorHandler.FlagError(
            common.PARSING,
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
    
    for {
        nodeType, ok := multiplicativeTypes[tokenType]
        if !ok {
            break
        }
        operandNode := p.ExecutionModelFactory.NewNode(nodeType, token)
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
        negate := p.ExecutionModelFactory.NewNode(ir.NEGATE, token)
        negate.AddChild(rootNode)
        rootNode = negate
    }
    
    
    token, _ = p.CurrentToken()
    
    tokenType = token.GetType()
    
   
    
    for  {
        nodeType, ok := additiveTypes[tokenType]
        
        if !ok {
            break
        }
        
        operandNode := p.ExecutionModelFactory.NewNode(nodeType, token)
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
        operandNode := p.ExecutionModelFactory.NewNode(irNodeType, token)
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
    
    return p.parseExpression(token)
}


func getMultiplicativeTypes() OperationMap {
    result := make(OperationMap)
    result[tokens.STAR] = ir.MULTIPLY
    result[tokens.SLASH] = ir.DIVIDE
    result[tokens.MOD] = ir.MODULO
    return result
}

func getAdditiveTypes() OperationMap {
    result := make(OperationMap)
    result[tokens.PLUS]         = ir.ADD
    result[tokens.MINUS]        = ir.SUBTRACT
    result[tokens.OR]           = ir.OR
    result[tokens.AND]          = ir.AND
    return result
}



func relationalTypes() OperationMap {
    result := make(OperationMap)
    result[tokens.EQUALS] = ir.EQUAL_TO
    result[tokens.NOT_EQUALS] = ir.NOT_EQUAL_TO
    result[tokens.LT] = ir.LESS_THAN
    result[tokens.GT] = ir.GREATER_THAN
    result[tokens.LTE] = ir.LTE
    result[tokens.GTE] = ir.GTE
    return result
}
