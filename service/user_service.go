package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GetUserList
// @Summary 获取用户
// @Schemes
// @Description 获取用户集合
// @Tags 用户模块
// @Accept json
// @Produce json
// @Success 200 {object} []models.UserBasic
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"code":    1,
		"message": "成功",
		"data":    data,
	})

	//fmt.Println(data)
}

// CreateUser
// @Summary 新增用户
// @Schemes
// @Description 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	passWord := c.Query("password")
	repassWord := c.Query("repassword")
	if passWord != repassWord {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "两次密码不一致！",
			"data":    "",
		})
		return
	}
	data := models.GetUserByName(user.Name)
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "用户已注册！",
			"data":    "",
		})
		return
	}
	salt := fmt.Sprintf("%06d", rand.Int31())
	passWord = utils.MakePassword(passWord, salt)
	user.PassWord = passWord
	user.Salt = salt
	models.CreateUser(&user)
	c.JSON(200, gin.H{
		"code":    1,
		"message": "新增用户成功！",
		"data":    user,
	})
}

// DeleteUser
// @Summary 删除用户
// @Schemes
// @Description 删除用户
// @Tags 用户模块
// @param id query string false "用户id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(&user)
	c.JSON(200, gin.H{
		"code":    1,
		"message": "删除用户成功！",
		"data":    user,
	})

}

// CreateUser
// @Summary 更新用户
// @Schemes
// @Description 更新用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "修改失败：" + err.Error(),
			"data":    "",
		})
		return
	}

	models.UpdateUser(&user)
	c.JSON(200, gin.H{
		"code":    1,
		"message": "更新用户成功！",
		"data":    user,
	})
}

// GetUserByNameAndPwd
// @Summary 根据用户名和密码获取用户
// @Schemes
// @Description 根据用户名和密码获取用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","message","data"}
// @Router /user/getUserByNameAndPwd [get]
func GetUserByNameAndPwd(c *gin.Context) {
	name := c.Query("name")
	passWord := c.Query("password")
	// 判断用户是否存在
	tmpUser := models.GetUserByName(name)

	if tmpUser.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户不存在！",
			"data":    "",
		})
		return
	}
	//校验密码
	pwd := utils.MakePassword(passWord, tmpUser.Salt)
	fmt.Println(pwd)
	flag := utils.ValidPassword(passWord, tmpUser.Salt, tmpUser.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "密码不正确",
			"data":    "",
		})
		return
	}
	user := models.GetUserByNameAndPwd(name, pwd)

	c.JSON(200, gin.H{
		"code":    1,
		"message": "获取用户成功！",
		"data":    user,
	})
}

// 防止跨域站点伪造请求
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 发送消息
func SendMsg(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	for {

		// 读取客户端发送的消息
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}

		log.Printf("Received message: %s\n", msg)
		MsgHandler(c, ws)

		// 发送消息给客户端
		// err = ws.WriteMessage(websocket.TextMessage, msg)
		// if err != nil {
		// 	log.Println("Failed to write message:", err)
		// 	break
		// }
	}

}

func MsgHandler(c *gin.Context, ws *websocket.Conn) {
	//for {
	msg, err := utils.Subscribe(c, utils.PublishKey)
	if err != nil {
		fmt.Println("MsgHandler 发送失败", err)
	}
	tm := time.Now().Format("2006-01-01 15:04:05")
	m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		log.Fatalln(err)
	}

	//}
}
