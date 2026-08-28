package store

import (
	"database/sql"
	"encoding/json"
	"example.com/receiptcheck/internal/model"
)

func (s *Store) ListReceipts() ([]model.Receipt, error) {
	rows, e := s.db.Query("SELECT data FROM receipts ORDER BY id")
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	var out []model.Receipt
	for rows.Next() {
		var b []byte
		if e = rows.Scan(&b); e != nil {
			return nil, e
		}
		var v model.Receipt
		if e = json.Unmarshal(b, &v); e != nil {
			return nil, e
		}
		out = append(out, v)
	}
	return out, rows.Err()
}
func (s *Store) ListAudits() ([]model.AuditEvent, error) {
	rows, e := s.db.Query("SELECT data FROM audits ORDER BY id")
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	var out []model.AuditEvent
	for rows.Next() {
		var b []byte
		if e = rows.Scan(&b); e != nil {
			return nil, e
		}
		var v model.AuditEvent
		if e = json.Unmarshal(b, &v); e != nil {
			return nil, e
		}
		out = append(out, v)
	}
	return out, rows.Err()
}
func scanOne(row *sql.Row, out any) error {
	var b []byte
	if e := row.Scan(&b); e != nil {
		return e
	}
	return json.Unmarshal(b, out)
}
