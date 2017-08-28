package pascal

import "testing"

func TestAndForSimpleBooleanVariableWorks(t *testing.T) {
    prg := 
    `
    BEGIN
        a := true
    END.
    `
    
    expectValue(t, prg, "a", true)
}


func TestConjoiningTrueBooleanLiteralsProducesTrue(t *testing.T) {
    prg := `
    BEGIN
        a := true AND true;
    END.
    `
    expectValue(t, prg, "a", true)
}

func TestComplexBooleanExpressionsWorkWithAssignments(t *testing.T) {
    prg :=  `
    BEGIN
    
        a := true;
        b := not a;
       
        d := not b;
        c := a OR b or d;
        e := c and (not c);
    END.
    `
   
    expectValue(t, prg, "c", true)
    expectValue(t, prg, "e", false)
    
}

func TestConjoiningTrueAndFalseBooleanLiteralsProducesFalse(t *testing.T) {
    prg := `
    BEGIN
        a := true AND false;
    END.
    `
    expectValue(t, prg, "a", false)
}
