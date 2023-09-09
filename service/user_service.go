package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 获取用户
// @Schemes
// @Description 获取用户集合
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {object} []models.UserBasic
// @Router /uesr/getlist [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})

	//fmt.Println(data)
}
