
进阶gorm</br>

题目1：模型定义</br>
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。</br>
要求 ： </br>
    使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。</br></br>
    编写Go代码，使用Gorm创建这些模型对应的数据库表。</br></br>
题目2：关联查询</br>
基于上述博客系统的模型定义。</br>
要求 ： </br>
    编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。 </br>
    编写Go代码，使用Gorm查询评论数量最多的文章信息。</br></br>
题目3：钩子函数</br>
继续使用博客系统的模型。</br>
要求 ：</br>
    为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。</br>
    为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。</br>

```go
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// User （用户）、 Post （文章）、 Comment （评论）
// 使用Gorm定义 User 、 Post 和 Comment 模型，
//	其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）
//	Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表

type User struct {
	ID        uint
	Name      string `gorm:"type:varchar(128);not null;default:'';comment:用户名"`
	Posts     []Post
	PostCount int
}

type Post struct {
	ID            uint
	Title         string `gorm:"type:varchar(128);not null;default:'';comment:标题"`
	UserID        uint
	Comments      []Comment
	CommentCount  int
	CommentStatus string
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + ?", 1)).
		Error
}

type Comment struct {
	ID      uint
	Content string `gorm:"type:varchar(1024);not null;default:'';comment:内容"`
	PostID  uint
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(&Post{}).Where("id = ?", c.PostID).
		Updates(map[string]interface{}{
			"comment_count":  gorm.Expr("comment_count + ?", 1),
			"comment_status": gorm.Expr("case when comment_count + 1 <= 0 then '无评论' when comment_count + 1 > 0 then '有评论' end"),
		}).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	err = tx.Debug().Model(&Post{}).Where("id = ?", c.PostID).
		Updates(map[string]interface{}{
			"comment_count":  gorm.Expr("comment_count - ?", 1),
			"comment_status": gorm.Expr("case when comment_count - 1 <= 0 then '无评论' when comment_count - 1 > 0 then '有评论' end"),
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func DelPostAllComment(db *gorm.DB, name string) error {
	var post Post
	err := db.Model(&Post{}).Preload("Comments").Where("title = ?", name).Take(&post).Error
	if err != nil {
		return err
	}

	if err = db.Delete(post.Comments).Error; err != nil {
		return err
	}

	return nil
}

func InitGormDB() *gorm.DB {
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

	return db
}

// GetDestUser 用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func GetDestUser(db *gorm.DB, name string) (*User, error) {
	var user User
	err := db.Model(&User{}).Preload("Posts.Comments").Where("name = ?", name).Take(&user).Error
	if err != nil {
		return nil, err
	}

	fmt.Println(user.Name)
	for _, post := range user.Posts {
		fmt.Println("\t", post.Title)
		for _, comment := range post.Comments {
			fmt.Println("\t\t", comment.Content)
		}
	}

	return &user, nil
}

func FindMostCommentPost(db *gorm.DB) error {
	var postID uint
	err := db.Model(&Comment{}).Select("post_id").Group("post_id").Order("COUNT(1) desc").Limit(1).Scan(&postID).Error
	if err != nil {
		return err
	}

	var post Post
	err = db.Model(&Post{}).Preload("Comments").Where("id = ?", postID).First(&post).Error
	if err != nil {
		return err
	}

	fmt.Println(post.Title)
	for _, comment := range post.Comments {
		fmt.Println("\t", comment.Content)
	}
	return nil
}

func SetupData(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		panic(err)
	}

	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Comment{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Post{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})

	err := db.Model(&User{}).Create([]*User{
		{
			Name: "天蚕土豆",
			Posts: []Post{
				{
					Title: "斗破苍穹",
					Comments: []Comment{
						{
							Content: "好看!!!",
						},
						{
							Content: "🐮B",
						},
					},
				},
				{
					Title: "大主宰",
					Comments: []Comment{
						{
							Content: "惊险，高潮迭起!!!",
						},
						{
							Content: "扮猪吃老虎",
						},
						{
							Content: "刺激，让人着迷😍",
						},
						{
							Content: "一般般",
						},
					},
				},
			},
		},
		{
			Name: "唐家三少",
			Posts: []Post{
				{
					Title: "斗罗大陆2",
					Comments: []Comment{
						{
							Content: "小白文",
						},
						{
							Content: "好看, 新奇",
						},
						{
							Content: "一般般",
						},
					},
				},
				{
					Title: "神印王座",
					Comments: []Comment{
						{
							Content: "老一套",
						},
						{
							Content: "扮猪吃老虎",
						},
						{
							Content: "没啥新意",
						},
					},
				},
			},
		},
		{
			Name: "辰东",
			Posts: []Post{
				{
					Title: "遮天",
					Comments: []Comment{
						{
							Content: "支持新作",
						},
						{
							Content: "新体系，有意思",
						},
						{
							Content: "先试试水",
						},
					},
				},
			},
		},
	}).Error

	if err != nil {
		panic(err)
	}
}

func main() {
	db := InitGormDB()

	SetupData(db)
	GetDestUser(db, "天蚕土豆")
	FindMostCommentPost(db)
	DelPostAllComment(db, "遮天")
}
```