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

func updateJson(c []CreateAccount.Customer){
	fl_pt,err:=os.OpenFile("CreateAccount/data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	check_error(err)
	defer fl_pt.Close()
	
	err=json.NewEncoder(fl_pt).Encode(c)
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

	var index int
	for i, v := range c {
		if v.Acc_num == il {
			index = i
			break
		}
	}

	cus = c[index]

	var epin int
	fmt.Printf("Enter your pin:")
	fmt.Scan(&epin)
checkpin:
	if epin != cus.Acc_pin{
		fmt.Println("Incorrect pin. Try again")
		goto checkpin
	}

	fmt.Printf("\nChoose one:\n1.Account Details\n2.Balance Enquiry\n3.Cash Withdrawal\n4.Cash Deposit\n5.Change Pin\n6.Exit\n")
	var n int
	fmt.Scan(&n)
	switch n {
	case 1:
		AccDetails(cus)
	case 2:
		BalanceEnquiry(cus)
	case 3:
		CashWithdrawl(&cus, index)
		c[index].Acc_bal=cus.Acc_bal
		updateJson(c)
		fmt.Println("\nCash has been withdrawn.")
	case 4:
		CashDeposit(&cus, index)
		c[index].Acc_bal=cus.Acc_bal
		updateJson(c)
		fmt.Println("Cash has been deposited.")
	case 5:
		ChangePin(&cus, index)
		c[index].Acc_pin=cus.Acc_pin
		updateJson(c)
	case 6:
	
	default:
		fmt.Println("Enter valid number.")

	}
}