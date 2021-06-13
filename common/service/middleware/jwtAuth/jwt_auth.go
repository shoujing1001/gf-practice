package jwtauth

import (
	"fmt"
	"gf-practice/common/response"
	"gf-practice/common/token"

	"github.com/gogf/gf/net/ghttp"
)

func Auth(r *ghttp.Request) {
	inputToken := r.GetHeader("Authorization")
	fmt.Println(inputToken)
	claim, err := token.VerifyAuthToken(inputToken)

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	// jwt中存的uid
	fmt.Println(claim.Claim.StandardClaims.Id)
	// 将当前请求的uid信息保存到请求的上下文c上
	r.SetParam("uid", claim.Claim.StandardClaims.Id)
	//判断是否有新的token生成
	if claim.NewToken != "" {
		r.Response.Header().Set("ssnt", claim.NewToken)
	}
	r.Middleware.Next()
}
