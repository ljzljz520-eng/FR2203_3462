package receipts

import (
	"crypto/sha256"
	"example.com/receiptcheck/internal/model"
	"example.com/receiptcheck/internal/rules"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(path string) (model.Receipt, string, error) {
	b, e := os.ReadFile(path)
	if e != nil {
		return model.Receipt{}, "", e
	}
	id := fmt.Sprintf("%x", sha256.Sum256(b))[:16]
	r := model.NewReceipt(id, filepath.Base(path), 0)
	content := rules.Normalize(string(b))
	amount, e := rules.ParseAmount(content)
	if e == nil {
		r.Amount = amount
	}
	r.Currency = rules.ParseCurrency(content)
	return r, content, e
}
func Parse(path string) (model.Receipt, string, error) {
	r, c, e := ReadFile(path)
	if e != nil {
		return r, c, e
	}
	if !rules.HasHeader(c) {
		return r, c, fmt.Errorf("invalid receipt header")
	}
	return r, c, nil
}
func Digest(content string) string { return fmt.Sprintf("%x", sha256.Sum256([]byte(content))) }
