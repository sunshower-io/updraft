package pascal

import (
	"testing"
	"strings"
	"github.com/sunshower-io/updraft/common/observer"
	"github.com/magiconair/properties/assert"
    "github.com/sunshower-io/updraft/front/parser"
    "github.com/sunshower-io/updraft/pascal/common"
    ccommon "github.com/sunshower-io/updraft/common"
    "github.com/sunshower-io/updraft/common/core"
    pir "github.com/sunshower-io/updraft/pascal/ir"
    "github.com/sunshower-io/updraft/common/ir"
)

const HELLO = `PROGRAM hello (output);

{Write 'hello, world' ten times.}

VAR

    i : integer;
BEGIN {hello}
    i := 1000e
    FOR i := 1 TO 10 DO BEGIN
    IF (+-:=<>=<==.......)
        writeln('hello, world');
    END;
END {hello}.


`

const COMPLEX = `


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
        read(number);
        
        IF number = 0 THEN BEGIN
            writeln(number:12, 0.0:12:6);
        END
        ELSE IF number < 0 THEN BEGIN
            writeln('*** ERROR: number < 0');
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





`




func TestSimpleCommentsGenerateCorrectLexingEvents(t *testing.T) {

	prog := `
	{this is a comment}
	{this is another comment}
	
	{frap}
	`
	
	cmp := NewPascal(strings.NewReader(prog))
	
	newlineListener := &countingListener{
        eventType: observer.SOURCE_LINE,
    }
	
	cmp.AddListener(
		ccommon.LEXING,
		newlineListener,
	)
	
	cmp.Compile()
	
	assert.Equal(t, newlineListener.count, 6)
}

func TestTokenLineNumberIsCorrectForEof(t *testing.T) {
    
    prog := `
    BEGIN
    
    END.
    `
    
    cmp := NewPascal(strings.NewReader(prog))
    cmp.AddListener(
        ccommon.PARSING,
        &common.ParserMessageListener{},
    )
    
    cmp.Compile()
}

func TestTokenErrorsGenerateErrorEvents(t *testing.T) {
    
    prog := `
	{this is a comment}
	{this is another comment}
	
	{frap}
	
	
	{
	        coobeanfrappers
	        
	        I'm not really sure how this works
	
	
	`
    
    cmp := NewPascal(strings.NewReader(prog))
    
    newlineListener := &countingListener{
        eventType:observer.SYNTAX_ERROR,
    }
    
    cmp.AddListener(
        ccommon.PARSING,
        newlineListener,
    )
    
    cmp.Compile()
    
    println(newlineListener.messages[0].Format())
    
    assert.Equal(t, newlineListener.count, 1)
}

func TestCommentsGenerateCommentEvents(t *testing.T) {
	
	prog := `
	{this is a comment}
	{this is another comment}
	
	{frap}
	
	
	{
	        coobeanfrappers
	        
	        I'm not really sure how this works {are you?}
	
	
	}
	`
	
	cmp := NewPascal(strings.NewReader(prog))
	
	newlineListener := &countingListener{
        eventType:observer.COMMENT,
    }
	
	cmp.AddListener(
		ccommon.LEXING,
		newlineListener,
	)
	
	cmp.Compile()
	
	assert.Equal(t, newlineListener.count, 4)
}


func TestReadingCommentConsumesTrailingBrace(t *testing.T) {
    
    prog := `
    {hello}
    `
    cmp := NewPascal(strings.NewReader(prog))
    
    newlineListener := &countingListener{
        eventType:observer.PARSER_SUMMARY,
    }
    
    cmp.AddListener(
        ccommon.PARSING,
        newlineListener,
    )
    cmp.AddListener(
        ccommon.PARSING,
        &parser.ParserMessageListener{},
    )
    cmp.Compile()
}


func TestReadingInvalidTokensProducesErrors(t *testing.T) {
    
    prog := `
    {hello}
   $$$
    `
    cmp := NewPascal(strings.NewReader(prog))
    
    newlineListener := &countingListener{
        eventType:observer.PARSER_SUMMARY,
    }
    
    cmp.AddListener(
        ccommon.PARSING,
        newlineListener,
    )
    cmp.AddListener(
        ccommon.PARSING,
        &parser.ParserMessageListener{},
    )
    cmp.Compile()
}


func TestReadingAdditiveAssignmentWithConstantsWorks(t *testing.T) {
    
    
    prg := `
    BEGIN
    
     a := -(1 + 2) * 3;
    END.
    `
    model := compile(prg).GetExecutionModel()
    result := new(ir.JsonExecutionModelPrinter).Print(model)
    println(result)
    
}



