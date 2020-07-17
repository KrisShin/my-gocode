package models

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Age    int
	Gender string
	Phone  string
	Email  string
}

func NewCustomer(id int, name string, age int, gender string, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Age:    age,
		Gender: gender,
		Phone:  phone,
		Email:  email,
	}
}

func AddCustomer(name string, age int, gender string, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Age:    age,
		Gender: gender,
		Phone:  phone,
		Email:  email,
	}

}

func (cus *Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%s\t%v\t%s\t%s\t%s", cus.Id, cus.Name, cus.Age, cus.Gender, cus.Phone, cus.Email)
}
