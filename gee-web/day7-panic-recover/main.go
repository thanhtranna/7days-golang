package main

import (
	"net/http"

	"gee"
)

/*
$ curl "http://localhost:9999"
Hello Jeremie
$ curl "http://localhost:9999/panic"
{"message":"Internal Server Error"}
$ curl "http://localhost:9999"
Hello Geektutu
>>> log
2021/04/12 15:29:31 Route  GET - /
2021/04/12 15:29:31 Route  GET - /panic
2021/04/12 15:29:35 [200] / in 7.897µs
2021/04/12 15:29:45 runtime error: index out of range [100] with length 1
Traceback:
        /usr/local/go/src/runtime/panic.go:969
        /usr/local/go/src/runtime/panic.go:88
        /home/Users/Documents/Golang/7days-golang/gee-web/day7-panic-recover/main.go:47
        /home/Users/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/context.go:41
        /home/Users/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/recovery.go:37
        /home/Users/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/context.go:41
        /home/Users/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/logger.go:15
        /home/Users/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/context.go:41
        /home/Users/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/router.go:99
        /home/Users/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/gee.go:130
        /usr/local/go/src/net/http/server.go:2837
        /usr/local/go/src/net/http/server.go:1925
        /usr/local/go/src/runtime/asm_amd64.s:1374
2021/04/12 15:29:45 [500] /panic in 96.121µs
*/

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"Jeremie"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
