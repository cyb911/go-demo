package question01

import (
	"context"
	"errors"
	"go-demo/task/03/db"

	"gorm.io/gorm"
)

/*
事务语句 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键，
from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。 要求 ： 编写一个事务，实现从账户 A
向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加
100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

type Account struct {
	gorm.Model
	Balance int64 `gorm:"not null;check:balance>=0"` // 余额（分）
}

type Transactions struct {
	gorm.Model
	FromAccountID uint
	ToAccountID   uint
	Amount        int64 // 金额（分）
}

func Transfer(ctx context.Context, fromID, toID uint, amount int64) error {
	if amount <= 0 {
		return errors.New("交易金额不能小于0")
	}

	if fromID == toID {
		return errors.New("转入转出账户id不能相同！")
	}

	return db.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var to Account
		if err := tx.Select("id").First(&to, toID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("收款账户不存在")
			}
			return err
		}

		res := tx.Model(&Account{}).
			Where("id = ? AND balance >= ?", fromID, amount).
			Update("balance", gorm.Expr("balance - ?", amount))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return errors.New("付款账户不存在或余额不足")
		}

		if err := tx.Model(&Account{}).
			Where("id = ?", toID).
			Update("balance", gorm.Expr("balance + ?", amount)).
			Error; err != nil {
			return err
		}

		rec := &Transactions{
			FromAccountID: fromID,
			ToAccountID:   toID,
			Amount:        amount,
		}
		if err := tx.Create(rec).Error; err != nil {
			return err
		}
		return nil

	})
}

func RegisterAccount(balance int64) (*Account, error) {
	if balance <= 0 {
		return nil, errors.New("初始金额不能小于0")
	}
	tx := db.DB
	account := &Account{Balance: balance}
	err := tx.Create(account).Error
	if err != nil {
		return nil, err
	}

	return account, nil
}
