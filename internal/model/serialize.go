package model

import (
	"encoding/json"
	"fmt"
)

func EncodeReceipt(r Receipt) ([]byte, error) { return json.Marshal(r) }
func DecodeReceipt(b []byte) (Receipt, error) { var r Receipt; e := json.Unmarshal(b, &r); return r, e }
func ReceiptKey(r Receipt) string {
	if r.ID == "" {
		return fmt.Sprintf("%s:%s", r.FileName, r.Status)
	}
	return r.ID
}
func BatchRatio(b ValidationBatch) float64 {
	if b.Total == 0 {
		return 0
	}
	return float64(b.Passed) / float64(b.Total)
}
