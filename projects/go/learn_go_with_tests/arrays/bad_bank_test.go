package arrays

import "testing"

func TestBadBank(t *testing.T) {
	t.Run("calculate the balance for a person", func(t *testing.T) {
		transactions := []Transaction{
			{
				From: "Chris",
				To:   "Riya",
				Sum:  100,
			},
			{
				From: "Adil",
				To:   "Chris",
				Sum:  25,
			},
		}

		AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
		AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
		AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
	})

	t.Run("calculate the balance for a person using NewBalanceFor", func(t *testing.T) {
		var (
			riya  = Account{Name: "Riya", Balance: 100}
			chris = Account{Name: "Chris", Balance: 75}
			adil  = Account{Name: "Adil", Balance: 200}

			transactions = []Transaction{
				NewTransaction(chris, riya, 100),
				NewTransaction(adil, chris, 25),
			}
		)

		newBalanceFor := func(account Account) float64 {
			return NewBalanceFor(account, transactions).Balance
		}

		AssertEqual(t, newBalanceFor(riya), 200)
		AssertEqual(t, newBalanceFor(chris), 0)
		AssertEqual(t, newBalanceFor(adil), 175)
	})
}
