package audit

import (
	"example.com/receiptcheck/internal/store"
	"path/filepath"
	"testing"
)

func TestAuditRoundTrip(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	l := NewLogger(s)
	if _, e := l.Record("x", "create", "ok"); e != nil {
		t.Fatal(e)
	}
}
