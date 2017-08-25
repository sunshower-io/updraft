package pascal

import "testing"

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