package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	bitcoinCents Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.bitcoinCents
}

func (w *Wallet) Deposit(sum Bitcoin) {
	w.bitcoinCents += sum
}

var ErrCustom = errors.New("oh oh")

func (w *Wallet) Withdraw(sum Bitcoin) error {
	if sum > w.Balance() {
		return ErrCustom
	}
	
	w.bitcoinCents -= sum
	return nil
}