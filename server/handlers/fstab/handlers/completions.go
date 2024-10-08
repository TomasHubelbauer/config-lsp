package handlers

import (
	"config-lsp/doc-values"
	"config-lsp/handlers/fstab/ast"
	"config-lsp/handlers/fstab/fields"

	"github.com/tliron/glsp/protocol_3_16"
)

func GetCompletion(
	entry *ast.FstabEntry,
	cursor uint32,
) ([]protocol.CompletionItem, error) {
	targetField := entry.GetFieldAtPosition(cursor)

	switch targetField {
	case ast.FstabFieldSpec:
		value, cursor := getFieldSafely(entry.Fields.Spec, cursor)

		return fstabdocumentation.SpecField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	case ast.FstabFieldMountPoint:
		value, cursor := getFieldSafely(entry.Fields.MountPoint, cursor)

		return fstabdocumentation.MountPointField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	case ast.FstabFieldFileSystemType:
		value, cursor := getFieldSafely(entry.Fields.FilesystemType, cursor)

		return fstabdocumentation.FileSystemTypeField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	case ast.FstabFieldOptions:
		fileSystemType := entry.Fields.FilesystemType.Value.Value

		var optionsField docvalues.DeprecatedValue

		if foundField, found := fstabdocumentation.MountOptionsMapField[fileSystemType]; found {
			optionsField = foundField
		} else {
			optionsField = fstabdocumentation.DefaultMountOptionsField
		}

		value, cursor := getFieldSafely(entry.Fields.Options, cursor)

		completions := optionsField.DeprecatedFetchCompletions(
			value,
			cursor,
		)

		return completions, nil
	case ast.FstabFieldFreq:
		value, cursor := getFieldSafely(entry.Fields.Freq, cursor)

		return fstabdocumentation.FreqField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	case ast.FstabFieldPass:
		value, cursor := getFieldSafely(entry.Fields.Pass, cursor)

		return fstabdocumentation.PassField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	}

	return nil, nil
}

// Safely get value and new cursor position
// If field is nil, return empty string and 0
func getFieldSafely(field *ast.FstabField, character uint32) (string, uint32) {
	if field == nil {
		return "", 0
	}

	if field.Value.Value == "" {
		return "", 0
	}

	return field.Value.Raw, character - field.Start.Character
}
