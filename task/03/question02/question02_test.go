package question02

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

/*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
func TestQuery(t *testing.T) {
	// 数据库连接
	db := ConnectDB()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ---- 题目1-1：查“技术部”的所有员工 ----
	emps, err := QueryEmployeesByDept(ctx, db, "技术部")
	if err != nil {
		log.Fatalf("queryEmployeesByDept error: %v", err)
	}
	fmt.Println("技术部员工:")
	for _, e := range emps {
		fmt.Printf("ID=%d, Name=%s, Dept=%s, Salary=%.2f\n", e.ID, e.Name, e.Department, e.Salary)
	}

	// ---- 题目1-2：查工资最高的员工 ----
	topEmp, err := QueryTopPaidEmployee(ctx, db)
	if err != nil {
		log.Fatalf("queryTopPaidEmployee error: %v", err)
	}
	fmt.Println("\n【工资最高的员工】")
	fmt.Printf("ID=%d, Name=%s, Dept=%s, Salary=%.2f\n", topEmp.ID, topEmp.Name, topEmp.Department, topEmp.Salary)
}

/*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

func TestQueryBooksPriceGTQuery(t *testing.T) {
	// 数据库连接
	db := ConnectDB()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ---- 题目2：查价格>50的书籍（类型安全映射）----
	books, err := QueryBooksPriceGT(ctx, db, 50)
	if err != nil {
		log.Fatalf("queryBooksPriceGT error: %v", err)
	}
	fmt.Println("价格>50的书籍:")
	for _, b := range books {
		fmt.Printf("ID=%d, Title=%s, Author=%s, Price=%.2f\n", b.ID, b.Title, b.Author, b.Price)
	}

}
