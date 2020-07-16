package util

import (
	"fmt"
	"strconv"
	"time"
)

func (ac *Account) Menu() {
	key := ""
	for ac.loop {
		fmt.Println("-------------------account: " + ac.Name + "---------------------")
		fmt.Println("                   1. detail")
		fmt.Println("                   2. income")
		fmt.Println("                   3. outcome")
		fmt.Println("                   4. exit")
		fmt.Print("                   choose1~4: ")
		fmt.Scan(&key)

		switch key {
		case "1":
			ac.getDetail()
		case "2":
			ac.setIncome()
		case "3":
			ac.setOutcome()
		case "4":
			ac.exitProgram()
		default:
			fmt.Println("Please choose again:")
		}
	}
}

func (ac *Account) getDetail() {
	fmt.Println("-------------------detail-------------------")
	if ac.detail == nil {
		fmt.Println("there are no details. Please record")
		return
	}
	fmt.Println("account	money	rest	date			reason")
	for _, rec := range ac.detail {
		fmt.Println(ac.Name + "\t" + rec.money + "\t" + rec.rest + "\t" + rec.date + "\t" + rec.reason)
	}
}
func (ac *Account) setIncome() {
	fmt.Println("-------------------income-------------------")
	ac.operateRecord()
}
func (ac *Account) setOutcome() {
	fmt.Println("-------------------outcome------------------")
	ac.operateRecord()
}
func (ac *Account) operateRecord() {
	rec := new(Record)
	fmt.Print("money: ")
	fmt.Scan(&rec.money)
	fmt.Print("reason: ")
	fmt.Scan(&rec.reason)
	rec.date = time.Now().Format("2006-01-02 15:04:05")
	money, _ := strconv.Atoi(rec.money)
	if ac.total+money < 0 {
		fmt.Println("你的余额不足")
		return
	}
	ac.total += money
	rec.rest = strconv.Itoa(ac.total)
	ac.detail = append(ac.detail, *rec)
	fmt.Println(rec.money + " for " + rec.reason + " at " + string(rec.date))
}
func (ac *Account) exitProgram() {
	confirm := ""
	fmt.Print("confirm exit? (Y or N):")
	fmt.Scan(&confirm)
	if confirm == "Y" {
		fmt.Println("                   Bye Bye")
		ac.loop = false
		return
	} else {
		fmt.Println("cancel exit")
		ac.loop = true
		return
	}
}

type Account struct {
	Name   string
	total  int
	loop   bool
	detail []Record
}

type Record struct {
	money  string
	rest   string
	reason string
	date   string
}

func MyAccount() *Account {
	return &Account{
		Name:   "Kris",
		total:  1000.0,
		loop:   true,
		detail: nil,
	}
}
