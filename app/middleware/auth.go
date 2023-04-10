package middleware

import (
	"encoding/base64"
	"fmt"
	"gigame.xyz/app/response"
	"gigame.xyz/app/vars"
	"gigame.xyz/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

type Headers struct {
	Authorization string `json:"Authorization"`
}

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headers := Headers{}
		if err := ctx.ShouldBindHeader(&headers); err != nil {
			response.AuthFail(ctx)
			return
		}
		if headers.Authorization == "" {
			response.AuthFail(ctx)
			return
		}
		token := strings.Split(headers.Authorization, " ")
		if len(token) != 2 || len(token[1]) < 10 {
			response.AuthFail(ctx)
			return
		}
		key := vars.YmlConfig.GetString(fmt.Sprintf("BoxKeys.%s.Key", token[0]))
		if key == "" {
			response.AuthFail(ctx)
			return
		}
		ek := vars.YmlConfig.GetString(fmt.Sprintf("BoxKeys.%s.EKey", token[0]))
		if !checkSecret(key, token[1], ek) {
			response.AuthFail(ctx)
			return
		}
		ctx.Next()
	}
}

func checkSecret(k, s, ek string) bool {
	decodeString, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return false
	}
	ds, err := utils.AesDecrypt(decodeString, []byte(k))
	if err != nil {
		return false
	}
	splits := strings.Split(string(ds), "-")
	if splits[0] != ek {
		return false
	}
	return true
}
