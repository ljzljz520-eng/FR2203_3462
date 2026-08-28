package store

import (
	"example.com/receiptcheck/internal/model"
	"path/filepath"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	s, e := Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if e = s.PutRule(model.NewRule("r", "required", "x", "error", true)); e != nil {
		t.Fatal(e)
	}
}
