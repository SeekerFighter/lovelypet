package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lovelypet/src/com/lovelypet/constant"
)

//返回客户端网络请求结果
func Make(results... interface{}) (gin.H,error) {
	if results == nil || len(results) < 0 {
		return nil,errors.New(constant.ParamNil)
	}
	size := len(results)
	res := gin.H{}
	if size >= 1 {
		res["code"] = results[0]
	}
	if size >= 2 {
		res["message"] = results[1]
	}
	if size >= 3 {
		res["data"] = results[2]
	}
	return res,nil
}