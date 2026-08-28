package store

import (
	"context"
	"time"
)

func (s *Store) Ping() error {
	ctx, c := context.WithTimeout(context.Background(), time.Second)
	defer c()
	return s.db.PingContext(ctx)
}
func (s *Store) Vacuum() error { _, e := s.db.Exec("VACUUM"); return e }
func (s *Store) Ready() bool   { return s != nil && s.Ping() == nil }
