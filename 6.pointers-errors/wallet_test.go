package pointers_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assert := func(t *testing.T, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	t.Run(`wallet balance test`, func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		assert(t, got, want)
	})
}
