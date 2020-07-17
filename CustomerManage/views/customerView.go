package main

import (
	"CustomerManage/models"
	"CustomerManage/services"
	"fmt"
)

func main() {
	cv := CustomerView{key: "", loop: true}
	cv.cs = services.NewCustomerService()
	cv.mainMenu()
}

type CustomerView struct {
	key  string
	loop bool

	cs *services.CustomerService
}

func (cv *CustomerView) mainMenu() {
	for cv.loop {
		fmt.Println("--------------------CustomerManageSystem--------------------")
		fmt.Println("                       1. add customer")
		fmt.Println("                       2. update customer")
		fmt.Println("                       3. delete customer")
		fmt.Println("                       4. show customers")
		fmt.Println("                       5. exit system")
		fmt.Print("                       Please choose: ")
		fmt.Scanln(&cv.key)

		switch cv.key {
		case "1":
			fmt.Println("------------------------add customer------------------------")
			cv.add()
		case "2":
			fmt.Println("-----------------------update customer----------------------")
			cv.update()
		case "3":
			fmt.Println("-----------------------delete customer----------------------")
			cv.delete()
		case "4":
			fmt.Println("------------------------show customer-----------------------")
			cv.show()
		case "5":
			cv.loop = false
			fmt.Println("---------------------------bye bye--------------------------")
		}
	}
}

func (cv *CustomerView) show() {
	fmt.Println("ID\tName\tAge\tGender\tPhone\t\tEmail")
	for _, cus := range cv.cs.List() {
		fmt.Println(cus.GetInfo())
	}
	fmt.Printf("-----------------------customer list------------------------\n\n")
}

func (cv *CustomerView) add() {
	name := ""
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	age := 0
	fmt.Print("Age: ")
	fmt.Scanln(&age)
	gender := ""
	fmt.Print("Gender: ")
	fmt.Scanln(&gender)
	phone := ""
	fmt.Print("Phone: ")
	fmt.Scanln(&phone)
	email := ""
	fmt.Print("Email: ")
	fmt.Scanln(&email)
	customer := models.AddCustomer(name, age, gender, phone, email)
	if cv.cs.NewCustomer(customer) {
		fmt.Printf("--------------------------success---------------------------\n\n")
	} else {
		fmt.Printf("--------------------------failed----------------------------\n\n")
	}
}

func (cv *CustomerView) update() {
	id := 0
	fmt.Print("Choose customer Id: ")
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Printf("--------------------------cancel----------------------------\n\n")
		return
	} else {
		if cv.cs.UpdateCustomer(id) {
			fmt.Printf("--------------------------success---------------------------\n\n")
		} else {
			fmt.Printf("--------------------------failed----------------------------\n\n")
		}
	}
}

func (cv *CustomerView) delete() {
	id := 0
	fmt.Print("Choose customer Id: ")
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Printf("--------------------------cancel----------------------------\n\n")
		return
	} else {
		if cv.cs.DeleteCustomer(id) {
			fmt.Printf("--------------------------success---------------------------\n\n")
		} else {
			fmt.Printf("--------------------------failed----------------------------\n\n")
		}
	}
}
