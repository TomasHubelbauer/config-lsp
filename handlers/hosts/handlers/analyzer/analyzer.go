package analyzer

import (
	"config-lsp/common"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Analyze(parser *HostsParser) []protocol.Diagnostic {
	errors := analyzeEntriesSetCorrectly(*parser)

	if len(errors) > 0 {
		return utils.Map(
			errors,
			func(err common.LSPError) protocol.Diagnostic {
				return err.ToDiagnostic()
			},
		)
	}

	errors = analyzeEntriesAreValid(*parser)

	if len(errors) > 0 {
		return utils.Map(
			errors,
			func(err common.LSPError) protocol.Diagnostic {
				return err.ToDiagnostic()
			},
		)
	}

	errors = append(errors, analyzeDoubleIPs(parser)...)
	errors = append(errors, analyzeDoubleHostNames(parser)...)

	return utils.Map(
		errors,
		func(err common.LSPError) protocol.Diagnostic {
			return err.ToDiagnostic()
		},
	)
}
