package question01

import (
	"context"
	"fmt"
	"go-demo/task/03/db"
	"log"
	"testing"
	"time"
)

func TestAutoMigrate(t *testing.T) {
	// 初始化数据库
	db.InitDB()
	defer db.CloseDB()

	// 自动建表
	err := db.DB.AutoMigrate(&Account{}, &Transactions{})

	if err != nil {
		log.Fatalf("模型自动执行数据库迁移失败：%v", err)
	}
}

func TestRegisterAccount(t *testing.T) {
	// 初始化数据库
	db.InitDB()
	defer db.CloseDB()

	account, err := RegisterAccount(999)
	if err != nil {
		log.Fatalf("注册失败：%v", err)
	}

	fmt.Println("账号注册成功。账户信息：", account)
}

func TestTransactions(t *testing.T) {
	// 初始化数据库
	db.InitDB()
	defer db.CloseDB()

	//添加一个超时上下文，用于超时交易引起事务回滚
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err1 := Transfer(ctx, 1, 2, 100)
	if err1 != nil {
		log.Fatalf("转账失败：%v", err1)
	}
}
