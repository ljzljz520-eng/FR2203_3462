package report

import (
	"encoding/json"
	"example.com/receiptcheck/internal/model"
	"fmt"
	"strings"
)

type Summary struct {
	Total, Passed, Failed int
	Errors                []string
}

func FromReceipt(r model.Receipt) Summary {
	s := Summary{Total: 1}
	if r.IsValid() {
		s.Passed = 1
	} else {
		s.Failed = 1
		s.Errors = append(s.Errors, r.Errors...)
	}
	return s
}
func FromBatch(b model.ValidationBatch) Summary {
	return Summary{Total: b.Total, Passed: b.Passed, Failed: b.Failed}
}
func Merge(a, b Summary) Summary {
	a.Total += b.Total
	a.Passed += b.Passed
	a.Failed += b.Failed
	a.Errors = append(a.Errors, b.Errors...)
	return a
}
func Text(s Summary) string {
	return fmt.Sprintf("total=%d passed=%d failed=%d errors=%s", s.Total, s.Passed, s.Failed, strings.Join(s.Errors, "|"))
}
func JSON(s Summary) string  { b, _ := json.Marshal(s); return string(b) }
func Success(s Summary) bool { return s.Failed == 0 }
