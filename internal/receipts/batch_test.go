package receipts

import (
	"example.com/receiptcheck/internal/config"
	"example.com/receiptcheck/internal/store"
	"os"
	"path/filepath"
	"testing"
)

func TestWorkflowTwo(t *testing.T) {
	d := t.TempDir()
	p := filepath.Join(d, "ok")
	os.WriteFile(p, []byte("RECEIPT\nAMOUNT=1"), 0644)
	s, _ := store.Open(filepath.Join(d, "db"))
	defer s.Close()
	b, _ := NewService(s, config.Default().Rules).ValidateBatch([]string{p})
	if b.Passed != 1 || b.Failed != 0 {
		t.Fatal(b)
	}
}
