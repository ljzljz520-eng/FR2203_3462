package main

import (
	"example.com/receiptcheck/internal/config"
	"example.com/receiptcheck/internal/receipts"
	"example.com/receiptcheck/internal/report"
	"example.com/receiptcheck/internal/store"
	"flag"
	"fmt"
	"os"
)

func main() {
	db := flag.String("db", "receipts.db", "database path")
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("usage: receiptcheck files...")
		return
	}
	s, e := store.Open(*db)
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
	defer s.Close()
	svc := receipts.NewService(s, config.Default().Rules)
	sum, e := svc.ReportBatch(flag.Args())
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
	fmt.Println(report.Text(sum))
	if !report.Success(sum) {
		os.Exit(2)
	}
}
