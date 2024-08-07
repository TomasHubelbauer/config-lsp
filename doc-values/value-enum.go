package docvalues

import (
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type ValueNotInEnumError struct {
	AvailableValues []string
	ProvidedValue   string
}

func (e ValueNotInEnumError) Error() string {
	if len(e.AvailableValues) <= 6 {
		return fmt.Sprintf("This value is not valid. Select one from: %s", strings.Join(e.AvailableValues, ","))
	} else {
		return fmt.Sprintf("This value is not valid")
	}
}

type EnumString struct {
	// What is actually inserted into the document
	InsertText string
	// What is shown in the completion list
	DescriptionText string
	// Documentation for this value
	Documentation string
}

func (v EnumString) ToCompletionItem() protocol.CompletionItem {
	textFormat := protocol.InsertTextFormatPlainText
	kind := protocol.CompletionItemKindEnum
	return protocol.CompletionItem{
		Label:            v.InsertText,
		InsertTextFormat: &textFormat,
		Kind:             &kind,
		Documentation:    &v.Documentation,
	}
}

func CreateEnumString(value string) EnumString {
	return EnumString{
		InsertText:      value,
		DescriptionText: value,
	}
}

func CreateEnumStringWithDoc(value string, doc string) EnumString {
	return EnumString{
		InsertText:      value,
		DescriptionText: value,
		Documentation:   doc,
	}
}

type EnumValue struct {
	Values []EnumString
	// If `true`, the value MUST be one of the values in the Values array
	// Otherwise an error is shown
	// If `false`, the value is just a hint
	EnforceValues bool
}

func (v EnumValue) GetTypeDescription() []string {
	if len(v.Values) == 1 {
		return []string{"'" + v.Values[0].DescriptionText + "'"}
	}

	lines := make([]string, len(v.Values)+1)
	lines[0] = "Enum of:"

	for index, value := range v.Values {
		lines[index+1] += "\t* " + value.DescriptionText
	}

	return lines
}
func (v EnumValue) CheckIsValid(value string) []*InvalidValue {
	if !v.EnforceValues {
		return nil
	}

	for _, validValue := range v.Values {
		if validValue.InsertText == value {
			return nil
		}

	}

	return []*InvalidValue{
		{
			Err: ValueNotInEnumError{
				ProvidedValue:   value,
				AvailableValues: utils.Map(v.Values, func(value EnumString) string { return value.InsertText }),
			},
			Start: 0,
			End:   uint32(len(value)),
		},
	}
}
func (v EnumValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, len(v.Values))

	for index, value := range v.Values {
		textFormat := protocol.InsertTextFormatPlainText
		kind := protocol.CompletionItemKindEnum

		completions[index] = protocol.CompletionItem{
			Label:            value.InsertText,
			InsertTextFormat: &textFormat,
			Kind:             &kind,
			Documentation:    &value.Documentation,
		}
	}

	return completions
}
