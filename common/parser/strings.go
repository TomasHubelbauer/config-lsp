package parser

type ParseFeatures struct {
	ParseDoubleQuotes      bool
	ParseEscapedCharacters bool
}

var FullFeatures = ParseFeatures{
	ParseDoubleQuotes:      true,
	ParseEscapedCharacters: true,
}

type ParsedString struct {
	Raw   string
	Value string

	Features ParseFeatures
}

func ParseRawString(
	raw string,
	features ParseFeatures,
) ParsedString {
	value := raw

	// Parse double quotes
	if features.ParseDoubleQuotes {
		value = ParseDoubleQuotes(value)
	}

	// Parse escaped characters
	if features.ParseEscapedCharacters {
		value = ParseEscapedCharacters(value)
	}

	return ParsedString{
		Raw:      raw,
		Value:    value,
		Features: features,
	}
}

func ParseDoubleQuotes(
	raw string,
) string {
	value := raw
	currentIndex := 0

	for {
		start, found := findNextDoubleQuote(value, currentIndex)

		if found && start < (len(value)-1) {
			currentIndex = max(0, start-1)
			end, found := findNextDoubleQuote(value, start+1)

			if found {
				insideContent := value[start+1 : end]
				value = modifyString(value, start, end+1, insideContent)

				continue
			}
		}

		break
	}

	return value
}

func ParseEscapedCharacters(
	raw string,
) string {
	value := raw
	currentIndex := 0

	for {
		position, found := findNextEscapedCharacter(value, currentIndex)

		if found {
			currentIndex = max(0, position-1)
			escapedCharacter := value[position+1]
			value = modifyString(value, position, position+2, string(escapedCharacter))
		} else {
			break
		}
	}

	return value
}

func modifyString(
	input string,
	start int,
	end int,
	newValue string,
) string {
	return input[:start] + newValue + input[end:]
}

// Find the next non-escaped double quote in [raw] starting from [startIndex]
// When no double quote is found, return -1
// Return as the second argument whether a double quote was found
func findNextDoubleQuote(
	raw string,
	startIndex int,
) (int, bool) {
	for index := startIndex; index < len(raw); index++ {
		if raw[index] == '"' {
			if index == 0 || raw[index-1] != '\\' {
				return index, true
			}
		}
	}

	return -1, false
}

func findNextEscapedCharacter(
	raw string,
	startIndex int,
) (int, bool) {
	for index := startIndex; index < len(raw); index++ {
		if raw[index] == '\\' && index < len(raw)-1 {
			return index, true
		}
	}

	return -1, false
}
