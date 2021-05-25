package router

import (
	"fmt"
	_ "gf-practice/app/api"
	"gf-practice/common/service/middleware/router"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	//注册路由
	// fmt.Print("我是注册路由init\n")
	// fmt.Printf("%v", router.GroupList)
	// fmt.Print("\n----\n")
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("哈喽世界！")
	})
	if len(router.GroupList) > 0 {
		for _, routerGroup := range router.GroupList {
			serverName := ""
			if routerGroup.ServerName == "api" {
				serverName = "api"
			}
			fmt.Println(routerGroup.RelativePath)
			g.Server(serverName).Group(routerGroup.RelativePath, func(group *ghttp.RouterGroup) {
				group.Middleware(routerGroup.Handlers...)
				for _, r := range routerGroup.Router {
					// fmt.Println(r)
					switch r.Method {
					case "ANY":
						group.ALL(r.RelativePath, r.HandlerFunc)
					case "GET":
						group.GET(r.RelativePath, r.HandlerFunc)
					case "POST":
						group.POST(r.RelativePath, r.HandlerFunc)
					case "PUT":
						group.PUT(r.RelativePath, r.HandlerFunc)
					case "HEAD":
						group.HEAD(r.RelativePath, r.HandlerFunc)
					case "PATCH":
						group.PATCH(r.RelativePath, r.HandlerFunc)
					case "DELETE":
						group.DELETE(r.RelativePath, r.HandlerFunc)
					case "OPTIONS":
						group.OPTIONS(r.RelativePath, r.HandlerFunc)
					case "CONNECT":
						group.CONNECT(r.RelativePath, r.HandlerFunc)
					case "TRACE":
						group.TRACE(r.RelativePath, r.HandlerFunc)
					}
				}
			})
		}
	}
}
