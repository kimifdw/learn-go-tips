package points

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) withDraw(amount Bitcoin) error {

	if amount > w.balance {
		return InsufficientFundsError
	}

	w.balance -= amount
	return nil
}

// String：自定义输出格式
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}