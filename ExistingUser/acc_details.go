package ExistingUser

import ("fmt"
"ATM/CreateAccount"
)

func AccDetails(c CreateAccount.Customer) {
	fmt.Printf("Account holder name:%s\nAccount number:%v\nAge:%v",c.Cus_name,c.Acc_num,c.Cus_age)
}