package tree

import (
	"config-lsp/common"
)

// Procedure
// Save options in fields
// Each type, such as Include, User etc is own type
// Each type inherits interface
// This interface has methods such as:
// - CheckIsUser()
// For user, this checks if the user is listed in passwd
// For include, this includes the file and parses it and validates it
//

// Parse content manually as the /etc/aliases file is so simple

type AliasKey struct {
	Location common.LocationRange
	Value    string
}

type AliasValues struct {
	Location common.LocationRange
	Values   []AliasValueInterface
}

type AliasEntry struct {
	Location  common.LocationRange
	Key       *AliasKey
	Separator *common.LocationRange
	Values    *AliasValues
}

type AliasesParser struct {
	Aliases      map[uint32]*AliasEntry
	CommentLines map[uint32]struct{}
}
