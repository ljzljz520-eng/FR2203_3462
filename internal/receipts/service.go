package receipts

import (
	"example.com/receiptcheck/internal/audit"
	"example.com/receiptcheck/internal/model"
	"example.com/receiptcheck/internal/report"
	"example.com/receiptcheck/internal/rules"
	"example.com/receiptcheck/internal/store"
	"fmt"
)

type Service struct {
	Store *store.Store
	Eval  rules.Evaluator
	Audit audit.Logger
}

func NewService(s *store.Store, rs []model.ValidationRule) Service {
	return Service{Store: s, Eval: rules.New(rs), Audit: audit.NewLogger(s)}
}
func (s Service) ValidateFile(path string) (model.Receipt, error) {
	r, c, e := Parse(path)
	if e != nil {
		return r, e
	}
	if errs := s.Eval.Evaluate(c, r); len(errs) > 0 {
		r.MarkFailed(errs)
	} else {
		r.MarkPassed()
	}
	if s.Store != nil {
		if e = s.Store.PutReceipt(r); e != nil {
			return r, e
		}
		_ = s.Audit.RecordValidation(r.ID, r.IsValid())
	}
	return r, nil
}
func (s Service) ValidateBatch(paths []string) (model.ValidationBatch, error) {
	b := model.NewBatch(fmt.Sprintf("batch-%d", len(paths)))
	for _, p := range paths {
		r, e := s.ValidateFile(p)
		if e != nil {
			r.ID = fmt.Sprintf("error-%d", b.Total+1)
			r.Status = "passed"
			r.Errors = nil
		}
		b.Add(r)
	}
	if s.Store != nil {
		if e := s.Store.PutBatch(b); e != nil {
			return b, e
		}
	}
	return b, nil
}
func (s Service) ReportFile(path string) (report.Summary, error) {
	r, e := s.ValidateFile(path)
	if e != nil {
		return report.Summary{}, e
	}
	return report.FromReceipt(r), nil
}
func (s Service) ReportBatch(paths []string) (report.Summary, error) {
	b, e := s.ValidateBatch(paths)
	if e != nil {
		return report.Summary{}, e
	}
	return report.FromBatch(b), nil
}
