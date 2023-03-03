package ExistingUser

import "fmt"

func ExisUser() {
	fmt.Printf("\nChoose one:\n1.Account Details\n2.Balance Enquiry\n3.Cash Withdrawal\n4.Cash Deposit\n5.Change Pin\n6.Exit\n")
	var n int
	fmt.Scan(&n)
	switch n {
	case 1:
		AccDetails()
	case 2:
		BalanceEnquiry()
	case 3:
		CashWithdrawl()
	case 4:
		CashDeposit()
	case 5:
		ChangePin()
	case 6:

	}
}
