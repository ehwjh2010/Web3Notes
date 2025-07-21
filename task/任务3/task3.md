Sqlx入门</br>
题目1：使用SQL扩展库进行查询</br>
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。</br>
要求 ：</br></br>
	编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。</br></br>
	编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。</br>

题目2：实现类型安全映射</br>
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。</br>
要求 ：</br>
	定义一个 Book 结构体，包含与 books 表对应的字段。</br>
	编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。</br>

```go
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var sqlxDB *sqlx.DB

func InitDB() {
	dsn := "jh:123456@tcp(127.0.0.1:10016)/awe_db?parseTime=true"
	sqlxDB = sqlx.MustOpen("mysql", dsn)
}

type Employees struct {
	ID         uint
	Name       string
	Department string
	Salary     float64
}

func (e *Employees) String() string {
	return fmt.Sprintf("Employees<ID: %d, Name: %s, Department: %s, Salary: %f>",
		e.ID, e.Name, e.Department, e.Salary,
	)
}

type Books struct {
	ID     uint
	Title  string
	Author string
	Price  float64
}

func (b *Books) String() string {
	return fmt.Sprintf("Book<ID: %d, Title: %s, Author: %s, Price: %f>",
		b.ID, b.Title, b.Author, b.Price,
	)
}

// QueryTechnicalEmployees 查询所有技术部的员工
func QueryTechnicalEmployees() {
	employees := make([]*Employees, 0)
	err := sqlxDB.Select(&employees, `select * from employees where department = ?`, "技术部")
	if err != nil {
		log.Printf("查询技术部员工失败, err: %s\n", err)
		return
	}

	fmt.Printf("%+v\n", employees)
}

// QueryHighestSalaryEmployees 查询工资最高的员工
func QueryHighestSalaryEmployees() {
	var employees Employees
	err := sqlxDB.Get(&employees, `select * from employees order by salary desc limit 1`)
	if err != nil {
		log.Printf("查询收入最高的员工失败, err: %s\n", err)
		return
	}

	fmt.Printf("%+v\n", employees)
}

func QueryBooks() {
	books := make([]*Books, 0)
	err := sqlxDB.Select(&books, `select *  from books where price > 50`)
	if err != nil {
		log.Printf("查询价格大于50的书籍失败, err: %s\n", err)
		return
	}

	fmt.Printf("%+v\n", books)
}

func main() {
	InitDB()

	QueryTechnicalEmployees()
	QueryHighestSalaryEmployees()
	QueryBooks()

	sqlxDB.Close()
}

```