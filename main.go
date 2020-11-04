package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

type Member struct {
	Id int `db:"id"`
	Username string `db:"username"`
	Gender int `db:"gender"`
	Age int `db:"age"`
	Score float32 `db:"score"`
	Gonde int `db:"gongde"`
}


//首先要初始化db，让后赋值给Db
func init()  {
	database,err := sqlx.Open("mysql","root:xxx@tcp(127.0.0.1:43306)/xxx")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
}


func InsertData(){
	Db.Exec("insert into member(username,gender,age,score,gongde) values('肯定',2,3,4,5)")
}

func SelectData(){
	var member []Member
	err := Db.Select(&member,"select * from member")
	if err != nil {
		panic(err)
	}

	fmt.Println(member)
}

func UpdateData(){
	Db.Exec("update member set gongde=99 where id=4")
}

func DeleteData(){
	res,err := Db.Exec("delete from member where id=5")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T",res)
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
}

func main() {
	//插入数据
	InsertData()

	//查询数据
	SelectData()

	//修改数据
	UpdateData()

	//删除数据
	DeleteData()
}