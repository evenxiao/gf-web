package main

import (
	"github.com/gogf/gf/frame/g"
	_ "mypj/boot"
	_ "mypj/router"
)

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	g.Server().Run()

}
