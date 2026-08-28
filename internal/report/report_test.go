package report

import (
	"example.com/receiptcheck/internal/model"
	"testing"
)

func TestReportFormatting(t *testing.T) {
	if Text(FromReceipt(model.NewReceipt("x", "x", 0))) == "" {
		t.Fatal("empty")
	}
}
