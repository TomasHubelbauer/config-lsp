package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
	"errors"
	"fmt"
	"regexp"
)

var whitespacePattern = regexp.MustCompile(`\S+`)

func NewSSHIndexes() *SSHIndexes {
	return &SSHIndexes{
		AllOptionsPerName: make(map[string](map[ast.SSHBlock]([]*ast.SSHOption)), 0),
		Includes: make([]*SSHIndexIncludeLine, 0),
	}
}

func CreateIndexes(config ast.SSHConfig) (*SSHIndexes, []common.LSPError) {
	errs := make([]common.LSPError, 0)
	indexes := NewSSHIndexes()

	it := config.Options.Iterator()
	for it.Next() {
		entry := it.Value().(ast.SSHEntry)

		switch entry.GetType() {
		case ast.SSHTypeOption:
			errs = append(errs, addOption(indexes, entry.GetOption(), nil)...)
		case ast.SSHTypeMatch:
			fallthrough
		case ast.SSHTypeHost:
			block := entry.(ast.SSHBlock)

			errs = append(errs, addOption(indexes, entry.GetOption(), block)...)

			it := block.GetOptions().Iterator()
			for it.Next() {
				option := it.Value().(*ast.SSHOption)

				errs = append(errs, addOption(indexes, option, block)...)
			}
		}
	}

	// Add Includes
	for block, options := range indexes.AllOptionsPerName["Include"] {
		includeOption := options[0]
		rawValue := includeOption.OptionValue.Value.Value
		pathIndexes := whitespacePattern.FindAllStringIndex(rawValue, -1)
		paths := make([]*SSHIndexIncludeValue, 0)

		for _, pathIndex := range pathIndexes {
			startIndex := pathIndex[0]
			endIndex := pathIndex[1]

			rawPath := rawValue[startIndex:endIndex]

			offset := includeOption.OptionValue.Start.Character
			path := SSHIndexIncludeValue{
				LocationRange: common.LocationRange{
					Start: common.Location{
						Line:      includeOption.Start.Line,
						Character: uint32(startIndex) + offset,
					},
					End: common.Location{
						Line:      includeOption.Start.Line,
						Character: uint32(endIndex) + offset,
					},
				},
				Value: rawPath,
				Paths: make([]ValidPath, 0),
			}

			paths = append(paths, &path)
		}

		indexes.Includes[includeOption.Start.Line] = &SSHIndexIncludeLine{
			Values:     paths,
			Option:     includeOption,
			Block: block,
		}
	}

	return indexes, errs
}

func addOption(
	i *SSHIndexes,
	option *ast.SSHOption,
	block ast.SSHBlock,
) []common.LSPError {
	var errs []common.LSPError

	if optionsMap, found := i.AllOptionsPerName[option.Key.Key]; found {
		if options, found := optionsMap[block]; found {
			if _, duplicatesAllowed := fields.AllowedDuplicateOptions[option.Key.Key]; !duplicatesAllowed {
				firstDefinedOption := options[0]
				errs = append(errs, common.LSPError{
					Range: option.Key.LocationRange,
					Err: errors.New(fmt.Sprintf(
						"Option '%s' has already been defined on line %d",
						option.Key.Key,
						firstDefinedOption.Start.Line+1,
					)),
				})
			} else {
				i.AllOptionsPerName[option.Key.Key][block] = append(
					i.AllOptionsPerName[option.Key.Key][block],
					option,
				)
			}
		} else {
			i.AllOptionsPerName[option.Key.Key][block] = []*ast.SSHOption{
				option,
			}
		}
	} else {
		i.AllOptionsPerName[option.Key.Key] = map[ast.SSHBlock]([]*ast.SSHOption){
			block: {
				option,
			},
		}
	}

	return errs
}