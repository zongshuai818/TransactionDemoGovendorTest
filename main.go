package main

import (
	Btc "github.com/zongshuai818/Bitcion"
	"fmt"
        "github.com/lxdeng/golibstrings"
)

func main() {
	fmt.Println("Init account...")
        fmt.Println(golibstrings.Reverse("!oG ,olleH"))
	var MyBTC Btc.BtcAccount
	MyBTC.Cache.Hash = make(map[string]string)
	_ = MyBTC.Check()
	MyBTC.Run()
}
