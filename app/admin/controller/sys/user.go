package sys

import (
	"github.com/gogf/gf/net/ghttp"
)

func World(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}
