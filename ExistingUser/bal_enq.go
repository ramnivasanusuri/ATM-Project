package ExistingUser

import (
	"ATM/CreateAccount"
	"fmt"
)

func BalanceEnquiry(c CreateAccount.Customer) {
	fmt.Printf("Available balance:%v", c.Acc_bal)
}
