package model

import "time"

type Receipt struct {
	ID, FileName, Hash, Currency, Status string
	Amount                               float64
	Errors                               []string
	CreatedAt                            time.Time
}
type ValidationBatch struct {
	ID                    string
	ReceiptIDs            []string
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
	if r.IsValid() {
		b.Passed++
	} else {
		b.Failed++
	}
}
func NewRule(id, name, pattern, severity string, enabled bool) ValidationRule {
	return ValidationRule{ID: id, Name: name, Pattern: pattern, Severity: severity, Enabled: enabled}
}
func NewAudit(id, eid, action, detail string) AuditEvent {
	return AuditEvent{ID: id, EntityID: eid, Action: action, Detail: detail, CreatedAt: time.Now().UTC()}
}
