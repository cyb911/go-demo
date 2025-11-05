package question02

import (
	"context"

	"github.com/jmoiron/sqlx"
)

/*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

type Book struct {
	ID     int64   `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func QueryBooksPriceGT(ctx context.Context, db *sqlx.DB, minPrice float64) ([]Book, error) {
	const q = `SELECT id, title, author, price FROM books WHERE price > :min_price ORDER BY price DESC, id ASC;`

	rows, err := db.NamedQueryContext(ctx, q, map[string]any{
		"min_price": minPrice,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Book
	for rows.Next() {
		var b Book
		if err := rows.StructScan(&b); err != nil {
			return nil, err
		}
		list = append(list, b)
	}
	return list, rows.Err()
}
