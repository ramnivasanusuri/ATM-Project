package ExistingUser

import (
	"ATM/CreateAccount"
	"fmt"
)

func CashWithdrawl(c *CreateAccount.Customer, i int) {

	var wa int
loo:
	fmt.Printf("Enter amount to withdraw:")
	fmt.Scan(&wa)

	if wa < c.Acc_bal {
		c.Acc_bal = c.Acc_bal - wa
		fmt.Printf("\nPlease collect your cash.")
	} else {
		fmt.Printf("\nInsufficient balance. Enter valid amount to withdraw.")
		goto loo
	}

}
