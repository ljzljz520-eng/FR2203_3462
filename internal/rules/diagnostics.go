package rules

import (
	"example.com/receiptcheck/internal/model"
	"sort"
)

func SortRules(rs []model.ValidationRule) []model.ValidationRule {
	out := append([]model.ValidationRule(nil), rs...)
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}
func HighestSeverity(rs []model.ValidationRule) string {
	best := ""
	score := 0
	for _, r := range rs {
		if s := SeverityWeight(r.Severity); s > score {
			score = s
			best = r.Severity
		}
	}
	return best
}
