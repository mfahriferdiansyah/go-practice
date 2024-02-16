package main

import "fmt"

func (s Student) GetIdentity() {
	fmt.Printf("ID : %d \n", s.ID)
	fmt.Printf("Name : %s \n", s.Name)
	fmt.Printf("Address : %s \n", s.Address)
	fmt.Printf("Job : %s \n", s.Job)
	fmt.Printf("Motivation : %s \n", s.Motivation)
}
