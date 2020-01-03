package result

import (
	"com/lovelypet/constant"
	"errors"
	"github.com/gin-gonic/gin"
)

func Response(results... interface{}) (gin.H,error) {
	if results == nil || len(results) < 0 {
		return nil,errors.New(constant.PARAM_NIL)
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