package pointers_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assert := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %#v, want %#v", got, want)
		}
	}
	t.Run(`wallet balance test`, func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		got := wallet.Balance()
		want := 10.0
		assert(t, got, want)
	})
}
