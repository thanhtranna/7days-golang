package main

import (
	"gee"
	"net/http"
)

// Result
// $ curl "http://localhost:9999"
// Hello ThanhTran
// $ curl "http://localhost:9999/panic"
// {"message":"Internal Server Error"}
// $ curl "http://localhost:9999"
// Hello ThanhTran

// >>> log
// 2023/09/17 15:47:32 Route  GET - /
// 2023/09/17 15:47:32 Route  GET - /panic
// 2023/09/17 15:47:52 [200] / in 8.351µs
// 2023/09/17 15:48:01 runtime error: index out of range [100] with length 1
// Traceback:
// 		/usr/local/opt/go/libexec/src/runtime/panic.go:914
// 		/usr/local/opt/go/libexec/src/runtime/panic.go:114
// 		/Users/thanhtran/Documents/Golang/7days-golang/gee-web/day7-panic-recover/main.go:48
// 		/Users/thanhtran/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/context.go:41
// 		/Users/thanhtran/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/recovery.go:37
// 		/Users/thanhtran/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/context.go:41
// 		/Users/thanhtran/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/logger.go:15
// 		/Users/thanhtran/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/context.go:41
// 		/Users/thanhtran/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/router.go:101
// 		/Users/thanhtran/Documents/Golang/7days-golang/gee-web/day7-panic-recover/gee/gee.go:131
// 		/usr/local/opt/go/libexec/src/net/http/server.go:2939
// 		/usr/local/opt/go/libexec/src/net/http/server.go:2010
// 		/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1651
//
// 2023/09/17 15:48:01 [500] /panic in 200.794µs
// 2023/09/17 15:48:49 [200] / in 3.258µs

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello ThanhTran\n")
	})

	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"thanhtran"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
