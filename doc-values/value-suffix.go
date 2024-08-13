package docvalues

import (
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Suffix struct {
	Suffix  string
	Meaning string
}

type SuffixWithMeaningValue struct {
	Suffixes []Suffix
	SubValue Value
}

func (v SuffixWithMeaningValue) GetTypeDescription() []string {
	subDescription := v.SubValue.GetTypeDescription()

	suffixDescription := utils.Map(v.Suffixes, func(suffix Suffix) string {
		return fmt.Sprintf("_%s_ -> %s", suffix.Suffix, suffix.Meaning)
	})

	return append(subDescription,
		append(
			[]string{"The following suffixes are allowed:"},
			suffixDescription...,
		)...,
	)
}

func (v SuffixWithMeaningValue) CheckIsValid(value string) []*InvalidValue {
	for _, suffix := range v.Suffixes {
		if strings.HasSuffix(value, suffix.Suffix) {
			return v.SubValue.CheckIsValid(value[:len(value)-len(suffix.Suffix)])
		}
	}

	return v.SubValue.CheckIsValid(value)
}

func (v SuffixWithMeaningValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	textFormat := protocol.InsertTextFormatPlainText
	kind := protocol.CompletionItemKindText

	suffixCompletions := utils.Map(v.Suffixes, func(suffix Suffix) protocol.CompletionItem {
		return protocol.CompletionItem{
			Label:            suffix.Suffix,
			Detail:           &suffix.Meaning,
			InsertTextFormat: &textFormat,
			Kind:             &kind,
		}
	})

	return append(suffixCompletions, v.SubValue.FetchCompletions(line, cursor)...)
}

func (v SuffixWithMeaningValue) FetchHoverInfo(line string, cursor uint32) []string {
	for _, suffix := range v.Suffixes {
		if strings.HasSuffix(line, suffix.Suffix) {
			return append([]string{
				fmt.Sprintf("Suffix: _%s_ -> %s", suffix.Suffix, suffix.Meaning),
			},
				v.SubValue.FetchHoverInfo(line[:len(line)-len(suffix.Suffix)], cursor)...,
			)
		}
	}

	return v.SubValue.FetchHoverInfo(line, cursor)
}
