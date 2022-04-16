package test

import (
	"hello/cmd/pointers"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet pointers.Wallet, want pointers.Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("👺 got %s, but want %s", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, pointers.Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := pointers.Test_new(20)
		wallet.Withdraw(10)
		assertBalance(t, wallet, pointers.Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := pointers.Test_new(20)
		err := wallet.Withdraw(pointers.Bitcoin(30))
		assertBalance(t, wallet, pointers.Bitcoin(10))

		if err == nil {
			t.Error("👺 wanted an error but didn't get one.")
		}
	})
}
