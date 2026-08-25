package report

import (
	"example.com/receiptcheck/internal/model"
	"fmt"
)

func Line(r model.Receipt) string { return fmt.Sprintf("%s | %s | %.2f", r.ID, r.Status, r.Amount) }
func Lines(rs []model.Receipt) []string {
	out := make([]string, 0, len(rs))
	for _, r := range rs {
		out = append(out, Line(r))
	}
	return out
}
