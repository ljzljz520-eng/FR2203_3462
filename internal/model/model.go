package model

import (
	"strings"
	"time"
)

type Receipt struct {
	ID, FileName, Hash, Currency, Status string
	Amount                               float64
	Errors                               []string
	CreatedAt                            time.Time
}
type ValidationBatch struct {
	ID                    string
	ReceiptIDs            []string
	Receipts              []Receipt
	Total, Passed, Failed int
	CreatedAt             time.Time
}
type ValidationRule struct {
	ID, Name, Pattern, Severity string
	Enabled                     bool
}
type AuditEvent struct {
	ID, EntityID, Action, Detail string
	CreatedAt                    time.Time
}

func NewReceipt(id, name string, amount float64) Receipt {
	return Receipt{ID: id, FileName: name, Amount: amount, Currency: "CNY", Status: "pending", CreatedAt: time.Now().UTC()}
}
func (r Receipt) IsValid() bool { return r.Status == "passed" && len(r.Errors) == 0 }
func (r *Receipt) MarkPassed()  { r.Status = "passed"; r.Errors = nil }
func (r *Receipt) MarkFailed(errs []string) {
	r.Status = "failed"
	r.Errors = append([]string(nil), errs...)
}
func NewBatch(id string) ValidationBatch { return ValidationBatch{ID: id, CreatedAt: time.Now().UTC()} }
func (b *ValidationBatch) Add(r Receipt) {
	b.Total++
	b.ReceiptIDs = append(b.ReceiptIDs, r.ID)
	b.Receipts = append(b.Receipts, r)
	if r.IsValid() {
		b.Passed++
	} else {
		b.Failed++
	}
}

// ErrorLines formats each failed receipt as "id: error" entries, prefixed
// with the file name so individual failures stay identifiable in a batch.
func ErrorLines(rs []Receipt) []string {
	out := make([]string, 0, len(rs))
	for _, r := range rs {
		if r.IsValid() {
			continue
		}
		name := r.ID
		if r.FileName != "" {
			name = r.FileName
		}
		msg := strings.Join(r.Errors, ", ")
		if msg == "" {
			msg = "validation failed"
		}
		out = append(out, name+": "+msg)
	}
	return out
}
func NewRule(id, name, pattern, severity string, enabled bool) ValidationRule {
	return ValidationRule{ID: id, Name: name, Pattern: pattern, Severity: severity, Enabled: enabled}
}
func NewAudit(id, eid, action, detail string) AuditEvent {
	return AuditEvent{ID: id, EntityID: eid, Action: action, Detail: detail, CreatedAt: time.Now().UTC()}
}
