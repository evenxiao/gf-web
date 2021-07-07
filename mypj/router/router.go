package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"mypj/app/api/hello"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		hello := new(hello.Hello)
		group.ALL("/hello", hello)

		group.GET("/test",func(r *ghttp.Request) {
			r.Response.WriteTpl("index.html", g.Map{
				"title": "列表页面",
				"show": true,
				"listData": g.List{
					g.Map{
						"date":    "2020-04-01",
						"name":    "朱元璋",
						"address": "江苏110号",
					},
					g.Map{
						"date":    "2020-04-02",
						"name":    "徐达",
						"address": "江苏111号",
					},
					g.Map{
						"date":    "2020-04-03",
						"name":    "李善长",
						"address": "江苏112号",
					},
				}})
		})
	})
}
