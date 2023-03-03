package CreateAccount

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path string = "CreateAccount/data.json"

type Customer struct {
	Acc_num  int    `json:"Account number"`
	Cus_name string `json:"Account Holder Name"`
	Cus_age  int    `json:"Account Holder Age"`
	Acc_bal  int    `json:"Balance"`
	Acc_pin  int    `json:"Pin"`
}

func strtoi(s string) int64 {
	a, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 0)
	return a
}

func check_error(e error) {
	if e != nil {
		fmt.Println("Error occured:", e)
	}
}

func check_file() {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err = os.Create(path)
		check_error(err)
	}
}

func generate_acc_num() int {
	var a []int

	file, err := os.ReadFile("CreateAccount/acc_num.json")
	check_error(err)
	err = json.Unmarshal(file, &a)
	check_error(err)

	v := a[0]
	a[0] = v + 1

	fl_pt, err := os.OpenFile("CreateAccount/acc_num.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	check_error(err)
	defer fl_pt.Close()

	err = json.NewEncoder(fl_pt).Encode(a)
	check_error(err)

	return v
}

func addToJSON(c Customer) {
	var data_list []Customer

	file, err := os.ReadFile(path)
	check_error(err)

	f, err := os.Stat(path)
	check_error(err)

	if f.Size() == 0 {
		fl_pt, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		check_error(err)
		defer fl_pt.Close()

		data_list = append(data_list, c)

		err = json.NewEncoder(fl_pt).Encode(data_list)
		check_error(err)
		return
	}

	err = json.Unmarshal(file, &data_list)
	check_error(err)
	fmt.Println(1)

	data_list = append(data_list, c)

	fl_pt, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	check_error(err)
	defer fl_pt.Close()

	err = json.NewEncoder(fl_pt).Encode(data_list)
	check_error(err)

}

func CreateAcc() {
	var c Customer
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name:")
	a1, _ := reader.ReadString('\n')
	c.Cus_name = strings.TrimRight(a1, "\r\n")

	fmt.Print("Enter age:")
	a2, _ := reader.ReadString('\n')
	c.Cus_age = int(strtoi(a2))

	//    Age check
	if c.Cus_age < 18 && c.Cus_age>100{
		fmt.Println("Sorry, you can't have a bank account!")
		return
	} 

	// fmt.Print("Enter account number:")
	// a3, _ := reader.ReadString('\n')
	c.Acc_num = generate_acc_num() //	int(strtoi(a3))

	// fmt.Print("Enter account balance:")
	// a4, _ := reader.ReadString('\n')
	c.Acc_bal = 0 //int(strtoi(a4))

	fmt.Print("Enter pin:")
	a5, _ := reader.ReadString('\n')
	c.Acc_pin = int(strtoi(a5))

	check_file()

	addToJSON(c)
}
