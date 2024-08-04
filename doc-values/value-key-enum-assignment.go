package docvalues

import (
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type KeyEnumAssignmentValue struct {
	Values map[EnumString]Value
	Separator string
	ValueIsOptional bool
}

func (v KeyEnumAssignmentValue) GetTypeDescription() []string {
	if len(v.Values) == 1 {
		firstKey := utils.KeysOfMap(v.Values)[0]
		valueDescription := v.Values[firstKey].GetTypeDescription()

		if (len(valueDescription) == 1) {
			return []string{
				fmt.Sprintf("Key-Value pair in form of '<%s>%s<%s>'", firstKey.DescriptionText, v.Separator, valueDescription[0]),
			}
		}
	} 


	var result []string
	for key, value := range v.Values {
		result = append(result, key.Documentation)
		result = append(result, value.GetTypeDescription()...)
	}

	return append([]string{
		"Key-Value pair in form of 'key%svalue'", v.Separator,
	}, result...)
}

func (v KeyEnumAssignmentValue) getValue(findKey string) (*Value, bool) {
	for key, value := range v.Values {
		if key.InsertText == findKey {
			switch value.(type) {
				case CustomValue:
					customValue := value.(CustomValue)
					context := KeyValueAssignmentContext{
						SelectedKey: findKey,
					}

					fetchedValue := customValue.FetchValue(context)

					return &fetchedValue, true
				default:
					return &value, true
			}
		}
	}

	return nil, false
}

func (v KeyEnumAssignmentValue) CheckIsValid(value string) error {
	parts := strings.Split(value, v.Separator)

	if len(parts) == 0 || parts[0] == "" {
		// Nothing to check for
		return nil
	}

	if len(parts) != 2 {
		if v.ValueIsOptional {
			return nil
		}

		return KeyValueAssignmentError{}
	}

	checkValue, found := v.getValue(parts[0])

	if !found {
		return ValueNotInEnumError{
			AvailableValues: utils.Map(utils.KeysOfMap(v.Values), func(key EnumString) string { return key.InsertText }),
			ProvidedValue: parts[0],
		}
	}

	err := (*checkValue).CheckIsValid(parts[1])

	if err != nil {
		return err
	}

	return nil
}

func (v KeyEnumAssignmentValue) FetchEnumCompletions() []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)

	for enumKey := range v.Values {
		textFormat := protocol.InsertTextFormatPlainText
		kind := protocol.CompletionItemKindEnum

		completions = append(completions, protocol.CompletionItem{
			Label:            enumKey.InsertText,
			InsertTextFormat: &textFormat,
			Kind:             &kind,
			Documentation:    &enumKey.Documentation,
		})
	}

	return completions
}

func (v KeyEnumAssignmentValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	if cursor == 0 {
		return v.FetchEnumCompletions()
	}

	relativePosition, found := utils.FindPreviousCharacter(line, v.Separator, int(cursor-1))

	if found {
		selectedKey := line[:uint32(relativePosition)]
		line = line[uint32(relativePosition+len(v.Separator)):]
		cursor -= uint32(relativePosition)

		keyValue, found := v.getValue(selectedKey)

		if !found {
			// Hmm... weird
			return v.FetchEnumCompletions()
		}

		return (*keyValue).FetchCompletions(line, cursor)
	} else {
		return v.FetchEnumCompletions()
	}
}