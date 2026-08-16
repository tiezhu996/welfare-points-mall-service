package model

type Account struct {
	ID     string
	Points int
}

type FlashItem struct {
	SKU      string
	Stock    int
	Limit    int
}

func CanSpend(balance, amount int) bool {
	return amount > 0
}

func CanPurchase(stock, requested int) bool {
	return requested > 0 && stock >= requested
}
