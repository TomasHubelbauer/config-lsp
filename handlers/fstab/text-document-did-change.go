package fstab

import (
	"config-lsp/common"
	"config-lsp/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidChange(
	context *glsp.Context,
	params *protocol.DidChangeTextDocumentParams,
) error {
	content := params.ContentChanges[0].(protocol.TextDocumentContentChangeEventWhole).Text
	common.ClearDiagnostics(context, params.TextDocument.URI)

	parser := documentParserMap[params.TextDocument.URI]
	parser.Clear()

	diagnostics := make([]protocol.Diagnostic, 0)
	errors := parser.ParseFromContent(content)

	if len(errors) > 0 {
		diagnostics = append(diagnostics, utils.Map(
			errors,
			func(err common.ParseError) protocol.Diagnostic {
				return err.ToDiagnostic()
			},
		)...)
	}

	diagnostics = append(diagnostics, parser.AnalyzeValues()...)

	if len(diagnostics) > 0 {
		common.SendDiagnostics(context, params.TextDocument.URI, diagnostics)
	}

	return nil
}
