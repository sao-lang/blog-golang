package middlewares

import (
	"blog/internal/global"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 构造返回错误格式相应
// func NewBindFailedResponse(tag string) *Response {
// 	return &Response{Code: WrongArgs, Msg: "wrong argument", Tag: tag}
// }

// reqVal表示具体struct类型
func ReqCheck(reqVal interface{}) func(ctx *gin.Context) {
	var reqType reflect.Type = nil
	if reqVal != nil {
		// 从interface{}还原，提取原值类型
		value := reflect.Indirect(reflect.ValueOf(reqVal))
		// 运行时: 拿到校验体原始类型
		reqType = value.Type()
	}

	return func(c *gin.Context) {
		// tag := c.Request.RequestURI

		var req interface{} = nil
		if reqType != nil {
			// 原始类型
			req = reflect.New(reqType).Interface()
			// 原始类型校验
			if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
				// 结构体绑定出错
				// c.JSON(http.StatusOK, NewBindFailedResponse(tag))
				global.ResponseError(c, http.StatusBadRequest, err)
				// 终止执行链
				c.Abort()
				return
			}
		}
		// 无需校验, 执行链往下
		c.Next()
	}
}
