package pascal

import "testing"


func TestBooleanLiteralWorks(t *testing.T) {
    prg := `
    BEGIN
        a := false;
    END.
    `
    expectValue(t, prg, "a", false)
}

func TestBooleanLiteralTrueValueWorks(t *testing.T) {
    prg := `
    BEGIN
        a := true;
    END.
    `
    expectValue(t, prg, "a", true)
}


func TestBooleanOrWorks(t *testing.T) {
    prg := `
    BEGIN
        a := true OR false;
    END.
    `
    expectValue(t, prg, "a", true)
}

func TestBooleanAndWorks(t *testing.T) {
    prg := `
    BEGIN
        a := true AND false;
    END.
    
    `
    
    expectValue(t, prg, "a", false)
}

func TestAdditionExpressionWorks(t *testing.T) {
    prg := `
    BEGIN
        a := 1 + 1;
    END.
    `
    
    expectValue(t, prg, "a", 2)
}


func TestSubtractionExpressionWorks(t *testing.T) {
    prg := `
    BEGIN
        c := 1 - 1;
    END.
    `
    expectValue(t, prg, "c", 0)
}





func TestModuloWorks(t *testing.T) {
    prg := `
    BEGIN
        c := 10 mod 8;
    END.
    `
    expectValue(t, prg, "c", 2)
    
}