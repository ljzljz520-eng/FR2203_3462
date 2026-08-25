package rules

import (
	"example.com/receiptcheck/internal/model"
	"fmt"
)

func DefaultRules() []model.ValidationRule {
	return []model.ValidationRule{model.NewRule("r1", "required", "content required", "error", true), model.NewRule("r2", "pattern", "^RECEIPT", "error", true), model.NewRule("r3", "amount", "positive amount", "error", true)}
}
func RuleByID(rs []model.ValidationRule, id string) (model.ValidationRule, bool) {
	for _, r := range rs {
		if r.ID == id {
			return r, true
		}
	}
	return model.ValidationRule{}, false
}
func Enabled(rs []model.ValidationRule) []model.ValidationRule {
	out := make([]model.ValidationRule, 0)
	for _, r := range rs {
		if r.Enabled {
			out = append(out, r)
		}
	}
	return out
}
func ValidateRule(r model.ValidationRule) error {
	if r.ID == "" || r.Name == "" {
		return fmt.Errorf("rule identity required")
	}
	return nil
}
