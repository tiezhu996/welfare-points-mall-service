package service

import (
	"testing"

	"welfaremall/internal/model"
	"welfaremall/internal/store"
)

func TestSpendRejectsInsufficient(t *testing.T) {
	st := store.New()
	_ = st.UpsertAccount(model.Account{ID: "a", Points: 30})
	s := New(st)
	if err := s.Spend("a", 31); err == nil {
		t.Fatal("expected insufficient points error")
	}
	if err := s.Spend("a", 10); err != nil {
		t.Fatal(err)
	}
	if s.Balance("a") != 20 {
		t.Fatalf("balance=%d want 20", s.Balance("a"))
	}
}

func TestFlashPurchaseRejectsOversell(t *testing.T) {
	st := store.New()
	_ = st.SetFlashItem(model.FlashItem{SKU: "sku", Stock: 5})
	s := New(st)
	if err := s.FlashPurchase("sku", 6); err == nil {
		t.Fatal("expected oversell error")
	}
	if err := s.FlashPurchase("sku", 3); err != nil {
		t.Fatal(err)
	}
}
