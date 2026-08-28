package receipts

import (
	"example.com/receiptcheck/internal/model"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type CheckResult struct {
	Name, Code string
	Passed     bool
	Detail     string
}
type Document struct {
	Lines    []string
	Fields   map[string]string
	ParsedAt time.Time
}

func ParseDocument(content string) Document {
	d := Document{Lines: strings.Split(content, "\n"), Fields: map[string]string{}, ParsedAt: time.Now().UTC()}
	for _, line := range d.Lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			d.Fields[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return d
}
func CheckHeader(d Document) CheckResult {
	if len(d.Lines) == 0 {
		return CheckResult{"header", "E001", false, "empty document"}
	}
	ok := strings.TrimSpace(d.Lines[0]) == "RECEIPT"
	return CheckResult{"header", "E001", ok, "receipt header"}
}
func CheckAmount(d Document) CheckResult {
	v, ok := d.Fields["AMOUNT"]
	if !ok {
		return CheckResult{"amount", "E002", false, "missing amount"}
	}
	n, e := strconv.ParseFloat(v, 64)
	if e != nil || n <= 0 {
		return CheckResult{"amount", "E002", false, "amount must be positive"}
	}
	return CheckResult{"amount", "E002", true, "amount accepted"}
}
func CheckCurrency(d Document) CheckResult {
	v := d.Fields["CURRENCY"]
	ok := v == "CNY" || v == "USD" || v == "EUR"
	return CheckResult{"currency", "E003", ok, "supported currency"}
}
func CheckDate(d Document) CheckResult {
	v := d.Fields["DATE"]
	if v == "" {
		return CheckResult{"date", "E004", false, "date missing"}
	}
	_, e := time.Parse("2006-01-02", v)
	return CheckResult{"date", "E004", e == nil, "iso date"}
}
func CheckID(d Document) CheckResult {
	v := d.Fields["ID"]
	ok, _ := regexp.MatchString(`^[A-Z0-9-]{4,40}$`, v)
	return CheckResult{"id", "E005", ok, "identifier format"}
}
func CheckTax(d Document) CheckResult {
	v := d.Fields["TAX"]
	if v == "" {
		return CheckResult{"tax", "E006", true, "tax optional"}
	}
	n, e := strconv.ParseFloat(v, 64)
	return CheckResult{"tax", "E006", e == nil && n >= 0, "tax nonnegative"}
}
func CheckTotal(d Document) CheckResult {
	a, _ := strconv.ParseFloat(d.Fields["AMOUNT"], 64)
	t, _ := strconv.ParseFloat(d.Fields["TOTAL"], 64)
	if d.Fields["TOTAL"] == "" {
		return CheckResult{"total", "E007", true, "total optional"}
	}
	return CheckResult{"total", "E007", t >= a, "total covers amount"}
}
func CheckVendor(d Document) CheckResult {
	return CheckResult{"vendor", "E008", strings.TrimSpace(d.Fields["VENDOR"]) != "", "vendor present"}
}
func CheckBuyer(d Document) CheckResult {
	return CheckResult{"buyer", "E009", strings.TrimSpace(d.Fields["BUYER"]) != "", "buyer present"}
}
func CheckNumber(d Document) CheckResult {
	v := d.Fields["NUMBER"]
	ok := len(v) >= 6
	return CheckResult{"number", "E010", ok, "number length"}
}
func CheckChecksum(d Document) CheckResult {
	v := d.Fields["CHECKSUM"]
	ok := len(v) == 64
	return CheckResult{"checksum", "E011", ok, "sha256 checksum"}
}
func CheckVersion(d Document) CheckResult {
	v := d.Fields["VERSION"]
	ok := v == "1" || v == "2"
	return CheckResult{"version", "E012", ok, "supported version"}
}
func CheckEncoding(content string) CheckResult {
	ok := !strings.ContainsRune(content, '\x00')
	return CheckResult{"encoding", "E013", ok, "utf8 content"}
}
func CheckLength(content string) CheckResult {
	ok := len(content) <= 1<<20
	return CheckResult{"length", "E014", ok, "document size"}
}
func CheckLineCount(d Document) CheckResult {
	ok := len(d.Lines) >= 1 && len(d.Lines) <= 1000
	return CheckResult{"lines", "E015", ok, "line count"}
}
func CheckFields(d Document) []CheckResult {
	return []CheckResult{CheckHeader(d), CheckAmount(d), CheckCurrency(d), CheckDate(d), CheckID(d), CheckTax(d), CheckTotal(d), CheckVendor(d), CheckBuyer(d), CheckNumber(d), CheckChecksum(d), CheckVersion(d), CheckLineCount(d)}
}
func FailedChecks(rs []CheckResult) []CheckResult {
	out := make([]CheckResult, 0)
	for _, r := range rs {
		if !r.Passed {
			out = append(out, r)
		}
	}
	return out
}
func CheckSummary(rs []CheckResult) string {
	n := 0
	for _, r := range rs {
		if !r.Passed {
			n++
		}
	}
	return fmt.Sprintf("%d checks failed", n)
}
func ApplyChecks(content string, r *model.Receipt) []CheckResult {
	d := ParseDocument(content)
	rs := append(CheckFields(d), CheckEncoding(content), CheckLength(content))
	if len(FailedChecks(rs)) > 0 {
		errs := make([]string, 0)
		for _, x := range FailedChecks(rs) {
			errs = append(errs, x.Code+":"+x.Detail)
		}
		r.MarkFailed(errs)
	} else {
		r.MarkPassed()
	}
	return rs
}
