package points

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, err error, want error) {
		if err == nil {
			t.Error("wanted an error but didn't got one")
		}

		if err != want {
			t.Errorf("got '%s', want '%s'", err, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withDraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}

		wallet.withDraw(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {

		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.withDraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		assertError(t, err, InsufficientFundsError)

	})
}
