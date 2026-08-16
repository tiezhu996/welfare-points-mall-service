package model

import "testing"

func TestCanSpend(t *testing.T) {
	cases := []struct {
		balance, amount int
		want            bool
	}{
		{100, 50, true},
		{50, 50, true},
		{49, 50, false},
		{100, 0, false},
	}
	for _, c := range cases {
		if got := CanSpend(c.balance, c.amount); got != c.want {
			t.Errorf("CanSpend(%d,%d)=%v want %v", c.balance, c.amount, got, c.want)
		}
	}
}

func TestCanPurchase(t *testing.T) {
	if CanPurchase(10, 5) != true {
		t.Fatal("expected purchase allowed")
	}
	if CanPurchase(5, 10) != false {
		t.Fatal("expected oversell rejected")
	}
}