func TestReadingMultipleAssignmentsWorks(t *testing.T) {
    prg := `

    BEGIN
        firstThing := 10;
        secondThing := 20;

    END.

    `
    
    model := compile(prg).GetExecutionModel()
    
    root := model.GetRoot()
    
    assert.Equal(t, root.GetType(), pir.COMPOUND)
    
    d, _ := ir.PathBy(ir.Index()).To("/0").Traverse(root)
    assert.Equal(t, d.GetType(), ir.ASSIGN)
    
    assert.Equal(t, d.Arity(), 2)
    
    c := d.ChildAt(0)
    
    assert.Equal(t, c.GetType(), ir.VARIABLE)
    assert.Equal(t, c.GetValue(), "firstthing")
    
    c = d.ChildAt(1)
    
    assert.Equal(t, c.GetType(), ir.INTEGER)
    assert.Equal(t, c.GetValue(), int64(10))
    
    d, _ = ir.PathBy(ir.Index()).To("/1").Traverse(root)
    assert.Equal(t, d.GetType(), ir.ASSIGN)
    
    assert.Equal(t, d.Arity(), 2)
    
    c = d.ChildAt(0)
    
    assert.Equal(t, c.GetType(), ir.VARIABLE)
    assert.Equal(t, c.GetValue(), "secondthing")
    
    c = d.ChildAt(1)
    
    assert.Equal(t, c.GetType(), ir.INTEGER)
    assert.Equal(t, c.GetValue(), int64(20))
}


func TestReadingAssignmentWorks(t *testing.T) {
    
    prg := `
    BEGIN
     a := 4
    END.
    `
    
    
    
    result := compile(prg) 
    
    model := result.GetExecutionModel()
    root := model.GetRoot()
    assert.Equal(t, root.GetType(), pir.COMPOUND)
    
    assert.Equal(t, root.Arity(), 1)
    
    child := root.GetChildren()[0]
    assert.Equal(t, child.GetType(), ir.ASSIGN)
    
    assert.Equal(t, child.Arity(), 2)
    
    lhs := child.GetChildren()[0]
    rhs := child.GetChildren()[1]
    
    assert.Equal(t, lhs.GetType(), ir.VARIABLE)
    assert.Equal(t, rhs.GetType(), ir.INTEGER)
    
    
}

func TestReadingComplexProgramWorks(t *testing.T) {
    
    cmp := NewPascal(strings.NewReader(COMPLEX))
    
    newlineListener := &countingListener{
        eventType:observer.PARSER_SUMMARY,
    }
    
    cmp.AddListener(
        ccommon.PARSING,
        newlineListener,
    )
    cmp.AddListener(
        ccommon.PARSING,
        &parser.ParserMessageListener{},
    )
    
    cmp.Compile()
  
    assert.Equal(t, newlineListener.count, 1)
    
    println(newlineListener.messages[0].Format())
}



//func TestComplexScriptReadsCorrectValues(t *testing.T) {
//
//	opts := PascalOptions{}
//
//	p := NewPascal(parser.NewSource(
//		strings.NewReader(COMPLEX),
//	), opts)
//
//	l := new(common.CommentMessageListener)
//	p.AddScannerListener(l)
//
//	p.Run()
//
//}
//
//func TestSimpleScriptReadsCorrectNumberOfComments(t *testing.T) {
//	opts := PascalOptions{}
//
//	p := NewPascal(parser.NewSource(
//		strings.NewReader(HELLO),
//	), opts)
//
//	l := new(common.CommentMessageListener)
//	p.AddScannerListener(l)
//
//	p.Run()
//
//	assert.Equal(t, len(l.Comments), 3)
//
//	assert.Equal(t, l.Comments[0], "Write 'hello, world' ten times.")
//
//}
//
//func TestDoubleQuotes(t *testing.T) {
//	opts := PascalOptions{}
//
//	p := NewPascal(parser.NewSource(
//		strings.NewReader(`hello ''world''''`),
//	), opts)
//	p.Run()
//}







type countingListener struct {
	count 		int
	eventType  	observer.EventType
    messages    []observer.Message
}



func(s *countingListener ) Id() string {
	return "newlines"
}

func(s *countingListener) ListensFor(m observer.Message) bool {
	return m.TopicId() == s.eventType
}

func (s *countingListener) OnMessage(m observer.Message) {
    s.messages = append(s.messages, m)
	s.count++
}


func printTree(prg string) ir.ExecutionModel {
    model := compile(prg).GetExecutionModel()
    str := new(ir.JsonExecutionModelPrinter).Print(model)
    println(str)
    
    return model
}


func compile(prg string) core.CompilationResult {
    
    cmp := NewPascal(strings.NewReader(prg))
    
    newlineListener := &countingListener{
        eventType:observer.PARSER_SUMMARY,
    }
    
    cmp.AddListener(
        ccommon.PARSING,
        newlineListener,
    )
    cmp.AddListener(
        ccommon.PARSING,
        &parser.ParserMessageListener{},
    )
    
    return cmp.Compile()
}
