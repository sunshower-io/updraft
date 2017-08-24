package common

import (
	"fmt"
	"github.com/docker/docker/api/types/time"
	"github.com/sunshower-io/updraft/common/io"
	"github.com/sunshower-io/updraft/common/observer"
	front "github.com/sunshower-io/updraft/common/frontend"
    "github.com/sunshower-io/updraft/backends/common"
)

const PARSER_SUMMARY_FORMAT = "\n%20d source lines." +
	"\n%20d syntax errors." +
	"\n%s seconds total parsing time"

type CommentMessageListener struct {
	observer.EventListener
	Comments []string
}

func (s *CommentMessageListener) Id() string {
	return "comment-scanned"
}

func (s *CommentMessageListener) ListensFor(e observer.Message) bool {
	return true
}

func (s *CommentMessageListener) OnMessage(m observer.Message) {
	if s.Comments == nil {
		s.Comments = make([]string, 0)
	}
	s.Comments = append(s.Comments, m.GetBody().(string))
}

type SourceMessageListener struct {
	observer.EventListener
}

func (s *SourceMessageListener) Id() string {
	return "source-message-listener"
}

func (s *SourceMessageListener) ListensFor(e observer.Message) bool {
	return true
}

func (s *SourceMessageListener) OnMessage(m observer.Message) {
	e := m.TopicId()

	switch e {
	case observer.SOURCE_LINE:
		s.printSourceLine(m.GetBody())
	case observer.SOURCE_LINE_FORMAT:
		s.formatSourceLine(m.GetBody())
	}
}

func (s *SourceMessageListener) formatSourceLine(body interface{}) {
	s.printSourceLine(body)
}

func (s *SourceMessageListener) printSourceLine(body interface{}) {
	switch body.(type) {
	case io.SourcePosition:
		b := body.(io.SourcePosition)
		println(fmt.Sprintf("%03d %s", b.LineNumber, b.Line))

	case *io.SourcePosition:
		b := body.(*io.SourcePosition)
		println(fmt.Sprintf("%03d %s", b.LineNumber, b.Line))
	}
}

type ParserMessageListener struct {
	observer.EventListener
	
	DumpSymbols  bool
}

func (s *ParserMessageListener) Id() string {
	return "parser-message-listener"
}

func (s *ParserMessageListener) ListensFor(m observer.Message) bool {
	return true
}


func printParserSummary(
		summary *front.ParserSummary,
		dump bool,
) {
	if dump {
		println(summary.SymbolTables.String())
	}
	println(fmt.Sprintf(
		PARSER_SUMMARY_FORMAT,
		summary.LineNumber,
		summary.ErrorCount,
		time.DurationToSecondsString(summary.Time),
	))
}

func (s *ParserMessageListener) OnMessage(m observer.Message) {
	

	body := m.GetBody()
	switch body.(type) {
	case front.ParserSummary:
		b := body.(front.ParserSummary)
		
		printParserSummary(&b, s.DumpSymbols)

	case *front.ParserSummary:
		b := body.(*front.ParserSummary)
		printParserSummary(b, s.DumpSymbols)
	}
}

type BackendMessageListener struct {
	observer.EventListener
}

func (s *BackendMessageListener) Id() string {
	return "parser-message-listener"
}

func (s *BackendMessageListener) ListensFor(m observer.Message) bool {
	return true
}

func (s *BackendMessageListener) OnMessage(m observer.Message) {
	switch m.TopicId() {
	case observer.COMPILER_SUMMARY:

		panic("noe")

	case observer.INTERPRETER_SUMMARY:

		var summary *common.Summary
		b := m.GetBody()

		switch b.(type) {
		case common.Summary:
			u := b.(common.Summary)
            
			summary = &u
		case *common.Summary:
			summary = b.(*common.Summary)

		}
		printInterpreterSummary(summary)
	}

}


func printInterpreterSummary(summary *common.Summary) {

}
