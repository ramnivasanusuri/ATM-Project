package ExistingUser

import (
	"ATM/CreateAccount"
	"fmt"
)

func ChangePin(c *CreateAccount.Customer, i int) {

	var np int
	fmt.Printf("Enter new pin:")
	fmt.Scan(&np)

	c.Acc_pin = np

}
