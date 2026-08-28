package receipts

import (
	"example.com/receiptcheck/internal/model"
	"strings"
)

func CountByStatus(rs []model.Receipt) map[string]int {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Status]++
	}
	return m
}
func HasErrors(r model.Receipt) bool    { return len(r.Errors) > 0 }
func JoinErrors(r model.Receipt) string { return strings.Join(r.Errors, "; ") }
