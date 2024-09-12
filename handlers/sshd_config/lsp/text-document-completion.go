package lsp

import (
	"config-lsp/common"
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/handlers"
	"regexp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var containsSeparatorPattern = regexp.MustCompile(`\s+$`)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	line := params.Position.Line
	cursor := params.Position.Character

	d := sshdconfig.DocumentParserMap[params.TextDocument.URI]

	if _, found := d.Config.CommentLines[line]; found {
		return nil, nil
	}

	entry, matchBlock := d.Config.FindOption(line)

	if entry == nil ||
		entry.Separator == nil ||
		entry.Key == nil ||
		(common.CursorToCharacterIndex(cursor)) <= entry.Key.End.Character {
		// Empty line
		return handlers.GetRootCompletions(
			d,
			matchBlock,
			entry == nil || containsSeparatorPattern.Match([]byte(entry.Value)),
		)
	}

	if entry.Separator != nil && cursor > entry.Separator.End.Character {
		return handlers.GetOptionCompletions(
			d,
			entry,
			cursor,
		)
	}

	return nil, nil
}
