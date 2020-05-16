package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Album struct {
	Id         int
	Filepath   string
	Filename   string
	Status     int
	Createtime int64
}

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
}

type User struct {
	Id            int
	Username      string
	Password      string
	Status        int
	Createtime    int64
}

func init() {
	fmt.Println("Register models into database")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Album), new(Article))
}

