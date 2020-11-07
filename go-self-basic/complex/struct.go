package main

import "time"

/**
struct 结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。
在java当中相当于对象

 */

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main(){
	dilbert.Salary -= 5000
	position := &dilbert.Position
	*position = "Senior " + *position


}