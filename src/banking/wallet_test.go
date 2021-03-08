package main
import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("Deposit:", func (t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw:", func (t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(0))
	})

	t.Run("Withdraw with insufficient funds:", func(t *testing.T) {
		startBalance := Bitcoin(10)
		wallet := Wallet{startBalance}
		err := wallet.Withdraw(Bitcoin(100))

		//third parameter below is startBalance since you don't want
		//balance to be reduced from your account if insufficient funds
		assertBalance(t, wallet, startBalance)
		if err == nil {
			t.Error("Wanted an error, didn't get one.")
		}
	})
}
