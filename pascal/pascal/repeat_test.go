package pascal

import "testing"

func TestSimpleRepeate(t *testing.T) {
    
    prg := `
    BEGIN
   
        i := 0;
        REPEAT
            j := i;
            k := i;
        UNTIL i <= j;
    
    END.
    `
    printTree(prg)
}
