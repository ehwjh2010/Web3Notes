题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
    编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
    在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
    向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
    如果余额不足，则回滚事务。

```go
package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Accounts struct {
	ID      uint
	Balance float64 `gorm:"type:decimal(10,2) unsigned;NOT NULL;default 0;comment:余额" json:"balance"`
}

type Transactions struct {
	ID            uint
	FromAccountId uint
	ToAccountId   uint
	Amount        float64 `gorm:"type:decimal(10,2) unsigned;NOT NULL;default 0;comment:转账金额"`
}

func InitGorm() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: `jh:123456@tcp(127.0.0.1:10016)/awe_db?charset=utf8mb4&parseTime=true&loc=Local`,
		//DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy:  &schema.NamingStrategy{SingularTable: true},
		CreateBatchSize: 1000,
		//Logger:          logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(&Accounts{}, &Transactions{}); err != nil {
		panic(err)
	}

	return db
}

func Transfer(db *gorm.DB, from uint, to uint, amount float64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var fromAccount Accounts
		err := tx.Model(&Accounts{}).Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", from).Take(&fromAccount).Error
		if err != nil {
			return err
		}
		if fromAccount.Balance < amount {
			return errors.New("from is insufficient balance")
		}

		var toAccount Accounts
		err = tx.Model(&Accounts{}).Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", to).Take(&toAccount).Error
		if err != nil {
			return err
		}

		err = tx.Model(&Accounts{}).Where("id = ?", from).Update("balance", gorm.Expr("balance - ?", amount)).Error
		if err != nil {
			return err
		}
		err = tx.Model(&Accounts{}).Where("id = ?", to).Update("balance", gorm.Expr("balance + ?", amount)).Error
		if err != nil {
			return err
		}

		err = tx.Create(&Transactions{
			FromAccountId: from,
			ToAccountId:   to,
			Amount:        amount,
		}).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func main() {
	db := InitGorm()

	// 添加 账户A id = 1, 账户B id = 2
	if err := db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create([]*Accounts{{ID: 1, Balance: 100}, {ID: 2, Balance: 50}}).Error; err != nil {
		panic(err)
	}

	fmt.Println(Transfer(db, 1, 2, 100))
}
```
