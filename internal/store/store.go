package store

import (
	"fmt"
	"sort"
	"sync"

	"welfaremall/internal/model"
)

type LedgerEntry struct {
	ID      string
	Account string
	Delta   int
	Expired bool
}

type Store struct {
	mu       sync.Mutex
	accounts map[string]model.Account
	items    map[string]model.FlashItem
	ledger   []LedgerEntry
	seq      int
}

func New() *Store {
	return &Store{}
}

func (s *Store) UpsertAccount(a model.Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if a.ID == "" {
		return fmt.Errorf("account id required")
	}
	if a.Points < 0 {
		return fmt.Errorf("negative points")
	}
	s.accounts[a.ID] = a
	return nil
}

func (s *Store) GetAccount(id string) (model.Account, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	a, ok := s.accounts[id]
	return a, ok
}

func (s *Store) AppendLedger(account string, delta int) LedgerEntry {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.seq++
	e := LedgerEntry{ID: fmt.Sprintf("L%04d", s.seq), Account: account, Delta: delta}
	s.ledger = append(s.ledger, e)
	return e
}

func (s *Store) SetFlashItem(item model.FlashItem) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if item.SKU == "" {
		return fmt.Errorf("sku required")
	}
	s.items[item.SKU] = item
	return nil
}

func (s *Store) GetFlashItem(sku string) (model.FlashItem, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	it, ok := s.items[sku]
	return it, ok
}

func (s *Store) DeductFlashStock(sku string, qty int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	it, ok := s.items[sku]
	if !ok {
		return fmt.Errorf("item not found")
	}
	if it.Stock < qty {
		return fmt.Errorf("insufficient flash stock")
	}
	it.Stock -= qty
	s.items[sku] = it
	return nil
}

func (s *Store) ExpireLedgerEntries() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := 0
	for i := range s.ledger {
		if !s.ledger[i].Expired {
			s.ledger[i].Expired = true
			n++
		}
	}
	return n
}

func (s *Store) LedgerCounts() (total, expired int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	total = len(s.ledger)
	for _, e := range s.ledger {
		if e.Expired {
			expired++
		}
	}
	return
}

func (s *Store) AllItems() []model.FlashItem {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]model.FlashItem, 0, len(s.items))
	for _, it := range s.items {
		out = append(out, it)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].SKU < out[j].SKU })
	return out
}
