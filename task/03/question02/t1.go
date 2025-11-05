package question02

import (
	"context"

	"github.com/jmoiron/sqlx"
)

/*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/

type Employee struct {
	ID         string `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

func QueryEmployeesByDept(ctx context.Context, db *sqlx.DB, dept string) ([]Employee, error) {
	const q = `SELECT id, name, department, salary FROM employees WHERE department = :dept ORDER BY id ASC;`

	rows, err := db.NamedQueryContext(ctx, q, map[string]any{
		"dept": dept,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Employee
	for rows.Next() {
		var e Employee
		if err := rows.StructScan(&e); err != nil {
			return nil, err
		}
		list = append(list, e)
	}
	return list, rows.Err()
}

func QueryTopPaidEmployee(ctx context.Context, db *sqlx.DB) (Employee, error) {
	const q = `SELECT id, name, department, salary FROM employees ORDER BY salary DESC, id ASC LIMIT 1;`
	var e Employee
	// Get 只拿一行，找不到会返回 sql.ErrNoRows
	if err := db.GetContext(ctx, &e, q); err != nil {
		return Employee{}, err
	}
	return e, nil
}
