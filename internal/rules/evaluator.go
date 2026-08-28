package rules

import (
	"example.com/receiptcheck/internal/model"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Evaluator struct{ Rules []model.ValidationRule }

func New(rs []model.ValidationRule) Evaluator { return Evaluator{Rules: rs} }
func (e Evaluator) Evaluate(content string, r model.Receipt) []string {
	var out []string
	for _, rule := range e.Rules {
		if !rule.Enabled {
			continue
		}
		if rule.Name == "required" && strings.TrimSpace(content) == "" {
			out = append(out, rule.Pattern)
		}
		if rule.Name == "pattern" {
			ok, _ := regexp.MatchString(rule.Pattern, content)
			if !ok {
				out = append(out, rule.Pattern)
			}
		}
		if rule.Name == "amount" && r.Amount <= 0 {
			out = append(out, "amount must be positive")
		}
	}
	return out
}
func ParseAmount(content string) (float64, error) {
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(line, "AMOUNT=") {
			return strconv.ParseFloat(strings.TrimPrefix(line, "AMOUNT="), 64)
		}
	}
	return 0, fmt.Errorf("amount missing")
}
func ParseCurrency(content string) string {
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(line, "CURRENCY=") {
			return strings.TrimPrefix(line, "CURRENCY=")
		}
	}
	return "CNY"
}
func Normalize(content string) string {
	return strings.TrimSpace(strings.ReplaceAll(content, "\r\n", "\n"))
}
func HasHeader(content string) bool { return strings.HasPrefix(Normalize(content), "RECEIPT") }
func SeverityWeight(s string) int {
	switch strings.ToLower(s) {
	case "error":
		return 3
	case "warning":
		return 1
	default:
		return 0
	}
}
