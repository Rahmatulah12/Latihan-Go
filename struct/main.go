package main

import "fmt"

func main() {
	var student Student
	student.FirstName = "James"
	student.LastName = "Bond"
	student.Age = 30
	fmt.Println(student)

	var student1 Student = Student{"James", "Bond", 30}
	fmt.Println(student1)

	var student2 = Student{}
	student2.FirstName = "James"
	student2.LastName = "Bond"
	student2.Age = 30
	fmt.Println(student2)

	// Variable Objek Pointer
	var student3 = Student{"James", "Bond", 30}

	var student4 *Student = &student3
	fmt.Println(student3)
	fmt.Println(student4)

	var employee Employee
	employee.Name = "James Bond"
	employee.Age = 30
	employee.Salary = 30000
	employee.Position = "Agent 007"
	fmt.Println(employee)

	var employee1 Employee = Employee{Person{"James Bond", 30}, "Agent 007", 30000}
	fmt.Println(employee1)

	var student5 = Student{
		FirstName: "James",
		LastName:  "Bond",
		Age:       30,
	}
	fmt.Println(student5)

	var newEmployee = Employee1{}
	newEmployee.Person1.Name = "James Bond"
	newEmployee.Age = 30
	newEmployee.Person1.NIK = "007"
	newEmployee.Position = "Agent 007"
	newEmployee.Salary = 30000
	newEmployee.NIK = "007"
	fmt.Println(newEmployee)

	// Anonymous Struct
	var AnonymousStruct = struct {
		FirstName string
		LastName  string
	}{}
	AnonymousStruct.FirstName = "James"
	AnonymousStruct.LastName = "Bond"
	fmt.Println(AnonymousStruct)

	// cara lain penulisan Anonymous Struct
	var anonymStruct struct {
		FirstName string
	}
	anonymStruct.FirstName = "James"
	fmt.Println(anonymStruct)

	var test1 = struct{ name string }{name: "James Bond"}
	var test2 = struct{ age int }{age: 30}
	fmt.Println(test1, test2)

	// kombinasi struct dan slice
	var Persons = []Person{
		{Name: "James Bond", Age: 30},
		{Name: "Diana Prince", Age: 25},
		{Name: "Maria Sharapova", Age: 25},
		{Name: "Dwayne Johnson", Age: 35},
	}

	for _, person := range Persons {
		fmt.Println("My name is", person.Name, "and I'm", person.Age, "years old")
	}

}

// Kapital pada Go adalah public
type Student struct {
	FirstName string
	LastName  string
	Age       int
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Position string
	Salary   int
}

type Person1 struct {
	NIK  string
	Name string
	Age  int
}

type Employee1 struct {
	Person1
	NIK      string
	Position string
	Salary   int
}

// Nested Struct
type Vehicles struct {
	Brand struct {
		Name string
	}
	Model       string
	Years       []string
	Manufacture struct {
		Country struct {
			Name string
		}
	}
}

// Deklarasi Struct Secara Horizontal
type Horizontal struct {
	first  int
	second string
	third  float32
}

// Tags pada property struct
type Person3 struct {
	Name string `name`
	Age  int    `age`
}

// Type Alias, Sebuah tipe data, seperti struct, bisa dibuatkan alias baru, caranya dengan type NamaAlias = TargetStruct
type People = Person
