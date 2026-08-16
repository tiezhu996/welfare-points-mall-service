package worker

import (
	"testing"

	"welfaremall/internal/store"
)

func TestRunExpiresOnce(t *testing.T) {
	s := store.New()
	s.AppendLedger("a", 1)
	w := New(s)
	if got := w.Run(); got != 1 {
		t.Fatalf("first run=%d want 1", got)
	}
	if got := w.Run(); got != 0 {
		t.Fatalf("second run=%d want 0", got)
	}
}
