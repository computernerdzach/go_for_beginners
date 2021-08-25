package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {
	jack := Employee{
		Name:     "Jack Smith",
		Age:      27,
		Salary:   40000,
		FullTime: false,
	}

	jill := Employee{
		Name:     "Jill Jones",
		Age:      33,
		Salary:   60000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, x := range employees {
		if x.Age > 30 {
			fmt.Println(x.Name, "is 30 or older")
		} else {
			fmt.Println(x.Name, "is under 30")
		}

		if x.Age > 30 && x.Salary > 50000 {
			fmt.Println(x.Name, "makes more than 50,000 and is over 30")
		} else {
			fmt.Println("Either", x.Name, "makes less than 50,000 or is over 30")
		}

		if x.Age > 30 || x.Salary > 50000 {
			fmt.Println(x.Name, "makes more than 50,000 or is over 30")
		} else {
			fmt.Println(x.Name, "makes less than 50,000 and is under 30")
		}

		if (x.Age > 30 || x.Salary < 50000) && x.FullTime {
			fmt.Println(x.Name, "matches our unclear criteria")
		}
	}
}
