package basecon

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Tuple struct {
	Key   string
	Value any
}

type BaseController struct {
}

func ReturnSucceed(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"code": "0", "msg": "succeed"})
}

func ReturnSucceedWithArgus(ctx *gin.Context, args ...any) {
	ctx.JSON(http.StatusOK, clusterPairs(GetSucceed(), args...))
}

func ReturnFailed(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"code": "-1", "msg": "failed"})
}

func ReturnFailedWithArgus(ctx *gin.Context, args ...any) {
	ctx.JSON(http.StatusOK, clusterPairs(GetFailed(), args...))
}

func ReturnCustom(ctx *gin.Context, code string, msg string) {
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": msg})
}

func ReturnCustomWithArgus(ctx *gin.Context, code *string, msg *string, args ...any) {
	ctx.JSON(http.StatusOK, clusterPairs(GetCustom(code, msg), args...))
}

func GetSucceed() gin.H {
	return gin.H{"code": "0", "msg": "succeed"}
}

func GetFailed() gin.H {
	return gin.H{"code": "-1", "msg": "failed"}
}

func GetCustom(code *string, msg *string) gin.H {
	return gin.H{"code": code, "msg": msg}
}

func clusterPairs(h gin.H, args ...any) gin.H {
	for _, v := range args {
		h[reflect.ValueOf(v).Field(0).String()] = reflect.ValueOf(v).Field(1).Interface()
	}
	return h
}

func NewJson(key string, value any) Tuple {
	return Tuple{Key: key, Value: value}
}
