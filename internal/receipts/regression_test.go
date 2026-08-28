package receipts

import (
	"example.com/receiptcheck/internal/config"
	"example.com/receiptcheck/internal/model"
	"example.com/receiptcheck/internal/store"
	"path/filepath"
	"testing"
)

func TestBatchReadErrorReported(t *testing.T) {
	d := t.TempDir()
	s, _ := store.Open(filepath.Join(d, "db"))
	defer s.Close()
	b, _ := NewService(s, config.Default().Rules).ValidateBatch([]string{filepath.Join(d, "missing")})
	if b.Failed != 1 {
		t.Fatalf("failed=%d", b.Failed)
	}
}
func TestPersistenceSurvivesReopen(t *testing.T) {
	d := t.TempDir()
	p := filepath.Join(d, "db")
	s, _ := store.Open(p)
	r := model.NewReceipt("x", "x", 1)
	s.PutReceipt(r)
	s.Close()
	s2, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s2.Close()
	if _, e = s2.GetReceipt("x"); e != nil {
		t.Fatal(e)
	}
}
