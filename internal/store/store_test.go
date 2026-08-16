package store

import (
	"testing"

	"welfaremall/internal/model"
)

func TestFlashStockDeduct(t *testing.T) {
	s := New()
	_ = s.SetFlashItem(model.FlashItem{SKU: "sku", Stock: 10})
	if err := s.DeductFlashStock("sku", 4); err != nil {
		t.Fatal(err)
	}
	it, _ := s.GetFlashItem("sku")
	if it.Stock != 6 {
		t.Fatalf("stock=%d want 6", it.Stock)
	}
	if err := s.DeductFlashStock("sku", 99); err == nil {
		t.Fatal("expected oversell error")
	}
}

func TestExpireLedgerDedup(t *testing.T) {
	s := New()
	s.AppendLedger("a", 1)
	s.AppendLedger("b", -2)
	if got := s.ExpireLedgerEntries(); got != 2 {
		t.Fatalf("expired=%d want 2", got)
	}
	if got := s.ExpireLedgerEntries(); got != 0 {
		t.Fatalf("second expire=%d want 0 (no double processing)", got)
	}
	total, expired := s.LedgerCounts()
	if total != 2 || expired != 2 {
		t.Fatalf("total=%d expired=%d", total, expired)
	}
}
