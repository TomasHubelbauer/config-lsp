package tree

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases/parser"

	"github.com/antlr4-go/antlr/v4"
)

type aliasesListenerContext struct {
	line                uint32
	currentIncludeIndex *uint32
}

type aliasesParserListener struct {
	*parser.BaseAliasesListener
	Parser       *AliasesParser
	Errors       []common.LSPError
	aliasContext aliasesListenerContext
}

func (s *aliasesParserListener) EnterEntry(ctx *parser.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	s.Parser.Aliases[location.Start.Line] = &AliasEntry{
		Location: location,
	}
}

func (s *aliasesParserListener) EnterSeparator(ctx *parser.SeparatorContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	entry := s.Parser.Aliases[location.Start.Line]
	entry.Separator = &location
}

func (s *aliasesParserListener) EnterKey(ctx *parser.KeyContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	entry := s.Parser.Aliases[location.Start.Line]
	entry.Key = &AliasKey{
		Location: location,
		Value:    ctx.GetText(),
	}
}

func (s *aliasesParserListener) EnterValues(ctx *parser.ValuesContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	entry := s.Parser.Aliases[location.Start.Line]
	entry.Values = &AliasValues{
		Location: location,
		Values:   make([]AliasValueInterface, 0, 5),
	}
}

// === Value === //

func (s *aliasesParserListener) EnterUser(ctx *parser.UserContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	user := AliasValueUser{
		AliasValue: AliasValue{
			Location: location,
			Value:    ctx.GetText(),
		},
	}

	entry := s.Parser.Aliases[location.Start.Line]
	entry.Values.Values = append(entry.Values.Values, user)
}

func (s *aliasesParserListener) EnterFile(ctx *parser.FileContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	if s.aliasContext.currentIncludeIndex != nil {
		// This `file` is inside an `include`, so we need to set the path on the include
		values := s.Parser.Aliases[location.Start.Line].Values
		rawValue := values.Values[*s.aliasContext.currentIncludeIndex]

		// Set the path
		include := rawValue.(AliasValueInclude)
		include.Path = AliasValueIncludePath{
			Location: location,
			Path:     path(ctx.GetText()),
		}
		values.Values[*s.aliasContext.currentIncludeIndex] = include

		// Clean up
		s.aliasContext.currentIncludeIndex = nil

		return
	}

	// Raw file, process it
	file := AliasValueFile{
		AliasValue: AliasValue{
			Location: location,
			Value:    ctx.GetText(),
		},
		Path: path(ctx.GetText()),
	}

	entry := s.Parser.Aliases[location.Start.Line]
	entry.Values.Values = append(entry.Values.Values, file)
}

func (s *aliasesParserListener) EnterCommand(ctx *parser.CommandContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	command := AliasValueCommand{
		AliasValue: AliasValue{
			Location: location,
			Value:    ctx.GetText(),
		},
		Command: ctx.GetText()[1:],
	}

	entry := s.Parser.Aliases[location.Start.Line]
	entry.Values.Values = append(entry.Values.Values, command)
}

func (s *aliasesParserListener) EnterInclude(ctx *parser.IncludeContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	include := AliasValueInclude{
		AliasValue: AliasValue{
			Location: location,
			Value:    ctx.GetText(),
		},
	}

	entry := s.Parser.Aliases[location.Start.Line]
	entry.Values.Values = append(entry.Values.Values, include)

	index := uint32(len(entry.Values.Values) - 1)
	s.aliasContext.currentIncludeIndex = &index
}

func (s *aliasesParserListener) EnterEmail(ctx *parser.EmailContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.aliasContext.line)

	email := AliasValueEmail{
		AliasValue: AliasValue{
			Location: location,
			Value:    ctx.GetText(),
		},
	}

	entry := s.Parser.Aliases[location.Start.Line]
	entry.Values.Values = append(entry.Values.Values, email)
}

func createListener(
	parser *AliasesParser,
	line uint32,
) aliasesParserListener {
	return aliasesParserListener{
		Parser: parser,
		aliasContext: aliasesListenerContext{
			line: line,
		},
		Errors: make([]common.LSPError, 0),
	}
}

type errorListener struct {
	*antlr.DefaultErrorListener
	Errors  []common.LSPError
	context aliasesListenerContext
}

func (d *errorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	_ int,
	character int,
	message string,
	error antlr.RecognitionException,
) {
	line := d.context.line
	d.Errors = append(d.Errors, common.LSPError{
		Range: common.CreateSingleCharRange(uint32(line), uint32(character)),
		Err: common.SyntaxError{
			Message: message,
		},
	})
}

func createErrorListener(
	line uint32,
) errorListener {
	return errorListener{
		Errors: make([]common.LSPError, 0),
		context: aliasesListenerContext{
			line: line,
		},
	}
}
