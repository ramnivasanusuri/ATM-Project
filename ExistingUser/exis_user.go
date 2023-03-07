package ExistingUser

import (
	"ATM/CreateAccount"
	"encoding/json"
	"fmt"
	"os"
)

func check_error(e error) {
	if e != nil {
		fmt.Println("Error occured:", e)
	}
}

func updateJson(c []CreateAccount.Customer) {
	fl_pt, err := os.OpenFile("CreateAccount/data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	check_error(err)
	defer fl_pt.Close()

	err = json.NewEncoder(fl_pt).Encode(c)
	check_error(err)
}

func ExisUser() {
	var c []CreateAccount.Customer
	var cus CreateAccount.Customer

	fmt.Printf("Enter account number:")
	var il int
	fmt.Scan(&il)

	file, err := os.ReadFile("CreateAccount/data.json")
	check_error(err)

	err = json.Unmarshal(file, &c)
	check_error(err)

	index := -1
	for i, v := range c {
		if v.Acc_num == il {
			index = i
			break
		}
	}
	if index != -1 {
		cus = c[index]
	}else{
		fmt.Println("\nAccount not found! Try creating an account.")
		return
	}

	var epin int
checkpin:
	fmt.Printf("Enter your pin:")
	fmt.Scan(&epin)
	if epin != cus.Acc_pin {
		fmt.Println("Incorrect pin. Try again")
		goto checkpin
	}
loop:
	fmt.Printf("\nChoose one:\n1.Account Details\n2.Balance Enquiry\n3.Cash Withdrawal\n4.Cash Deposit\n5.Change Pin\n6.Exit\n")
	var n int
	fmt.Scan(&n)
	switch n {
	case 1:
		AccDetails(cus)
		goto loop
	case 2:
		BalanceEnquiry(cus)
		goto loop
	case 3:
		CashWithdrawl(&cus, index)
		c[index].Acc_bal = cus.Acc_bal
		updateJson(c)
		fmt.Println("\nCash has been withdrawn.")
		goto loop
	case 4:
		CashDeposit(&cus, index)
		c[index].Acc_bal = cus.Acc_bal
		updateJson(c)
		fmt.Println("Cash has been deposited.")
		goto loop
	case 5:
		ChangePin(&cus, index)
		c[index].Acc_pin = cus.Acc_pin
		updateJson(c)
		goto loop
	case 6:
		os.Exit(0)

	default:
		fmt.Println("Enter valid number.")

	}
}
