package handlers

import (
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/ast"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func FormatDocument(
	d *sshconfig.SSHDocument,
	textRange protocol.Range,
	options protocol.FormattingOptions,
) ([]protocol.TextEdit, error) {
	edits := make([]protocol.TextEdit, 0)

	it := d.Config.Options.Iterator()
	for it.Next() {
		line := it.Key().(uint32)
		entry := it.Value().(ast.SSHEntry)

		if !(line >= textRange.Start.Line && line <= textRange.End.Line) {
			continue
		}

		switch entry.GetType() {
		case ast.SSHTypeOption:
			option := entry.(*ast.SSHOption)
			edits = append(edits, formatSSHOption(
				option,
				options,
				optionTemplate,
			)...)
		case ast.SSHTypeMatch:
			matchBlock := entry.(*ast.SSHMatchBlock)

			edits = formatSSHMatchBlock(textRange, matchBlock, options)
		case ast.SSHTypeHost:
			hostBlock := entry.(*ast.SSHHostBlock)

			edits = formatSSHHostBlock(textRange, hostBlock, options)
		}
	}

	return edits, nil
}

