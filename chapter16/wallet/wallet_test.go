package wallet

import (
	"errors"
	"fmt"
	"testing"
)

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

// 类型别名
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is", &wallet.balance)
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return InsufficientFundsError
	}
	wallet.balance -= amount
	return nil
}

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, err error, want error) {

		if err == nil {
			// 停止测试
			t.Fatal("didn't get an error but wanted one")
		}

		if err != want {
			t.Errorf("got '%s' ,want '%s'", err, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		assertError(t, err, InsufficientFundsError)
	})
}
