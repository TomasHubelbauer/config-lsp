package ast

import (
	"config-lsp/utils"
	"testing"
)

func TestExample1(
	t *testing.T,
) {
	input := utils.Dedent(`
LABEL=test /mnt/test ext4 defaults 0 0
`)
	c := NewFstabConfig()

	errors := c.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if c.Entries.Size() != 1 {
		t.Fatalf("Expected 1 entry, got %d", c.Entries.Size())
	}

	rawFirstEntry, _ := c.Entries.Get(uint32(0))
	firstEntry := rawFirstEntry.(FstabEntry)
	if !(firstEntry.Fields.Spec.Value.Value == "LABEL=test" && firstEntry.Fields.MountPoint.Value.Value == "/mnt/test" && firstEntry.Fields.FilesystemType.Value.Value == "ext4" && firstEntry.Fields.Options.Value.Value == "defaults" && firstEntry.Fields.Freq.Value.Value == "0" && firstEntry.Fields.Pass.Value.Value == "0") {
		t.Fatalf("Expected entry to be LABEL=test /mnt/test ext4 defaults 0 0, got %v", firstEntry)
	}

	if !(firstEntry.Fields.Spec.LocationRange.Start.Line == 0 && firstEntry.Fields.Spec.LocationRange.Start.Character == 0) {
		t.Errorf("Expected spec start to be 0:0, got %v", firstEntry.Fields.Spec.LocationRange.Start)
	}

	if !(firstEntry.Fields.Spec.LocationRange.End.Line == 0 && firstEntry.Fields.Spec.LocationRange.End.Character == 10) {
		t.Errorf("Expected spec end to be 0:10, got %v", firstEntry.Fields.Spec.LocationRange.End)
	}

	if !(firstEntry.Fields.MountPoint.LocationRange.Start.Line == 0 && firstEntry.Fields.MountPoint.LocationRange.Start.Character == 11) {
		t.Errorf("Expected mountpoint start to be 0:11, got %v", firstEntry.Fields.MountPoint.LocationRange.Start)
	}

	if !(firstEntry.Fields.MountPoint.LocationRange.End.Line == 0 && firstEntry.Fields.MountPoint.LocationRange.End.Character == 20) {
		t.Errorf("Expected mountpoint end to be 0:20, got %v", firstEntry.Fields.MountPoint.LocationRange.End)
	}

	if !(firstEntry.Fields.FilesystemType.LocationRange.Start.Line == 0 && firstEntry.Fields.FilesystemType.LocationRange.Start.Character == 21) {
		t.Errorf("Expected filesystemtype start to be 0:21, got %v", firstEntry.Fields.FilesystemType.LocationRange.Start)
	}

	if !(firstEntry.Fields.FilesystemType.LocationRange.End.Line == 0 && firstEntry.Fields.FilesystemType.LocationRange.End.Character == 25) {
		t.Errorf("Expected filesystemtype end to be 0:25, got %v", firstEntry.Fields.FilesystemType.LocationRange.End)
	}

	if !(firstEntry.Fields.Options.LocationRange.Start.Line == 0 && firstEntry.Fields.Options.LocationRange.Start.Character == 26) {
		t.Errorf("Expected options start to be 0:26, got %v", firstEntry.Fields.Options.LocationRange.Start)
	}

	if !(firstEntry.Fields.Options.LocationRange.End.Line == 0 && firstEntry.Fields.Options.LocationRange.End.Character == 34) {
		t.Errorf("Expected options end to be 0:34, got %v", firstEntry.Fields.Options.LocationRange.End)
	}

	if !(firstEntry.Fields.Freq.LocationRange.Start.Line == 0 && firstEntry.Fields.Freq.LocationRange.Start.Character == 35) {
		t.Errorf("Expected freq start to be 0:35, got %v", firstEntry.Fields.Freq.LocationRange.Start)
	}

	if !(firstEntry.Fields.Freq.LocationRange.End.Line == 0 && firstEntry.Fields.Freq.LocationRange.End.Character == 36) {
		t.Errorf("Expected freq end to be 0:36, got %v", firstEntry.Fields.Freq.LocationRange.End)
	}

	if !(firstEntry.Fields.Pass.LocationRange.Start.Line == 0 && firstEntry.Fields.Pass.LocationRange.Start.Character == 37) {
		t.Errorf("Expected pass start to be 0:37, got %v", firstEntry.Fields.Pass.LocationRange.Start)
	}
}
