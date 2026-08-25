package rules

import (
	"example.com/receiptcheck/internal/model"
	"testing"
)

func TestRuleEvaluation(t *testing.T) {
	r := model.NewReceipt("1", "x", 2)
	if len(New(DefaultRules()).Evaluate("RECEIPT\nAMOUNT=2", r)) != 0 {
		t.Fatal("unexpected")
	}
}
