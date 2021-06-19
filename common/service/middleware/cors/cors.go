package cors

import (
	"github.com/gogf/gf/net/ghttp"
)

// 允许跨域
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
