package receipts

import (
	"example.com/receiptcheck/internal/model"
	"example.com/receiptcheck/internal/report"
	"example.com/receiptcheck/internal/rules"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestWorkflowTwo_RejectsBadFiles guards the batch summary: a batch that
// contains unreadable / malformed receipt files must report those files as
// failed with their errors surfaced, instead of marking the whole batch as
// passed.
func TestWorkflowTwo_RejectsBadFiles(t *testing.T) {
	dir := t.TempDir()

	good := filepath.Join(dir, "good.txt")
	mustWrite(t, good, "RECEIPT\nAMOUNT=10.00\nCURRENCY=CNY\n")

	badHeader := filepath.Join(dir, "bad_header.txt")
	mustWrite(t, badHeader, "NOTARECEIPT\nAMOUNT=5.00\nCURRENCY=CNY\n")

	missing := filepath.Join(dir, "missing.txt") // never created

	svc := NewService(nil, rules.DefaultRules())
	paths := []string{good, badHeader, missing}

	sum, err := svc.ReportBatch(paths)
	if err != nil {
		t.Fatalf("ReportBatch error: %v", err)
	}

	if sum.Total != 3 {
		t.Fatalf("total = %d, want 3", sum.Total)
	}
	if sum.Passed != 1 {
		t.Fatalf("passed = %d, want 1", sum.Passed)
	}
	if sum.Failed != 2 {
		t.Fatalf("failed = %d, want 2", sum.Failed)
	}

	// The two failed files must be named in the error summary.
	txt := report.Text(sum)
	if !strings.Contains(txt, "bad_header.txt") {
		t.Errorf("summary missing bad_header.txt:\n%s", txt)
	}
	if !strings.Contains(txt, "missing.txt") {
		t.Errorf("summary missing missing.txt:\n%s", txt)
	}
	// Passed files must not be lumped into the failures section.
	if strings.Contains(strings.SplitN(txt, "failures:", 2)[1], "good.txt") {
		t.Errorf("passed file leaked into failures:\n%s", txt)
	}
	// Bad files must never be reported as passed.
	if !strings.Contains(txt, "failed=2") {
		t.Errorf("summary missing failed=2:\n%s", txt)
	}
}

// TestWorkflowTwo_AllPassed confirms a clean batch still reports success.
func TestWorkflowTwo_AllPassed(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "good.txt")
	mustWrite(t, good, "RECEIPT\nAMOUNT=10.00\nCURRENCY=CNY\n")

	svc := NewService(nil, rules.DefaultRules())
	sum, err := svc.ReportBatch([]string{good})
	if err != nil {
		t.Fatalf("ReportBatch error: %v", err)
	}
	if !report.Success(sum) {
		t.Fatalf("expected success, got %+v", sum)
	}
}

func mustWrite(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

// guard against the model layer silently miscounting receipts.
func TestBatchAdd_CountsFailures(t *testing.T) {
	b := model.NewBatch("b1")
	b.Add(model.Receipt{ID: "a", Status: "passed"})
	b.Add(model.Receipt{ID: "b", Status: "failed", Errors: []string{"boom"}})
	b.Add(model.Receipt{ID: "c", Status: "failed", Errors: []string{"nope"}})
	if b.Total != 3 || b.Passed != 1 || b.Failed != 2 {
		t.Fatalf("counts = total=%d passed=%d failed=%d, want 3/1/2", b.Total, b.Passed, b.Failed)
	}
	if len(b.Receipts) != 3 {
		t.Fatalf("Receipts len = %d, want 3", len(b.Receipts))
	}
}
