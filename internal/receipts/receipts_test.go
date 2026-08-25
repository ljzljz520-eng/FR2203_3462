package receipts

import (
	"example.com/receiptcheck/internal/config"
	"example.com/receiptcheck/internal/store"
	"os"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	d := t.TempDir()
	p := filepath.Join(d, "ok.txt")
	os.WriteFile(p, []byte("RECEIPT\nAMOUNT=12\nCURRENCY=CNY"), 0644)
	s, _ := store.Open(filepath.Join(d, "db"))
	defer s.Close()
	r, e := NewService(s, config.Default().Rules).ValidateFile(p)
	if e != nil || !r.IsValid() {
		t.Fatalf("unexpected %#v %v", r, e)
	}
}
