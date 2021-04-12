package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"gee"
)

/*
(1) render array
$ curl http://localhost:9999/date
<html>
<body>
    <p>hello, Jeremie</p>
    <p>Date: 2021-04-10</p>
</body>
</html>
*/

/*
(2) custom render function
$ curl http://localhost:9999/students
<html>
<body>
    <p>hello, Jeremie</p>
    <p>0: Jeremie is 20 years old</p>
    <p>1: Jack is 22 years old</p>
</body>
</html>
*/

/*
(3) serve static files
$ curl http://localhost:9999/assets/css/style.css
p {
    color: orange;
    font-weight: 700;
    font-size: 20px;
}
*/

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Jeremie", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "Jeremie",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "Jeremie",
			"now":   time.Date(2021, 4, 10, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
