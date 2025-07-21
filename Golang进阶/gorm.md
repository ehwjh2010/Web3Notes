## Gorm

Tag

![gormTag.png](..%2Fimages%2FgormTag.png)


1. 初始化Gorm并自动迁移
```go
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Description string `gorm:"column:description;type:varchar(128);not null;default:''" json:"description"` // 描述
// ElementType int32  `gorm:"column:element_type;type:tinyint;not null;default:0" json:"element_type"`     // 元素类型

type Person struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(128);index:idx_name;not null;default:''" json:"name"`
	Age  uint8  `gorm:"column:age;type:tinyint unsigned;not null;default:0" json:"age"`
}

func InitGorm() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: `jh:123456@tcp(127.0.0.1:10016)/awe_db?charset=utf8mb4&parseTime=true&loc=Local`,
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{SingularTable: true},
	})

	if err != nil {
		panic(err)
	}
	return db
}

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&Person{}); err != nil {
		panic(err)
	}
}

```

## 钩子(Hook)

1. 创建时可用的 hook
```markdown
// 开始事务
BeforeSave
BeforeCreate
// 关联前的 save
// 插入记录至 db
// 关联后的 save
AfterCreate
AfterSave
// 提交或回滚事务
```