package ExistingUser

import (
	"ATM/CreateAccount"
	"fmt"
)

func CashDeposit(c *CreateAccount.Customer, i int) {

	var da int
	fmt.Printf("Enter amount to deposit:")
	fmt.Scan(&da)

	c.Acc_bal = c.Acc_bal + da

	
}
