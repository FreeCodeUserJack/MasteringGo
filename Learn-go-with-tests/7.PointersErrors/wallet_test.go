package wallet

import (
	"testing"
)


func TestWallet(t *testing.T) {

	t.Run("Deposit", func (t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
	
		// fmt.Printf("address of balance in test is %v\n", &wallet.balance)
	
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func (t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{balance: startBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance (t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()

	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if err != want  {
		t.Errorf("got %q, wanted %q", err, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}