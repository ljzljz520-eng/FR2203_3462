package audit

import (
	"example.com/receiptcheck/internal/model"
	"example.com/receiptcheck/internal/store"
	"fmt"
	"time"
)

type Logger struct{ Store *store.Store }

func NewLogger(s *store.Store) Logger { return Logger{Store: s} }
func (l Logger) Record(entity, action, detail string) (model.AuditEvent, error) {
	if l.Store == nil {
		return model.AuditEvent{}, fmt.Errorf("store unavailable")
	}
	v := model.NewAudit(fmt.Sprintf("audit-%d", time.Now().UnixNano()), entity, action, detail)
	return v, l.Store.PutAudit(v)
}
func (l Logger) RecordValidation(id string, passed bool) error {
	state := "failed"
	if passed {
		state = "passed"
	}
	_, e := l.Record(id, "validation", state)
	return e
}
func Summarize(events []model.AuditEvent) map[string]int {
	m := map[string]int{}
	for _, e := range events {
		m[e.Action]++
	}
	return m
}
