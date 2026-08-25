package model

import "time"

type ReceiptHistory struct {
	ReceiptID string
	Status    string
	Message   string
	At        time.Time
}
type BatchHistory struct {
	BatchID string
	Event   string
	Count   int
	At      time.Time
}

func NewReceiptHistory(id, status, msg string) ReceiptHistory {
	return ReceiptHistory{ReceiptID: id, Status: status, Message: msg, At: time.Now().UTC()}
}
func NewBatchHistory(id, event string, count int) BatchHistory {
	return BatchHistory{BatchID: id, Event: event, Count: count, At: time.Now().UTC()}
}
func LatestReceiptHistory(h []ReceiptHistory) ReceiptHistory {
	if len(h) == 0 {
		return ReceiptHistory{}
	}
	latest := h[0]
	for _, v := range h[1:] {
		if v.At.After(latest.At) {
			latest = v
		}
	}
	return latest
}
func LatestBatchHistory(h []BatchHistory) BatchHistory {
	if len(h) == 0 {
		return BatchHistory{}
	}
	latest := h[0]
	for _, v := range h[1:] {
		if v.At.After(latest.At) {
			latest = v
		}
	}
	return latest
}
func HistoryStatuses(h []ReceiptHistory) map[string]int {
	m := map[string]int{}
	for _, v := range h {
		m[v.Status]++
	}
	return m
}
func HistoryEvents(h []BatchHistory) map[string]int {
	m := map[string]int{}
	for _, v := range h {
		m[v.Event]++
	}
	return m
}
