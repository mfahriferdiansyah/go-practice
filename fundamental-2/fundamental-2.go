package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]
	stringArg := strings.Join(arg, " ")

	indexTarget, err := checkName(stringArg)

	if err != nil {
		fmt.Println("==========================================")
		fmt.Println("Student name not found, available student: ")
		for _, s := range StudentList {
			fmt.Printf("Name : %s \n", s.Name)
		}

	} else {
		StudentList[indexTarget].GetIdentity()
	}
}

func checkName(name string) (int, error) {
	var indexTarget int
	flagFound := false
	for index, s := range StudentList {
		if s.Name == name {
			indexTarget = index
			flagFound = true
		}
	}

	if !flagFound {
		return -1, fmt.Errorf("Student name not found (%s)", name)
	}

	return indexTarget, nil
}
