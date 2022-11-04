package userserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"unicode/utf8"
)

type UserInfo struct {
	Id   string `validate:"uuid" json:"id"`           //UUID 类型
	Name string `validate:"checkName" json:"name"`    //自定义校验
	Age  uint8  `validate:"min=0,max=130" json:"age"` //  0<=Age<=130
}

var valildate *validator.Validate

func Login(c *gin.Context) {

	valildate = validator.New()
	valildate.RegisterValidation("checkName", checkNameFunc)

	var user UserInfo

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "请求参数错误！") //请求的参数可以多，但不能少
		return
	}
	//校验
	err = valildate.Struct(user)
	if err != nil {
		//输出错误的校验值
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("错误的字段：", e.Field())
			fmt.Println("错误的值：", e.Value())
			fmt.Println("错误的tag：", e.Tag())
		}
		c.JSON(http.StatusBadRequest, "数据校验失败"+err.Error())
		return
	}
	c.JSON(http.StatusOK, "数据成功！")
}

func checkNameFunc(f validator.FieldLevel) bool {
	count := utf8.RuneCountInString(f.Field().String())
	if count >= 2 && count <= 12 {
		return true
	}
	return false
}
