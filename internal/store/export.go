package store

import (
	"encoding/json"
	"example.com/receiptcheck/internal/model"
	"fmt"
	"os"
)

type Snapshot struct {
	Receipts []model.Receipt         `json:"receipts"`
	Batches  []model.ValidationBatch `json:"batches"`
	Rules    []model.ValidationRule  `json:"rules"`
	Audits   []model.AuditEvent      `json:"audits"`
}

func (s *Store) Snapshot() (Snapshot, error) {
	rs, e := s.ListReceipts()
	if e != nil {
		return Snapshot{}, e
	}
	as, e := s.ListAudits()
	if e != nil {
		return Snapshot{}, e
	}
	return Snapshot{Receipts: rs, Audits: as}, nil
}
func (s *Store) Export(path string) error {
	snap, e := s.Snapshot()
	if e != nil {
		return e
	}
	b, e := json.MarshalIndent(snap, "", "  ")
	if e != nil {
		return e
	}
	return os.WriteFile(path, b, 0644)
}
func ImportSnapshot(path string) (Snapshot, error) {
	b, e := os.ReadFile(path)
	if e != nil {
		return Snapshot{}, e
	}
	var snap Snapshot
	if e = json.Unmarshal(b, &snap); e != nil {
		return Snapshot{}, e
	}
	return snap, nil
}
func SnapshotSummary(s Snapshot) string {
	return fmt.Sprintf("receipts=%d batches=%d rules=%d audits=%d", len(s.Receipts), len(s.Batches), len(s.Rules), len(s.Audits))
}
