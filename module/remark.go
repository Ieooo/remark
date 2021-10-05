package module

import (
	"fmt"
	"joke/connection"
	"time"
)

type Remark struct {
	Id         int
	Content    string
	CreateTime *time.Time
	UpdateTime *time.Time
}

func (Remark) TableName() string {
	return "remark"
}

func GetOne() (remark *Remark) {
	remark = &Remark{}
	sql := "select * from remark order by rand() limit 1"
	connection.DB.Raw(sql).Scan(remark)
	if remark == nil {
		fmt.Print("error!")
	}
	return remark
}

func AddOne(remark *Remark) bool {
	err := connection.DB.Create(remark)
	if err != nil {
		fmt.Printf("err: %v", err)
		return false
	}
	return true
}
