package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, e error, want string) {
		t.Helper()

		if e == nil {
			t.Fatal("!!!")
		}

		if e.Error() != want {
			t.Errorf("got %s, want %s", e.Error(), want)
		}
	}

	t.Run("deposit", func(t *testing.T){
		w := Wallet{}
		w.Deposit(1164)
		want := Bitcoin(1164)
		assertBalance(t, w, want)
	})

	t.Run("withdraw", func(t *testing.T){
		w := Wallet{Bitcoin(10000)}
		w.Withdraw(1164)
		want := Bitcoin(8836)
		assertBalance(t, w, want)
	})

	t.Run("withdraw more than balance", func(t *testing.T){
		w := Wallet{Bitcoin(1000)}
		err := w.Withdraw(1164)
		want := Bitcoin(1000)
		assertBalance(t, w, want)
		assertError(t, err, "oh oh")
	})
}