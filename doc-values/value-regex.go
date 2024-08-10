package docvalues

import (
	"fmt"
	"regexp"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type RegexInvalidError struct {
	Regex string
}

func (e RegexInvalidError) Error() string {
	return fmt.Sprintf("This value does not match the regular expression (Pattern: `%s`)", e.Regex)
}

type RegexValue struct {
	Regex regexp.Regexp
}

func (v RegexValue) GetTypeDescription() []string {
	return []string{
		fmt.Sprintf("String matching the regular expression (Pattern: `%s`)", v.Regex.String()),
	}
}

func (v RegexValue) CheckIsValid(value string) []*InvalidValue {
	if !v.Regex.MatchString(value) {
		return []*InvalidValue{{
			Err:   RegexInvalidError{Regex: v.Regex.String()},
			Start: 0,
			End:   uint32(len(value)),
		}}
	}

	return []*InvalidValue{}
}

func (v RegexValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return []protocol.CompletionItem{}
}

func (v RegexValue) FetchHoverInfo(line string, cursor uint32) []string {
	return []string{
		fmt.Sprintf("Pattern: `%s`", v.Regex.String()),
	}
}
