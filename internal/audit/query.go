package audit

import "example.com/receiptcheck/internal/model"

func ForEntity(es []model.AuditEvent, id string) []model.AuditEvent {
	out := make([]model.AuditEvent, 0)
	for _, e := range es {
		if e.EntityID == id {
			out = append(out, e)
		}
	}
	return out
}
func Last(es []model.AuditEvent) model.AuditEvent {
	if len(es) == 0 {
		return model.AuditEvent{}
	}
	return es[len(es)-1]
}
