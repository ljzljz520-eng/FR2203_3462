package receipts

import (
	"example.com/receiptcheck/internal/model"
	"example.com/receiptcheck/internal/report"
)

func WorkflowOne(s Service, path string) (report.Summary, error)    { return s.ReportFile(path) }
func WorkflowTwo(s Service, paths []string) (report.Summary, error) { return s.ReportBatch(paths) }
func WorkflowThree(s Service, id string) (model.Receipt, error)     { return s.Store.GetReceipt(id) }
