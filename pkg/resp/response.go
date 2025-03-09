package resp

import (
	"fmt"
	"net/http"
	"simple/pkg/consts"

	"github.com/gin-gonic/gin"
)

/*
   @NAME    : response
   @author  : 清风
   @desc    :
   @time    : 2025/3/6 23:20
*/

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
	TID  string `json:"tid,omitempty"`
}

type PageResp struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
}

func NewResponse(code int, data any, msg string) *Response {
	return &Response{
		Code: code,
		Data: data,
		Msg:  msg,
	}
}

func NewPageResp(list any, total int64) *PageResp {
	return &PageResp{
		List:  list,
		Total: total,
	}
}

func ok(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, NewResponse(200, data, "ok"))
}

func okNil(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, NewResponse(200, nil, "ok"))
}

func fail(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, NewResponse(consts.GC(err), nil, err.Error()))
}

// Res 通用响应
func Res(ctx *gin.Context, err error, data ...any) {
	if err != nil {
		fail(ctx, err)
		return
	}
	if len(data) == 0 {
		okNil(ctx)
		return
	}

	ok(ctx, data)
}

// RefreshToken 刷新token
func RefreshToken(ctx *gin.Context, token, refresh string, expire int64) {
	ctx.Header("u", token)
	ctx.Header("refresh", refresh)
	ctx.Header("expire", fmt.Sprintf("%d", expire))

	ctx.JSON(http.StatusUnauthorized, NewResponse(consts.GC(consts.ErrUnauthorized), nil, consts.ErrUnauthorized.Error()))
}

// 用于处理404错误
func NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, NewResponse(consts.GC(consts.ErrNotFound), nil, consts.ErrNotFound.Error()))
}
