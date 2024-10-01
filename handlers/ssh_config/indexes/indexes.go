package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
)

type ValidPath string

func (v ValidPath) AsURI() string {
	return "file://" + string(v)
}

type SSHIndexIncludeValue struct {
	common.LocationRange
	Value string

	// Actual valid paths, these will be set by the analyzer
	Paths []ValidPath
}

type SSHIndexIncludeLine struct {
	Values []*SSHIndexIncludeValue
	Option *ast.SSHOption
	Block  ast.SSHBlock
}

type SSHIndexIgnoredUnknowns struct {
	OptionValue    *ast.SSHOption
	IgnoredOptions map[fields.NormalizedOptionName]struct{}
}

type SSHIndexes struct {
	AllOptionsPerName map[fields.NormalizedOptionName](map[ast.SSHBlock]([]*ast.SSHOption))

	Includes []*SSHIndexIncludeLine

	BlockRanges map[uint32]ast.SSHBlock

	// Map of <block|nil (for global)> to a list of ignored options
	IgnoredOptions map[ast.SSHBlock]SSHIndexIgnoredUnknowns
}
