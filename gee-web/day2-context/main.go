package main

import (
	"net/http"

	"gee"
)

// Result:

// (1)
// $ curl -i http://localhost:9999/
// HTTP/1.1 200 OK
// Content-Type: text/html
// Date: Sun, 17 Sep 2023 07:14:48 GMT
// Content-Length: 18

// <h1>Hello Gee</h1>%

// (2)
// $ curl http://localhost:9999/hello?name=thanhtran
// hello thanhtran, you're at /hello

// (3)
// $ curl "http://localhost:9999/login" -X POST -d 'username=thanhtran&password=1234'
// {"password":"1234","username":"thanhtran"}

// (4)
// $ curl "http://localhost:9999/xxx"
// 404 NOT FOUND: /xxx

func main() {
	r := gee.New()

	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=thanhtran
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
