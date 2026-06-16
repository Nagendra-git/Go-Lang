package main

import (
	"fmt"
)

type Organization struct {
	Id        int
	Name      string
	Employees []*Employee
}

type Employee struct {
	Id   int
	Name string
	Org  *Organization
}

func main() {
	google := &Organization{Id: 1, Name: "google"}

	nagendra := &Employee{Id: 100, Name: "Nagendra", Org: google}
	vinay := &Employee{Id: 101, Name: "Vinay", Org: google}

	google.Employees = append(google.Employees, nagendra, vinay)

	microsoft := &Organization{Id: 2, Name: "microsoft"}

	sonu := &Employee{Id: 200, Name: "sonu", Org: microsoft}
	banu := &Employee{Id: 201, Name: "banu", Org: microsoft}

	microsoft.Employees = append(microsoft.Employees, sonu, banu)

	fmt.Println(nagendra.Name)
	fmt.Println(nagendra.Org.Name)
	google.Name = "Google Inc."
	fmt.Println(nagendra.Org.Name)

	fmt.Println(google.Employees[0].Name)
	nagendra.Name = "Nagendra Burusu"
	fmt.Println(google.Employees[0].Name)
}
