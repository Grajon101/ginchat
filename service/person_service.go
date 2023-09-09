package service

import (
	"ginchat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostPerson
// @Summary 测试post
// @Schemes
// @Description 测试post
// @Tags 测试模块
// @param  person body models.Person false "数据"
// @Accept json
// @Produce json
// @Success 200 {object} models.Person
// @Router /postPerson [post]
func PostPerson(c *gin.Context) {
	var person models.Person
	if err := c.BindJSON(&person); err != nil {
		// 如果不能解析JSON，返回一个错误
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 在这里处理我们得到的JSON数据
	// 返回一个成功的JSON响应
	c.JSON(http.StatusOK, gin.H{
		"name":    person.Name,
		"age":     person.Age,
		"address": person.Address,
	})
}
