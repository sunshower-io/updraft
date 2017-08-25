package pascal

import (
    "testing"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/magiconair/properties/assert"
)

func TestSimpleAdditionWithVariableWorks(t *testing.T) {
    prg := `
    BEGIN
        a := 2;
        b := (a + a) * a;
        c := a * b + a;
    END.
    `
    printTree(prg)
    //model := compile(prg)
    //st := model.GetSymbolTables().Peek()
    //symbol, _ := st.Lookup("b")
    //assert.Equal(t, symbol.GetAttribute(ir.DATA_VALUE), int64(8))
    //symbol, _ = st.Lookup("a")
    //assert.Equal(t, symbol.GetAttribute(ir.DATA_VALUE), int64(2))
    //
    //symbol, _ = st.Lookup("c")
    //assert.Equal(t, symbol.GetAttribute(ir.DATA_VALUE), int64(18))
}

func TestNegationOfSimpleIntegerWorks(t *testing.T) {
    
    prg := `
    BEGIN
        a := -1;
    END.
    `
    
    model := compile(prg)
    st := model.GetSymbolTables().Peek()
    symbol, _ := st.Lookup("a")
    assert.Equal(t, symbol.GetAttribute(ir.DATA_VALUE), int64(-1))
}

func TestEmptyCompoundStatementWithSimpleAssignmentWorks(t *testing.T) {
    prg := `
    BEGIN
        a := 1;
    END.
    `
    
    model := compile(prg)
    st := model.GetSymbolTables().Peek()
    symbol, _ := st.Lookup("a")
    assert.Equal(t, symbol.GetAttribute(ir.DATA_VALUE), int64(1))
}

func TestEmptyCompoundStatementWithSimpleBinaryExpressionAssignmentWorks(t *testing.T) {
    
    prg := `
    BEGIN
        a := 1 + 2;
    END.
    `
    
    model := compile(prg)
    st := model.GetSymbolTables().Peek()
    symbol, _ := st.Lookup("a")
    assert.Equal(t, symbol.GetAttribute(ir.DATA_VALUE), int64(3))
}

func TestNestedEmptyNestedCompoundStatementsWorks(t *testing.T) {
    prg := `
    
    BEGIN
        BEGIN
        END;
    END.
    
`
    model := printTree(prg)
    
    node, _ := ir.PathBy(ir.Index()).To("/0").Traverse(model.GetRoot())
    
    assert.Equal(t, node.GetType(), ir.SCOPE)
    
    node, _ = ir.PathBy(ir.Index()).To("/").Traverse(model.GetRoot())
    assert.Equal(t, node, model.GetRoot())
}


func TestNestedAssignmentsWithAllProductionsInAllPositions(t *testing.T) {
    prg := `
    
    
    BEGIN
        a := 15;
        b := -(1 * 1) + a;
        BEGIN {conversions}
            five := -1 + 2 - 3 + 4 + 3;
            ratio := five / 9.0;
            fahrenheit := 72;
            centigrade := (fahrenheit - 32) * ratio;
            centigrade := 25;
            fahrenheit := centigrade / ratio + 32;
            centigrade := 25;
            fahrenheit := 32 + centigrade / ratio;
    
        END;
        dze := fahrenheit / (ratio - ratio);
        BEGIN {calculate a square root using Newton's method}
            number := 2;
            root := number;
            root := (number / root + root) / 2;
            root := -(number / root + root) / 2;
        END;
        ch := 'x';
        str := 'hello, world';
    END.
    `
    
    printTree(prg)
}
