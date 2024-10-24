package roothandler

import (
	"config-lsp/common"
	"fmt"

	aliases "config-lsp/handlers/aliases/lsp"
	fstab "config-lsp/handlers/fstab/lsp"
	hosts "config-lsp/handlers/hosts/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	wireguard "config-lsp/handlers/wireguard/lsp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	common.ClearDiagnostics(context, params.TextDocument.URI)

	// Find the file type
	content := params.TextDocument.Text
	language, err := initFile(
		context,
		content,
		params.TextDocument.URI,
		params.TextDocument.LanguageID,
	)

	if err != nil {
		return err
	}

	switch *language {
	case LanguageFstab:
		return fstab.TextDocumentDidOpen(context, params)
	case LanguageSSHDConfig:
		return sshdconfig.TextDocumentDidOpen(context, params)
	case LanguageSSHConfig:
		return sshconfig.TextDocumentDidOpen(context, params)
	case LanguageWireguard:
		return wireguard.TextDocumentDidOpen(context, params)
	case LanguageHosts:
		return hosts.TextDocumentDidOpen(context, params)
	case LanguageAliases:
		return aliases.TextDocumentDidOpen(context, params)
	}

	panic(fmt.Sprintf("unexpected roothandler.SupportedLanguage: %#v", language))
}

func showParseError(
	context *glsp.Context,
	uri protocol.DocumentUri,
	err common.ParseError,
) {
	context.Notify(
		"window/showMessage",
		protocol.ShowMessageParams{
			Type:    protocol.MessageTypeError,
			Message: err.Err.Error(),
		},
	)
}

func initFile(
	context *glsp.Context,
	content string,
	uri protocol.DocumentUri,
	advertisedLanguage string,
) (*SupportedLanguage, error) {
	language, err := DetectLanguage(content, advertisedLanguage, uri)

	if err != nil {
		parseError := err.(common.ParseError)
		showParseError(
			context,
			uri,
			parseError,
		)

		return nil, parseError.Err
	}

	openedFiles[uri] = struct{}{}

	// Everything okay, now we can handle the file
	rootHandler.AddDocument(uri, language)

	return &language, nil
}
