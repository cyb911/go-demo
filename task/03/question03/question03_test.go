package question02

import (
	"fmt"
	"go-demo/task/03/db"
	"go-demo/task/03/question03/models"
	"log"
	"testing"

	"gorm.io/gorm"
)

/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

// 使用Gorm创建这些模型对应的数据库表
func TestAutoMigrate(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	// 建表
	if err := db.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		log.Fatalf("自动建表失败: %v", err)
	}
	fmt.Println("数据表创建完成")
}

// 初始化数据
func TestInitDate(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	var count int64
	db.DB.Model(&models.User{}).Count(&count)
	if count > 0 {
		return
	}

	u := models.User{Username: "Amor"}
	db.DB.Create(&u)

	p1 := models.Post{Title: "GORM 入门", Content: "第一章 CURD", UserID: u.ID}
	p2 := models.Post{Title: "GORM 进阶", Content: "ORM框架的原理之GORM", UserID: u.ID}
	db.DB.Create(&p1)
	db.DB.Create(&p2)

	c1 := models.Comment{Content: "good！", PostID: p1.ID, UserID: u.ID}
	c2 := models.Comment{Content: "I was amazed.", PostID: p1.ID, UserID: u.ID}
	db.DB.Create(&c1)
	db.DB.Create(&c2)
}

// 关联查询
func TestRefQuery(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	// 查询用户文章及评论
	u, err := GetUserPostsWithComments(2)
	if err != nil {
		log.Printf("查询用户文章失败: %v", err)
	} else {
		fmt.Printf("用户 %s 的文章数：%d\n", u.Username, len(u.Posts))
		for _, p := range u.Posts {
			fmt.Printf("  -《%s》评论：\n", p.Title)
			for _, comment := range p.Comments {
				fmt.Printf("  用户id：%d-->：%s\n", comment.UserID, comment.Content)
			}
		}
	}

	// 评论数最多的文章
	post, err := GetMostCommentedPost()
	if err != nil {
		log.Printf("查询最多评论文章失败: %v", err)
	} else {
		fmt.Printf("评论最多的文章：《%s》，评论数：%d\n", post.Title, post.CommentCount)
	}

}

func TestHookCreate(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()
	c1 := models.Comment{Content: "Test:good！", PostID: uint(1), UserID: uint(1)}
	db.DB.Create(&c1)
}

// 删除评论时，hook调用
func TestHookDel(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	err := db.DB.Delete(&models.Comment{Model: gorm.Model{ID: 4}}).Error
	if err != nil {
		log.Printf("删除评论失败: %v", err)
	}
}
