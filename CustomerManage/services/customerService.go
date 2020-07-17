package services

import (
	"CustomerManage/models"
)

// 该结构体主要完成对customer的具体操作, 包括增删改查
type CustomerService struct {
	customers   []models.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	cs := &CustomerService{}
	customer := models.NewCustomer(1, "kris", 13, "male", "13113113111", "kris@123.com")
	cs.customerNum = 1
	cs.customers = append(cs.customers, customer)
	return cs
}

func (cs *CustomerService) List() []models.Customer {
	return cs.customers
}

func (cs *CustomerService) NewCustomer(customer models.Customer) bool {
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

func (cs *CustomerService) UpdateCustomer(id int) bool {
	index := cs.FindCustumerById(id)
	if index == -1{
		return false
	}else{
		customer := cs.customers[index]
		name := ""
		fmt.Printf("Name(%s): ", customer.Name)
		fmt.Scanln(&name)
		if name!=nil {
			customer.Name = name
		}
		age := 0
		fmt.Print("Age(%v): ", customer.Age)
		fmt.Scanln(&age)
		if age!=-1{
			customer.Age = age
		}
		gender := ""
		fmt.Print("Gender(%s): ", customer.Gender)
		fmt.Scanln(&gender)
		if gender!=nil{
			customer.Gender = gender
		}
		phone := ""
		fmt.Print("Phone(%s): ", customer.Phone)
		fmt.Scanln(&phone)
		if phone!=nil{
			customer.Phone = phone
		}
		email := ""
		fmt.Print("Email(%s): ", customer.Email)
		fmt.Scanln(&email)
		if age!=0 && age!=customer.Age{
			customer.Age = age
		}
	}

	customer := models.AddCustomer(name, age, gender, phone, email)
	return cs.customers
}

func (cs *CustomerService) DeleteCustomer(id int) bool {
	index := cs.FindCustumerById(id)
	if index == -1{
		return false
	}else{
		cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
		return true
	}
}

func (cs *CustomerService) FindCustumerById(id int) int {
	for index, cus := range cs.customers {
		if cus.Id == id {
			return index
		}
	}
	return -1
}
