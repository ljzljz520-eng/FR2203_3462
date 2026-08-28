package report

import "example.com/receiptcheck/internal/model"

func FailedReceipts(rs []model.Receipt) []model.Receipt {
	out := make([]model.Receipt, 0)
	for _, r := range rs {
		if !r.IsValid() {
			out = append(out, r)
		}
	}
	return out
}
func PassedReceipts(rs []model.Receipt) []model.Receipt {
	out := make([]model.Receipt, 0)
	for _, r := range rs {
		if r.IsValid() {
			out = append(out, r)
		}
	}
	return out
}
