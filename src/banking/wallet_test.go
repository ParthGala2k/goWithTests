package main
import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit:", func (t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw:", func (t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(10))
		if err != nil {
			t.Fatal("Didn't want an error, got one.")
		}
		assertBalance(t, wallet, Bitcoin(0))
	})

	t.Run("Withdraw with insufficient funds:", func(t *testing.T) {
		startBalance := Bitcoin(10)
		wallet := Wallet{startBalance}
		err := wallet.Withdraw(Bitcoin(100))

		//third parameter below is startBalance since you don't want
		//balance to be reduced from your account if insufficient funds
		assertBalance(t, wallet, startBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
       t.Helper()
       got := wallet.Balance()
       if got != want {
	       t.Errorf("got %s want %s", got, want)
       }
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
	     t.Fatal("Wanted an error, didn't get one.")
	}
	if got != want {
	     t.Errorf("got %q, want %q", got, want)
	}
}
