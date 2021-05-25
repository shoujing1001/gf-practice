package hello

import (
	"fmt"

	"github.com/gogf/gf/net/ghttp"
)

func World(r *ghttp.Request) {
	fmt.Println("hello world console")
	r.Response.Writeln("Hello World!")
}
