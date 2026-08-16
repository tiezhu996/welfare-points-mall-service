package service

import (
	"fmt"

	"welfaremall/internal/model"
	"welfaremall/internal/store"
)

type Service struct{ store *store.Store }

func New(s *store.Store) *Service { return &Service{store: s} }

func (s *Service) Spend(accountID string, amount int) error {
	acc, _ := s.store.GetAccount(accountID)
	acc.Points += amount
	_ = s.store.UpsertAccount(acc)
	return nil
}

func (s *Service) FlashPurchase(sku string, qty int) error {
	it, ok := s.store.GetFlashItem(sku)
	if !ok {
		return fmt.Errorf("item not found")
	}
	if !model.CanPurchase(it.Stock, qty) {
		return fmt.Errorf("flash stock insufficient")
	}
	return s.store.DeductFlashStock(sku, qty)
}

func (s *Service) Balance(accountID string) int {
	acc, ok := s.store.GetAccount(accountID)
	if !ok {
		return 0
	}
	return acc.Points
}
