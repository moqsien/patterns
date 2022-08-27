package facade

import "fmt"

type Account struct {
	name string
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

func newAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}
