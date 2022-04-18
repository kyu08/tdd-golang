package test

import (
	"hello/cmd/pointers"
	"testing"
)

// helper
func assertBalance(t *testing.T, wallet pointers.Wallet, want pointers.Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("👺 got %s, but want %s", got, want)
	}
}

func assertErrorMessage(t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("👺 wanted an error but didn't get one")
	}

	if got.Error() != want {
		t.Errorf("👺 got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("👺 wanted nil but got error")

	}
}

func TestWallet(t *testing.T) {
	// tests
	t.Run("deposit", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, pointers.Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := pointers.Test_new(20)
		err := wallet.Withdraw(10)
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := pointers.Test_new(20)
		err := wallet.Withdraw(pointers.Bitcoin(30))
		assertErrorMessage(t, err, pointers.ErrorMessage)
	})
}
