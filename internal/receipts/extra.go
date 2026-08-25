package receipts

import (
	"example.com/receiptcheck/internal/model"
	"fmt"
	"strings"
)

func ValidateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("name required")
	}
	if strings.ContainsAny(name, "/\\") {
		return fmt.Errorf("name must be base")
	}
	return nil
}
func CloneReceipt(r model.Receipt) model.Receipt {
	r.Errors = append([]string(nil), r.Errors...)
	return r
}
func StatusLabel(r model.Receipt) string {
	if r.IsValid() {
		return "PASS"
	}
	return "FAIL"
}
func ErrorCount(r model.Receipt) int { return len(r.Errors) }
