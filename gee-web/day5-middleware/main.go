package main

import (
	"gee"
	"net/http"
	"time"
)

// Result:

// (1) global middleware Logger
// $ curl http://localhost:9999/
// <h1>Hello Gee</h1>
// >>> log
// 2023/09/17 15:12:02 [200] / in 5.041µs

// (2) global + group middleware
// $ curl http://localhost:9999/v2/hello/thanhtran
// {"message":"Internal Server Error"}

// >>> log
// 2023/09/17 15:12:43 [500] /v2/hello/thanhtran in 60.264µs for group v2
// 2023/09/17 15:12:43 [500] /v2/hello/thanhtran in 74.021µs

import (
	"log"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(http.StatusInternalServerError, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger()) // global middleware
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/thanhtran
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
