package Bitcoin

import (
	Ccl "github.com/zongshuai818/Cachelab"
	"fmt"
	"strconv"
)

type BtcAccount struct {
	Total int
	Cache Ccl.Cache
}

// in
func (BA *BtcAccount) Check() int {
	fmt.Printf("Current account balance:%d\n", BA.Total)
	return BA.Total
}

func (BA *BtcAccount) ChangeIn(name string, number int) bool {
	if number > 2000 || number < 1000 {
		fmt.Println("ChangeIn fail. Value should be between 1000 and 2000!")
		return false
	}
	// Record
	BA.Total += number
	BA.Cache.Set(name, "+"+strconv.Itoa(number))
	fmt.Println("ChangeIn successful")
	return true
}

// out
func (BA *BtcAccount) ChangeOut(name string, number int) bool {
	if BA.Total < number {
		fmt.Println("ChangeOut fail. Insufficient account balance!")
		return false
	}
	BA.Total -= number
	BA.Cache.Set(name, "-"+strconv.Itoa(number))
	fmt.Println("ChangeOut successful")
	return true
}

// Inquire
func (BA *BtcAccount) InquiryTransfer(name string) (string, bool) {
	transfer, get := BA.Cache.Get(name)
	return transfer, get
}

// Run this account
func (BA *BtcAccount) Run() {
	var opt, amount int
	var account string

	for true {
		fmt.Printf("Select a operation(1=Change into, 2=change out, 3=check amount, 4=check transfer, 0=exit): ")
		_, _ = fmt.Scanln(&opt)
		switch opt {
		case 1: //in
			{
				fmt.Printf("Please enter the account name and amount:")
				_, _ = fmt.Scanln(&account, &amount)
				BA.ChangeIn(account, amount)
			}
		case 2: //out
			{
				fmt.Printf("Please enter the account name and amount:")
				_, _ = fmt.Scanln(&account, &amount)
				BA.ChangeOut(account, amount)
			}
		case 3: //check amount
			{
				fmt.Println("Account balance:", BA.Total)
			}
		case 4: //check transfer
			{
				fmt.Printf("Please enter account name:")
				_, _ = fmt.Scanln(&account)
				transfer, get := BA.InquiryTransfer(account)
				if get {
					fmt.Println("Transfer record:", account, transfer)
				} else {
					fmt.Println("No record")
				}
			}
		case 0: //exit
			goto label
		}
	}
label:
	fmt.Println("Exit！")
}
