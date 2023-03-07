package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
	"ATM/CreateAccount"
	"ATM/ExistingUser"
)

// String to Integer conversion
// func strtoi(s string) int64 {
// 	a, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 0)
// 	return a
// }

func main() {
	// reader := bufio.NewReader(os.Stdin)
loop:
	fmt.Println("\nChoose one:\n1.Create an account.\n2.Existing user.")
	// n, _ := reader.ReadString('\n')
	// var i int = int(strtoi(n))
	var i int
	fmt.Scan(&i)
	switch i {
	case 1:
		CreateAccount.CreateAcc()
	goto loop
	case 2:
		ExistingUser.ExisUser()
		goto loop
	}
}

// func accountDetails(a [4]string, b [4]int, c int) {
// 	fmt.Printf("Account details :\n")
// 	fmt.Printf("Name :%s\n", a[c])
// 	fmt.Printf("Account number :%v\n", b[c])
// }

// func balanceEnquiry(a [4]int, b int) {
// 	fmt.Printf("Available balance :%v\n", a[b])
// }

// func cashWithdrawl(a *[4]int, b int) {
// 	fmt.Printf("\nEnter the amout you want to withdraw :")
// 	var cw int
// 	fmt.Scan(&cw)
// 	if cw > (*a)[b] {
// 		fmt.Println("Insufficient balance! Try again.")
// 		cashWithdrawl(a, b)
// 	} else {
// 		(*a)[b] -= cw
// 		fmt.Println("Cash has been withdrawn.")
// 		fmt.Printf("Available balance: %v", (*a)[b])
// 	}

// }

// func cashDeposit(a *[4]int, b int) {
// 	var cd int
// 	fmt.Println("Enter the amount you want to deposit:")
// 	fmt.Scan(&cd)
// 	(*a)[b] += cd
// 	fmt.Println("Cash has been deposited.")
// 	fmt.Printf("Available balance: %v", (*a)[b])
// }

// func changePin(a *[4]int, b int) {
// 	fmt.Println("Enter a 4 digit new PIN:")
// 	var np int
// 	fmt.Scan(&np)
// 	if len(strconv.FormatInt(int64(np), 10)) == 4 {
// 		(*a)[b] = np
// 		fmt.Println("Your PIN has been updated")
// 		fmt.Println(a)
// 	} else {
// 		fmt.Println("Please enter a four digit pin!")
// 		changePin(a, b)
// 	}
// }

// func main() {
// 	accName := [4]string{"ram", "vinay", "mani", "chandu"}
// 	accNumber := [4]int{1111, 2222, 3333, 4444}
// 	accPin := [4]int{1000, 2000, 3000, 4000}
// 	accBalance := [4]int{20000, 25000, 10000, 28000}
// 	var can string
// 	var cap int
// 	var ni int = -1
// 	fmt.Printf("Enter your name :")
// 	fmt.Scan(&can)

// 	for i, v := range accName {
// 		if can == v {
// 			ni = i
// 			break
// 			// } else {
// 			// 	fmt.Printf("Entered name is not registered in the database.")
// 			// 	os.Exit(0)
// 		}
// 	}
// 	if ni < 0 || ni >= 4 {
// 		fmt.Printf("Entered name is not registered in the database.")
// 		os.Exit(0)
// 	}

// 	fmt.Printf("\nEnter your Pin :")
// 	fmt.Scan(&cap)
// 	if cap != accPin[ni] {
// 		fmt.Printf("The PIN entered is incorrect! Please try again.")
// 	} else {
// 		var process_int int
		// fmt.Printf("\nChoose your process:\n1.Account Details\n2.Balance Enquiry\n3.Cash Withdrawal\n4.Cash Deposit\n5.Change Pin\n6.Exit\n")
// 		fmt.Scan(&process_int)
// 		switch process_int {
// 		case 1:
// 			accountDetails(accName, accNumber, ni)
// 		case 2:
// 			balanceEnquiry(accBalance, ni)
// 		case 3:
// 			cashWithdrawl(&accBalance, ni)
// 		case 4:
// 			cashDeposit(&accBalance, ni)
// 		case 5:
// 			changePin(&accPin, ni)
// 		case 6:
// 			os.Exit(0)
// 		default:
// 			fmt.Printf("Choose the correct process and try again.")
// 		}
// 	}
// }
