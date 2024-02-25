package op

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/simon/orderManager/internal/db"
	logger "github.com/simon/orderManager/internal/log"
	"github.com/simon/orderManager/internal/model"
	"github.com/simon/orderManager/pkg/util"
	"github.com/simon/orderManager/server/common"
)

func Login(c *gin.Context) {
	var login model.User
	if err := c.ShouldBind(&login); err == nil {
		fmt.Printf("login: %v\n", login)
		password, err := db.GetByUsername(login.Name)
		if err != nil {
			common.ErrorWithMsg(c, "未查询到账户", 403)
			return
		}
		if password.Password != login.Password {
			logger.Log.Sugar().Info("%s 登陆失败，账户密码错误.", login.Name)
			common.ErrorWithMsg(c, "账户密码错误", 401)
			return
		}
		token, err := util.CrateToken(login.Name)
		common.SuccessRespond(c, token)
		return
	}
}

func CreateUser(c *gin.Context) {
	var login model.User
	if err := c.ShouldBind(&login); err == nil {
		err = db.CreateUser(login)
		if err != nil {
			common.ErrorWithMsg(c, "创建用户失败", 500)
			logger.Log.Sugar().Error("创建用户失败")
			return
		}
		common.SuccessRespond(c)
		return
	}
	common.ErrorWithMsg(c, "参数解析失败", 500)
}

func CurrentUser(c *gin.Context) {
	username := c.GetString("username")
	user := c.Request.URL.User.Username()
	fmt.Printf("user: %v\n", user)
	if s, err := db.GetByUsername(username); err == nil {
		common.SuccessRespond(c, s)
		return
	}
	common.ErrorWithMsg(c, "参数解析失败", 500)

}
