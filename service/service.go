package service

import (
	"encoding/json"
	 "github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"github.com/monkeyliu/rest/model"
	"github.com/monkeyliu/rest/dbinit"
	"fmt"
)
func GETuser(c *gin.Context) {
	var user model.Users
	var idType model.IDtype
	data, err := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &idType)
	id := idType.ID
	fmt.Printf(id)
	var res gin.H
	db := dbinit.DbInit()
	fmt.Printf("数据库链接成功")
	row := db.QueryRow("select id, name,age,sex from users where id = ?;", id)
	err = row.Scan(&user.ID, &user.NAME, &user.AGE, &user.SEX)
	if err != nil {
		res = gin.H{
			"result": nil,
		}
		c.JSON(http.StatusNotFound,res)
	} else {
		res :=gin .H{
			"result": user,

		}
		c.JSON(http.StatusOK, res)
	}
	db.Close()
}
func DELETEuser(c *gin.Context) {
	var user  model.Users
	var idType model.IDtype
	data, err := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &idType)
	var res gin.H
	id1 := user.ID
	id := idType.ID
	fmt.Printf(id)
	fmt.Printf(id1)
	db := dbinit.DbInit()
	//查询需要删除的id是否在表中存在,存在才能删，不存在返回错误信息
	row := db.QueryRow("select id, name,age,sex from users where id = ?;", id)
	err = row.Scan(&user.ID, &user.NAME, &user.AGE, &user.SEX)
	if err != nil {
		res =gin .H{
			"code":  0,
			"result": id,
		}
		c.JSON(http.StatusNotFound,res)
	} else {
		res = gin.H{
			"result": "DELETE SUCCESS",
		}
		db.QueryRow("delete from users where id = ?;", id)
		c.JSON(http.StatusOK, res)
	}
		db.Close()
}
func CREATEuser(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	db := dbinit.DbInit()
	var user model.Users
	json.Unmarshal(data, &user)
	var res gin .H
	id := user.ID
	//查询所增加的data的id是否存在，存在无法增加，返回错误信息，不存在才增加信息
	row := db.QueryRow("select id, name,age,sex from users where id = ?;", id)
	err = row.Scan(&user.ID, &user.NAME, &user.AGE, &user.SEX)
	if err != nil {
		res =gin .H{
			"result": user,
		}
		id := user.ID
		name := user.NAME
		age := user.AGE
		sex := user.SEX
		db.QueryRow("insert into users (ID,NAME,AGE,SEX) values (?,?,?,?)", id, name, age, sex)
		c.JSON(http.StatusOK, res)
	} else {
		res = gin.H{
			"result": user,
		}
		c.JSON(http.StatusNotFound, res)
	}
	db.Close()
}
func UPDATEuser(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	db := dbinit.DbInit()
	var user model.Users
	json.Unmarshal(data, &user)



	id := user.ID
	name := user.NAME
	age := user.AGE
	sex := user.SEX

	fmt.Printf(name)
	fmt.Printf(id)
	var res gin.H
	//查询更新信息的Id是否存在，存在才跟新，不存在无法更新
	row := db.QueryRow("select id, name,age,sex from users where id = ?;", id)
	err = row.Scan(&user.ID, &user.NAME, &user.AGE, &user.SEX)
	err=nil
		if err != nil {
			res= gin.H{
				"result": "所更改的列信息不存在",
			}
			c.JSON(http.StatusNotFound, res)
		} else {
			db.QueryRow("update users set id=?,name=?,age=?,sex=?  where id = ?", id ,name, age, sex,id)
			res =gin .H{
				"result": user,
			}



		c.JSON(http.StatusOK, res)
	}
	db.Close()
}
