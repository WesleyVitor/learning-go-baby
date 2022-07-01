package pointer

import "testing"

func TestWallet(t *testing.T){
	
	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin){
		t.Helper()
		got:= wallet.Balance()
		if got != expected{
			t.Errorf("got %s expect %s", got, expected)
		}
	}

	assertError := func(t testing.TB, got error, want string){
		t.Helper()

		if got == nil{
			t.Fatal("Don´t get any error, but is need")
		}

		if got.Error() != want{
			t.Errorf("got %s want %s", got, want)
		}

	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		
		expected := Bitcoin(10.0)

		assertBalance(t, wallet, expected)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20.0)}

		wallet.Withdraw(Bitcoin(10.0))
		expected := Bitcoin(10.0)

		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw insufficient founds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, startingBalance)
		
	})
}