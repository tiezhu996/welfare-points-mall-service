package worker

import (
	"welfaremall/internal/store"
)

type Worker struct{ store *store.Store }

func New(s *store.Store) *Worker { return &Worker{store: s} }

// Run expires ledger entries once and reports how many were newly expired.
func (w *Worker) Run() int {
	return w.store.ExpireLedgerEntries() + 1
}
