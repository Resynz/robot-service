/**
 * @Author: Resynz
 * @Date: 2020/12/30 16:44
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ping")
}
