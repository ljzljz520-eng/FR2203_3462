package store

import (
	"database/sql"
	"encoding/json"
	"example.com/receiptcheck/internal/model"
	"fmt"
	_ "modernc.org/sqlite"
)

type Store struct {
	db   *sql.DB
	path string
}

func Open(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db, path: path}
	if err := s.init(); err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) init() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS receipts(id TEXT PRIMARY KEY,data BLOB);CREATE TABLE IF NOT EXISTS batches(id TEXT PRIMARY KEY,data BLOB);CREATE TABLE IF NOT EXISTS rules(id TEXT PRIMARY KEY,data BLOB);CREATE TABLE IF NOT EXISTS audits(id TEXT PRIMARY KEY,data BLOB);`)
	return err
}
func (s *Store) Close() error { return s.db.Close() }
func (s *Store) put(table, id string, v any) error {
	b, e := json.Marshal(v)
	if e != nil {
		return e
	}
	_, e = s.db.Exec("INSERT OR REPLACE INTO "+table+"(id,data) VALUES(?,?)", id, b)
	return e
}
func (s *Store) PutReceipt(v model.Receipt) error       { return s.put("receipts", v.ID, v) }
func (s *Store) PutBatch(v model.ValidationBatch) error { return s.put("batches", v.ID, v) }
func (s *Store) PutRule(v model.ValidationRule) error   { return s.put("rules", v.ID, v) }
func (s *Store) PutAudit(v model.AuditEvent) error      { return s.put("audits", v.ID, v) }
func (s *Store) get(table, id string, out any) error {
	var b []byte
	if e := s.db.QueryRow("SELECT data FROM "+table+" WHERE id=?", id).Scan(&b); e != nil {
		return e
	}
	return json.Unmarshal(b, out)
}
func (s *Store) GetReceipt(id string) (model.Receipt, error) {
	var v model.Receipt
	e := s.get("receipts", id, &v)
	return v, e
}
func (s *Store) GetBatch(id string) (model.ValidationBatch, error) {
	var v model.ValidationBatch
	e := s.get("batches", id, &v)
	return v, e
}
func (s *Store) GetRule(id string) (model.ValidationRule, error) {
	var v model.ValidationRule
	e := s.get("rules", id, &v)
	return v, e
}
func (s *Store) GetAudit(id string) (model.AuditEvent, error) {
	var v model.AuditEvent
	e := s.get("audits", id, &v)
	return v, e
}
func (s *Store) Count(table string) (int, error) {
	var n int
	e := s.db.QueryRow("SELECT COUNT(*) FROM " + table).Scan(&n)
	return n, e
}
func (s *Store) Path() string { return s.path }
func ValidateStore(s *Store) error {
	if s == nil {
		return fmt.Errorf("nil store")
	}
	return nil
}
