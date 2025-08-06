package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area()
	Perimeter()
}

// Rectangle 矩形
type Rectangle struct {
	Length int
	Width  int
}

func NewRectangle(length int, width int) Shape {
	return &Rectangle{Length: length, Width: width}
}

func (r *Rectangle) Area() {
	fmt.Println(r.Width * r.Length)
}

func (r *Rectangle) Perimeter() {
	fmt.Println(2 * (r.Length + r.Width))
}

// Circle 圆形
type Circle struct {
	Radius int
}

func NewCircle(radius int) Shape {
	return &Circle{Radius: radius}
}

func (c *Circle) Area() {
	fmt.Printf("%.2f\n", math.Pi*float64(c.Radius)*float64(c.Radius))
}

func (c *Circle) Perimeter() {
	fmt.Printf("%.2f\n", 2*math.Pi*float64(c.Radius))
}

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

type Employee struct {
	*Person
	EmployeeID int
}

func NewEmployee(person *Person, employeeID int) *Employee {
	return &Employee{Person: person, EmployeeID: employeeID}
}

func (e Employee) PrintInfo() {
	fmt.Printf("Employee<Name: %s, Age: %d, EmployeeID: %d>\n", e.Name, e.Age, e.EmployeeID)
}

func main() {
	rectangle := NewRectangle(10, 2)
	circle := NewCircle(1)

	rectangle.Area()
	rectangle.Perimeter()

	circle.Area()
	circle.Perimeter()

	person := NewPerson("Tom", 20)
	employee := NewEmployee(person, 1)
	employee.PrintInfo()
}
