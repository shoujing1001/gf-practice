package main

import (
	"fmt"
	_ "gf-practice/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	serverSwitch := g.Cfg().GetBool("status.admin")
	apiSwitch := g.Cfg().GetBool("status.api")

	if serverSwitch {
		g.Server().Start()
	}

	if apiSwitch {
		address := g.Cfg().GetString("api.Address")
		api := g.Server("api")
		api.SetAddr(address)
		api.Start()
		fmt.Println("api running")
	}

	g.Wait()
}
