package hello

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"mypj/app/service"
)

type Hello struct {

}

func (h *Hello) Index(r *ghttp.Request) {
	// Hello is a demonstration route handler for output "Hello World!".
	r.Response.Writeln("Hello World!")

}

func (h *Hello) Users(r *ghttp.Request) {
	u, err := service.GetUserInfo()
	if err != nil {

	}
	var n = 3
	//r.Response.WriteJson(u)

	//r.Response.WriteTpl("user.html", g.Map{"name":"awen", "id":1})
	r.Response.WriteTpl("users_list.html", g.Map{
		"list" : u,
		"name": "用户列表",
		"n":n,
	})
}
