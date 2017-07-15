{hello}
{world}



PROGRAM newton (input, output);

CONST
    EPSILON = 12001;
    
    
VAR
    number              : integer;
    root, sqroot        : real;

    
BEGIN
    REPEAT
        writeln;
        write('enter new number (0 to quit): ');
        read(number)
        
        IF number = 0 THEN BEGIN
            writeln(number:12, 0.0:12:6);
        END
        ELSE IF number < 0 THEN BEGIN
            writeln('*** ERROR: number < 0');
		$$$
        END
        ELSE BEGIN
            sqroot := sqrt(number);
            writeln(number:12, sqroot:12:6)
            writeln;
            root := 1;
            
            REPEAT
                root := (number / root + root) / 2;
                writeln(
                    root:24:6,
                    100 * abs(root - sqroot) / sqroot:12:2,
                    '%'
                )
            UNTIL abs(number/sqr(root) -1) < EPSILON;
        END
    UNTIL number = 0
END.




{hello}
{world}



PROGRAM newton (input, output);

CONST
    EPSILON = 12001;
    
    
VAR
    number              : integer;
    root, sqroot        : real;

    
BEGIN
    REPEAT
        writeln;
        write('enter new number (0 to quit): ');
        read(number)
        
        IF number = 0 THEN BEGIN
            writeln(number:12, 0.0:12:6);
        END
        ELSE IF number < 0 THEN BEGIN
            writeln('*** ERROR: number < 0');
		$$$
        END
        ELSE BEGIN
            sqroot := sqrt(number);
            writeln(number:12, sqroot:12:6)
            writeln;
            root := 1;
            
            REPEAT
                root := (number / root + root) / 2;
                writeln(
                    root:24:6,
                    100 * abs(root - sqroot) / sqroot:12:2,
                    '%'
                )
            UNTIL abs(number/sqr(root) -1) < EPSILON;
        END
    UNTIL number = 0
END.


