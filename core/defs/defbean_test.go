package defs

import (
	"testing"

	"github.com/zhangga/luban-go/core/rawdefs"
)

func TestFieldGroupFilter(t *testing.T) {
	f := NewDefField(&rawdefs.RawField{
		Name:   "ClientOnly",
		Groups: []string{"client"},
	})

	if !f.NeedExport([]string{"client"}) {
		t.Errorf("Expected NeedExport true for matching group")
	}

	if f.NeedExport([]string{"server"}) {
		t.Errorf("Expected NeedExport false for mismatching group")
	}

	if !f.NeedExport([]string{}) {
		t.Errorf("Expected NeedExport true when target groups is empty")
	}

	if !f.NeedExport([]string{"client", "server"}) {
		t.Errorf("Expected NeedExport true for partially matching group list")
	}

	fAll := NewDefField(&rawdefs.RawField{
		Name:   "All",
		Groups: []string{}, // No groups defined
	})

	if !fAll.NeedExport([]string{"client"}) {
		t.Errorf("Expected NeedExport true when field has no groups")
	}
}
