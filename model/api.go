package model

import (
	db "todoList/database"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id           int    `json:"新增順序"`
	Dolist       string `json:"要做什麼呢?"`
	Sort         string `json:"分類"`
	Important    int    `json:"重要程度"`
	Last_update1 int64  `json:"創建時間" form:"time"`
	Last_update2 int64  `json:"更新時間" form:"time"`
	Nowstatus    string `gorm:"default:'未開始'" json:"目前狀態"`
}

func GetAllList(c *gin.Context) {
	var user []User
	db.DB.Find(&user)
	c.JSON(http.StatusOK, user)
}
func Get3List(c *gin.Context) {
	var user []User
	db.DB.Order("Important desc, Dolist ").Limit(3).Find(&user)
	c.JSON(http.StatusOK, user)
}

func UpdateList(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	if user.Nowstatus == "進行中" || user.Nowstatus == "完成" {
		//db.DB.Table("ToDoList8").Updates(map[string]interface{}{"Last_update2": time.Now().Unix()})
		user.Last_update2 = time.Now().Unix()
		db.DB.Save(&user)
		c.JSON(http.StatusOK, gin.H{
			"message": "已更改紀錄!",
			"data":    &user,
		})
	} else {
		c.JSON(487, gin.H{"error": "狀態錯誤哦 !!"})
	}
}

func DeleteList(c *gin.Context) {
	var user User

	todoID := c.Param("id")
	db.DB.First(&user, todoID)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "沒有此代辦事項哦!"})
		return
	}
	db.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "刪除成功囉!"})

}

func CreateList(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	if user.Important <= 5 {
		user.Last_update1 = time.Now().Unix()
		if user.Sort == "work" || user.Sort == "other" {
			db.DB.Create(&user)
			c.JSON(http.StatusOK, gin.H{
				"message": "已成功紀錄!",
				"data":    &user,
			})
		} else {
			c.JSON(487, gin.H{"error": "分類錯誤!!"})
		}

	} else {
		c.JSON(487, gin.H{"error": "重要程度只能1~5 !!"})
	}

}
