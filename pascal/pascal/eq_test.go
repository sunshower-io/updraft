package pascal

import "testing"

func TestSimpleEqualityBetweenTwoEqualIntegersWorks(t *testing.T) {
    prg := `
    BEGIN
        a := 1 = 1;
    END.
    `
    expectValue(t, prg, "a", true)
}

func TestSimpleInequalityBetweenTwoNonEqualIntegersFails(t *testing.T) {
    prg := `
    
    BEGIN
        a := 1 <> 1;
    END.
    `
    
    expectValue(t, prg, "a", false)
}

func TestOperatorPrecedenceIsCorrectForEqualityAndInequality(t *testing.T) {
    
    prg := `
    
    BEGIN
        a :=  (1 = 1) or true;
    END.
    `
    
    //expectValue(t, prg, "a", true)
    
    printTree(prg)
}

func TestBooleanOperatorPrecedenceIsHigherThanRelational(t *testing.T) {
    prg := `
    
    BEGIN
        b := (1 <> 1) or true and false;
    END.
    `
  
    expectValue(t, prg, "b", false)
    
}